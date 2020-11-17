FROM golang:1.15

ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct \
    GIN_MODE=release \
    PORT=8001

WORKDIR /app

COPY . .

RUN go build .

EXPOSE 8001

ENTRYPOINT ["./app"]