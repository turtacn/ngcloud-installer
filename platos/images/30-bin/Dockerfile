ARG REPO
ARG TAG
FROM ${REPO}/platos-rootfs:${TAG} as rootfs

ARG REPO
ARG TAG
FROM ${REPO}/platos-progs:${TAG} as progs

ARG REPO
ARG TAG
FROM ${REPO}/platos-base:${TAG}

COPY --from=rootfs /output/rootfs.squashfs /usr/src/
COPY install.sh /output/platos-install.sh
COPY --from=progs /output/platos /output/platos
RUN echo -n "_sqmagic_" >> /output/platos
RUN cat /usr/src/rootfs.squashfs >> /output/platos
