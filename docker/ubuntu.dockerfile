# Use a base image with the desired OS (e.g., Ubuntu, Debian, etc.)
FROM ubuntu:22.04

RUN echo 'root:root' | chpasswd && \
    printf '#!/bin/sh\nexit 0' > /usr/sbin/policy-rc.d

RUN apt-get update && \
 apt-get install -y openssh-server sudo curl iputils-ping \
    systemd systemd-sysv dbus dbus-user-session iptables-persistent

RUN printf "systemctl start systemd-logind" >> /etc/profile && \
    systemctl disable ufw nftables && \
    systemctl enable ssh iptables && \
    useradd -rm -d /home/sshuser -s /bin/bash -g root -G sudo sshuser && \
    echo 'sshuser:password' | chpasswd
# Expose the SSH port
EXPOSE 22 6443
# Start SSH server on container startup

ENTRYPOINT ["/usr/sbin/init"]