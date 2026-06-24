#!/usr/bin/env bash
# Source the ROS 2 environment, then exec the container's command. This ensures
# AMENT_PREFIX_PATH, LD_LIBRARY_PATH and friends are set even for non-login,
# non-interactive invocations such as `docker run <image> make test`.
set -e

# shellcheck disable=SC1090
source "/opt/ros/${ROS_DISTRO}/setup.bash"

exec "$@"
