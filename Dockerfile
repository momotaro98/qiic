FROM golang:1.12.5-stretch as builder

WORKDIR /workdir

ENV GO111MODULE="on"
COPY go.mod go.sum ./
RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go build -o qiic

FROM alpine:3.8

WORKDIR /root/

# Multi stage build function of Docker
COPY --from=builder /workdir/qiic .

CMD ./qiic
