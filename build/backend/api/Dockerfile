FROM golang:1.16
LABEL Maintainer="Souki Aoba <souki.aoba@ginco.co.jp>"

RUN mkdir -p /go/src/github.com/GincoInc/iost-explorer
COPY ./backend /go/src/github.com/GincoInc/iost-explorer/backend
WORKDIR /go/src/github.com/GincoInc/iost-explorer/backend
RUN make clean
RUN make

EXPOSE 8088
ENTRYPOINT ["./explorer"]
