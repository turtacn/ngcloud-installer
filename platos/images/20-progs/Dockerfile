ARG REPO
ARG TAG
ARG VERSION
FROM ${REPO}/platos-gobuild:${TAG} as gobuild

ENV LINUXKIT v0.8

FROM gobuild as linuxkit
RUN git clone https://github.com/linuxkit/linuxkit.git $GOPATH/src/github.com/linuxkit/linuxkit
WORKDIR $GOPATH/src/github.com/linuxkit/linuxkit/pkg/metadata
RUN git checkout -b current $LINUXKIT
RUN gobuild -o /output/metadata

FROM gobuild as platos
ARG VERSION
#COPY /cmd/ $GOPATH/src/github.com/rancher/k3os/cmd/
COPY /pkg/ $GOPATH/src/github.com/rancher/k3os/pkg/
COPY /main.go $GOPATH/src/github.com/rancher/k3os/
COPY /vendor/ $GOPATH/src/github.com/rancher/k3os/vendor/
WORKDIR $GOPATH/src/github.com/rancher/k3os
RUN gobuild -o /output/k3os

FROM gobuild
COPY --from=linuxkit /output/ /output/
COPY --from=platos /output/ /output/
WORKDIR /output
RUN git clone --branch v0.7.0 https://github.com/ahmetb/kubectx.git \
 && chmod -v +x kubectx/kubectx kubectx/kubens

# Additional utils
RUN apk --no-cache add zstd
