ARG REPO
ARG TAG
FROM ${REPO}/platos-k8s:${TAG} as k3s

ARG REPO
ARG TAG
FROM ${REPO}/platos-bin:${TAG} as bin

ARG REPO
ARG TAG
FROM ${REPO}/platos-base:${TAG} as base
ARG VERSION

COPY --from=k3s /output/  /output/platos/system/k3s/
COPY --from=bin /output/  /output/platos/system/platos/${VERSION}/

WORKDIR /output/platos/system/k3s
RUN mkdir -vp $(cat version) /output/sbin
RUN mv -vf crictl ctr kubectl /output/sbin/
RUN ln -sf $(cat version) current
RUN mv -vf install.sh current/k3s-install.sh
RUN mv -vf k3s current/
RUN rm -vf version *.sh
RUN ln -sf /platos/system/k3s/current/k3s /output/sbin/k3s

WORKDIR /output/platos/system/platos
RUN ln -sf ${VERSION} current
RUN ln -sf /platos/system/platos/current/platos /output/sbin/platos
RUN ln -sf platos /output/sbin/init
