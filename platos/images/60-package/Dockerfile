ARG REPO
ARG TAG
FROM ${REPO}/platos-kernel:${TAG} as kernel

ARG REPO
ARG TAG
FROM ${REPO}/platos-package:${TAG}
ARG VERSION

COPY --from=kernel /output/ /output/k3os/system/kernel/

WORKDIR /output/k3os/system/kernel
RUN mkdir -vp $(cat version)
RUN ln -sf $(cat version) current
RUN mv -vf initrd kernel.squashfs current/
RUN rm -vf version vmlinuz
