ARG ROS_DISTRO=iron

FROM ros:${ROS_DISTRO}

# This is pushed to the docker hub to be accessible to the GitHub Actions

# Install the golang package
RUN apt update && apt install -y \
  wget

# Install the specific golang version
RUN cd /tmp \
  && wget https://go.dev/dl/go1.23.2.linux-amd64.tar.gz \
  && tar -C /usr/local -xzf go1.23.2.linux-amd64.tar.gz

ENV PATH="$PATH:/usr/local/go/bin"

ENV CGO_CFLAGS="-I/opt/ros/${ROS_DISTRO}/include/action_msgs \
  -I/opt/ros/${ROS_DISTRO}/include/builtin_interfaces \
  -I/opt/ros/${ROS_DISTRO}/include/rcl \
  -I/opt/ros/${ROS_DISTRO}/include/rcl_action \
  -I/opt/ros/${ROS_DISTRO}/include/rcl_yaml_param_parser \
  -I/opt/ros/${ROS_DISTRO}/include/rcutils \
  -I/opt/ros/${ROS_DISTRO}/include/rmw \
  -I/opt/ros/${ROS_DISTRO}/include/rosidl_dynamic_typesupport \
  -I/opt/ros/${ROS_DISTRO}/include/rosidl_runtime_c \
  -I/opt/ros/${ROS_DISTRO}/include/rosidl_typesupport_interface \
  -I/opt/ros/${ROS_DISTRO}/include/service_msgs \
  -I/opt/ros/${ROS_DISTRO}/include/std_msgs \
  -I/opt/ros/${ROS_DISTRO}/include/type_description_interfaces \
  -I/opt/ros/${ROS_DISTRO}/include/unique_identifier_msgs "

ENV CGO_LDFLAGS='-L/opt/ros/${ROS_DISTRO}/lib -Wl,-rpath=/opt/ros/${ROS_DISTRO}/lib '

# Source the ros2 installation (based on the original Dockerfile: https://github.com/josegron/fog-ros-baseimage/tree/main)
SHELL [ "/bin/bash", "-c" ]

ENV BASH_ENV="/opt/ros/$ROS_DISTRO/setup.bash"
RUN echo "source /opt/ros/$ROS_DISTRO/setup.bash" >> /etc/bash.bashrc