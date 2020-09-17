FROM golang:1-alpine as build

RUN apk add --no-cache git curl

# 定义构建时需要的参数　　制定go语言的包管理软件dep版本为0.5.0
ARG DEP_VERSION="0.5.0"

# go语言的包管理软件
RUN curl -L -s https://github.com/golang/dep/releases/download/v${DEP_VERSION}/dep-linux-amd64 \
        -o /go/bin/dep && chmod +x /go/bin/dep

# 将资源编译到二进制语言中的软件
RUN go get -u github.com/gobuffalo/packr/v2/packr2
RUN git clone https://github.com/oipwg/oip /go/src/github.com/oipwg/oip
WORKDIR /go/src/github.com/oipwg/oip
RUN dep ensure -v
WORKDIR /go/src/github.com/oipwg/oip/cmd/oipd
RUN packr2 -v

# GitCommitHash为最新的commit id
# BuildDate为当下时间
RUN go build -ldflags \
        "-X github.com/oipwg/oip/version.GitCommitHash=e929a52fd \
        -X github.com/oipwg/oip/version.BuildDate=$(date +'%Y.%m.%d.%H%M%S') \
        -X github.com/oipwg/oip/version.BuiltBy=Docker \
        -X 'github.com/oipwg/oip/version.GoVersion=$(go version)' \
        -s -w" -o oipd github.com/oipwg/oip/cmd/oipd
RUN cp oipd /go/bin/oipd

FROM alpine:latest

# 在测试环境中,因此更改此处的docker为testnet
#ARG network=mainnet
ARG network=testnet

COPY --from=build /go/bin/oipd /oip/bin/oipd
COPY ./config.${network}.yml /oip/config.yml

# 在测试环境中,因此更改此处的端口
# EXPOSE 1606
EXPOSE 11606
ENV LOG=*
ENTRYPOINT ["/oip/bin/oipd", "--appdir=/oip"]