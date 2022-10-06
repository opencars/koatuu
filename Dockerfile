FROM golang:1.19-alpine AS build

ENV GO111MODULE=on

WORKDIR /go/src/app

LABEL maintainer="ashanaakh@gmail.com"

RUN apk add bash ca-certificates git gcc g++ libc-dev

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o /go/bin/server ./cmd/grpc-server/main.go && \
    go build -o /go/bin/parser ./cmd/http-server/main.go && \
    go build -o /go/bin/parser ./cmd/parser/main.go

FROM alpine

RUN apk update && apk upgrade && apk add curl

WORKDIR /app

COPY --from=build /go/bin/ ./
COPY ./config ./config

EXPOSE 8080

CMD ["./http-server"]
