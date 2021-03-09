FROM golang:1.15 as builder

ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct

WORKDIR /app
COPY . .
RUN GOOS=linux GOARCH=amd64 go build -tags netgo -o app cmd/serviceLogin/main.go

RUN mkdir publish && cp app publish

FROM alpine
WORKDIR /app
COPY --from=builder /app/publish .
ENV GIN_MODE=release \
    PORT=8082

EXPOSE 8082
ENTRYPOINT ["./app"]