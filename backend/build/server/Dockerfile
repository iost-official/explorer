FROM golang:1.16 as builder
LABEL Maintainer="Souki Aoba <souki.aoba@ginco.co.jp>"

RUN mkdir -p /go/src/github.com/GincoInc/iost-explorer
COPY . /go/src/github.com/GincoInc/iost-explorer/backend
WORKDIR /go/src/github.com/GincoInc/iost-explorer/backend
RUN make clean
RUN make

ENTRYPOINT ["./explorer"]
