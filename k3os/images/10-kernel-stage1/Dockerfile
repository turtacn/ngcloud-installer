ARG REPO
ARG TAG
FROM ubuntu:focal


ARG http_proxy=$http_proxy
ARG https_proxy=$https_proxy
ARG no_proxy=$no_proxy
ENV http_proxy $http_proxy
ENV https_proxy $https_proxy
ENV no_proxy $no_proxy
ENV HTTP_PROXY $http_proxy
ENV HTTPS_PROXY $https_proxy
RUN env | grep -i proxy
RUN echo 'Acquire::http { Proxy "$HTTP_PROXY"; };' >> /etc/apt/apt.conf.d/01proxy && cat /etc/apt/apt.conf.d/01proxy

RUN apt-get --assume-yes -o Acquire::http::proxy="$http_proxy" -o Acquire::https::proxy="$https_proxy" update
RUN apt-get --assume-yes -o Acquire::http::proxy="$http_proxy" -o Acquire::https::proxy="$https_proxy" install \
    curl \
    initramfs-tools \
    kmod \
    lz4 \
    rsync \
    xz-utils \
 && echo 'r8152' >> /etc/initramfs-tools/modules \
 && echo 'hfs' >> /etc/initramfs-tools/modules \
 && echo 'hfsplus' >> /etc/initramfs-tools/modules \
 && echo 'nls_utf8' >> /etc/initramfs-tools/modules \
 && echo 'nls_iso8859_1' >> /etc/initramfs-tools/modules

ARG ARCH
ENV KVERSION=5.4.0-54-generic
ENV URL=https://github.com/rancher/k3os-kernel/releases/download/5.4.0-54.60-rancher1
ENV KERNEL_XZ=${URL}/kernel-generic_${ARCH}.tar.xz
ENV KERNEL_EXTRA_XZ=${URL}/kernel-extra-generic_${ARCH}.tar.xz
ENV KERNEL_HEADERS_XZ=${URL}/kernel-headers-generic_${ARCH}.tar.xz

# Download kernel
RUN mkdir -p /usr/src
RUN curl -fL $KERNEL_XZ -o /usr/src/kernel.tar.xz
RUN curl -fL $KERNEL_EXTRA_XZ -o /usr/src/kernel-extra.tar.xz
RUN curl -fL $KERNEL_HEADERS_XZ -o /usr/src/kernel-headers.tar.xz

### Extract to /usr/src/root
RUN mkdir -p /usr/src/root && ls /usr/src
RUN cd /usr/src/root && \
    tar xvf /usr/src/kernel.tar.xz > /dev/null
RUN cd /usr/src/root && \
    tar xvf /usr/src/kernel-extra.tar.xz > /dev/null
RUN cd /usr/src/root && \
    tar xvf /usr/src/kernel-headers.tar.xz > /dev/null

# Create initrd
RUN mkdir /usr/src/initrd && \
    rsync -a /usr/src/root/lib/ /lib/ && \
    depmod $KVERSION && \
    mkinitramfs -k $KVERSION -c lz4 -o /usr/src/initrd.tmp

# Generate initrd firmware and module lists
RUN mkdir -p /output/lib && \
    mkdir -p /output/headers && \
    cd /usr/src/initrd && \
    lz4cat /usr/src/initrd.tmp | cpio -idmv && \
    find lib/modules -name \*.ko > /output/initrd-modules && \
    echo lib/modules/${KVERSION}/modules.order >> /output/initrd-modules && \
    echo lib/modules/${KVERSION}/modules.builtin >> /output/initrd-modules && \
    find lib/firmware -type f > /output/initrd-firmware && \
    find usr/lib/firmware -type f | sed 's!usr/!!' >> /output/initrd-firmware

# Copy output assets
RUN cd /usr/src/root && \
    cp -r usr/src/linux-headers* /output/headers && \
    cp -r lib/firmware /output/lib/firmware && \
    cp -r lib/modules /output/lib/modules && \
    cp boot/System.map* /output/System.map && \
    cp boot/config* /output/config && \
    cp boot/vmlinuz-* /output/vmlinuz && \
    echo ${KVERSION} > /output/version
