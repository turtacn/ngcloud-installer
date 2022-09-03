ARG REPO
ARG TAG
FROM ${REPO}/platos-tar:${TAG} as tar

ARG REPO
ARG TAG
FROM ${REPO}/platos-iso:${TAG} as iso

ARG REPO
ARG TAG
FROM ${REPO}/platos-kernel:${TAG} as kernel

ARG REPO
ARG TAG
FROM ${REPO}/platos-base:${TAG}
ARG ARCH

COPY --from=iso /output/platos.iso /output/platos-${ARCH}.iso
