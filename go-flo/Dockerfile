# Alpine Linux 是一个社区开发的面向安全应用的轻量级Linux发行版，适合用来做Docker镜像
FROM golang:1-alpine as build

RUN apk add --no-cache git

# Golang 包管理工具
RUN go get -u github.com/Masterminds/glide

# FLO 的 Golang 语言全节点
RUN git clone https://github.com/bitspill/flod /go/src/github.com/bitspill/flod
#RUN git clone https://github.com/GuoxiW/flod /go/src/github.com/bitspill/flod
#RUN git clone https://github.com/btcsuite/btcd /go/src/github.com/btcsuite/btcd
#RUN git clone https://github.com/btcsuite/btclog /go/src/github.com/btcsuite/btclog
#RUN git clone https://github.com/btcsuite/btcutil /go/src/github.com/btcsuite/btcutil
#RUN git clone https://github.com/btcsuite/go-socks /go/src/github.com/btcsuite/go-socks
#RUN git clone https://github.com/btcsuite/goleveldb /go/src/github.com/btcsuite/goleveldb
#RUN git clone https://github.com/davecgh/go-spew /go/src/github.com/davecgh/go-spew
#RUN git clone https://github.com/decred/dcrd /go/src/github.com/decred/dcrd
#RUN git clone https://github.com/aead/siphash /go/src/github.com/aead/siphash
#RUN git clone https://github.com/btcsuite/snappy-go /go/src/github.com/btcsuite/snappy-go
#RUN git clone https://github.com/kkdai/bstream /go/src/github.com/kkdai/bstream
#RUN go get -u golang.org/x/crypto/ripemd160 
WORKDIR /go/src/github.com/bitspill/flod
RUN glide install
RUN go install . ./cmd/...


FROM alpine:latest

# 构建参数用于设置环境变量
ARG network=mainnet

COPY --from=build /go/bin/flod /flod/bin/flod
COPY ./${network}.conf /flod/flod.conf

VOLUME ["/root/.flod/"]
ENTRYPOINT ["/flod/bin/flod", "--configfile=/flod/flod.conf"]
