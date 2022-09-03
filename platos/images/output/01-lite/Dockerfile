ARG REPO
ARG TAG
FROM ${REPO}/platos-tar:${TAG} as tar

ARG REPO
ARG TAG
FROM ${REPO}/platos-base:${TAG}
ARG ARCH

COPY --from=tar /output/userspace.tar /output/platos-rootfs-${ARCH}.tar
RUN gzip /output/platos-rootfs-${ARCH}.tar
