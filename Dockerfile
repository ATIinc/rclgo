FROM ros:iron

# This is pushed to the docker hub to be accessible to the GitHub Actions

# Install the golang package
RUN apt update && apt install -y \
  wget

# Install the specific golang version
RUN cd /tmp \
&& wget https://go.dev/dl/go1.23.2.linux-amd64.tar.gz \
&& tar -C /usr/local -xzf go1.23.2.linux-amd64.tar.gz

ENV PATH="$PATH:/usr/local/go/bin"

# Set the default entrypoint
# ENTRYPOINT [ "/bin/bash", "-c"]