GO111MODULE = on

.ONESHELL:
SHELL = /bin/bash

# default make target
all: build

build:
	cd cmd/rclgo-gen && go build

install:
	cd cmd/rclgo-gen && go install

# Tests run on rmw_zenoh against a private router, so they never join another stack's graph on the
# default port 7447. The port is keyed on the checkout path; override it with ZENOH_TEST_PORT.
define with_zenoh_router
	set -euo pipefail
	port="$${ZENOH_TEST_PORT:-$$((20000 + 16#$$(printf '%s' "rclgo:$(CURDIR)" | md5sum | cut -c1-8) % 10000))}"
	listening() { ss -ltnH "( sport = :$$port )" | grep -q .; }
	if listening; then echo "zenoh test port $$port is in use; set ZENOH_TEST_PORT" >&2; exit 1; fi
	ZENOH_CONFIG_OVERRIDE="listen/endpoints=[\"tcp/127.0.0.1:$$port\"]" \
		/opt/ros/$${ROS_DISTRO}/lib/rmw_zenoh_cpp/rmw_zenohd >/dev/null 2>&1 &
	router=$$!
	trap 'kill $$router 2>/dev/null; wait $$router 2>/dev/null || true' EXIT
	for _ in $$(seq 100); do listening && break; sleep 0.1; done
	if ! listening; then echo "zenoh router did not listen on port $$port" >&2; exit 1; fi
	export RMW_IMPLEMENTATION=rmw_zenoh_cpp
	export ZENOH_CONFIG_OVERRIDE="connect/endpoints=[\"tcp/127.0.0.1:$$port\"]"
	$(1)
endef

.PHONY: test
test:
	$(call with_zenoh_router,go test -count=1 -timeout=90s ./...)

# cgocheck2 makes the runtime check every pointer write into non-Go memory.
.PHONY: test-cgocheck2
test-cgocheck2:
	$(call with_zenoh_router,GOEXPERIMENT=cgocheck2 go test -count=1 -timeout=90s ./...)

.PHONY: test-verbose
test-verbose:
	$(call with_zenoh_router,go test -v -count=1 -timeout=90s ./...)

generate:
	dest_path=internal/msgs

	rm -rf "$$dest_path/"*
	go run ./cmd/rclgo-gen generate \
		--root-path /usr \
		--root-path /opt/ros/${ROS_DISTRO} \
		--dest-path "$$dest_path" \
		--message-module-prefix "github.com/ATIinc/rclgo/$$dest_path" \
		--license-header-path ./license-header.txt \
		--include-go-package-deps ./... \
		--cgo-flags-path "" \
		|| exit 1
	rm "$$dest_path/msgs.gen.go" || exit 1
	go run ./cmd/rclgo-gen generate-rclgo \
		--root-path /usr \
		--root-path /opt/ros/${ROS_DISTRO} \
		--license-header-path ./license-header.txt

lint:
	golangci-lint run ./...
