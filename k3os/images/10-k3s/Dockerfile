ARG REPO
ARG TAG
FROM ${REPO}/k3os-base:${TAG}

ARG ARCH
ENV ARCH ${ARCH}
ENV VERSION v1.19.4+k3s1

ARG http_proxy=$http_proxy
ARG https_proxy=$https_proxy
ARG no_proxy=$no_proxy

ENV http_proxy $http_proxy
ENV https_proxy $https_proxy
ENV no_proxy $no_proxy
ENV HTTP_PROXY $HTTP_PROXY
ENV HTTPS_PROXY $HTTPS_PROXY


ADD https://raw.githubusercontent.com/rancher/k3s/${VERSION}/install.sh /output/install.sh
ENV INSTALL_K3S_VERSION=${VERSION} \
    INSTALL_K3S_SKIP_START=true \
    INSTALL_K3S_BIN_DIR=/output
RUN chmod +x /output/install.sh
RUN env | grep  -i proxy
RUN cat /output/install.sh
RUN /output/install.sh
RUN echo "${VERSION}" > /output/version
