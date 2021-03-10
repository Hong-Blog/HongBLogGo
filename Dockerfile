FROM golang:1.15 as builder

ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct

WORKDIR /app
COPY . .
RUN GOOS=linux GOARCH=amd64 go build -tags netgo -o serviceLogin cmd/serviceLogin/main.go
RUN GOOS=linux GOARCH=amd64 go build -tags netgo -o servicePost cmd/servicePost/main.go
RUN GOOS=linux GOARCH=amd64 go build -tags netgo -o serviceUserCenter cmd/serviceUserCenter/main.go

RUN mkdir publish && cp serviceLogin servicePost serviceUserCenter publish

FROM alpine as login
WORKDIR /app
COPY --from=builder /app/publish/serviceLogin .
ENV GIN_MODE=release \
    PORT=8082 \
    REDIS_URL="redis://127.0.0.1:6379/1" \
    PREFIX_USER_SERVICE="http://127.0.0.1:18081" \
    DATA_SOURCE="root:123qwe@tcp(127.0.0.1:3306)/hongblog?parseTime=true"
EXPOSE 8082
ENTRYPOINT ["./serviceLogin"]

FROM alpine as post
WORKDIR /app
COPY --from=builder /app/publish/servicePost .
ENV GIN_MODE=release \
    PORT=80 \
    REDIS_URL="redis://127.0.0.1:6379/1" \
    DATA_SOURCE="root:123qwe@tcp(127.0.0.1:3306)/hongblog?parseTime=true"
EXPOSE 80
ENTRYPOINT ["./servicePost"]

FROM alpine as userCenter
WORKDIR /app
COPY --from=builder /app/publish/serviceUserCenter .
ENV GIN_MODE=release \
    PORT=8081 \
    REDIS_URL="redis://127.0.0.1:6379/1" \
    DATA_SOURCE="root:123qwe@tcp(127.0.0.1:3306)/hongblog?parseTime=true"
EXPOSE 8081
ENTRYPOINT ["./serviceUserCenter"]