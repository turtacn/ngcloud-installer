FROM golang:1.13-alpine AS gobuild
## alpine source: turta
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories

ENV http_proxy ""
ENV https_proxy ""
ENV no_proxy ""
ENV HTTP_PROXY ""
ENV HTTPS_PROXY ""

RUN apk -U add git gcc linux-headers musl-dev make libseccomp libseccomp-dev bash
COPY gobuild /usr/bin/
RUN rm -f /bin/sh && ln -s /bin/bash /bin/sh
WORKDIR /output