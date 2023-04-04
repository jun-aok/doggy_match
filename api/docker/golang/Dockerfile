FROM golang:1.19.0-alpine3.16

SHELL ["/bin/ash", "-c"]
WORKDIR /go/src/app
COPY ./app ./
EXPOSE 8080
ENV GO111MODULE=on

RUN apk add --no-cache alpine-sdk

# Golang ホットリロード(freshのインストール)
RUN go install github.com/pilu/fresh@latest

# Golang 環境構築
RUN go install github.com/rogpeppe/godef@latest
RUN go install golang.org/x/tools/cmd/goimports@latest
RUN go install golang.org/x/tools/cmd/gorename@latest
RUN go install sourcegraph.com/sqs/goreturns@latest

# debug
RUN go install github.com/uudashr/gopkgs/v2/cmd/gopkgs@latest
RUN go install github.com/ramya-rao-a/go-outline@latest
RUN go install github.com/cweill/gotests/gotests@latest
RUN go install github.com/fatih/gomodifytags@latest
RUN go install github.com/josharian/impl@latest
RUN go install github.com/haya14busa/goplay/cmd/goplay@latest
RUN go install github.com/go-delve/delve/cmd/dlv@latest
RUN go install github.com/go-delve/delve/cmd/dlv@master
RUN go install honnef.co/go/tools/cmd/staticcheck@latest
RUN go install golang.org/x/tools/gopls@latest

# RUN go install github.com/labstack/echo@latest
RUN go get github.com/labstack/echo/v4
RUN go get github.com/labstack/echo/v4/middleware
RUN go get firebase.google.com/go/v4
RUN go get github.com/go-sql-driver/mysql
RUN go get github.com/shopspring/decimal

# ENV GOOGLE_APPLICATION_CREDENTIALS /go/src/app/cred/fire-base.json
