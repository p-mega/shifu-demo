# 第一步编译go文件
FROM golang:alpine AS builder
WORKDIR /build
ADD go.mod .
COPY . .
RUN go build -o shifu-demo main.go

# 第二步部生成镜像
FROM alpine
WORKDIR /build
COPY --from=builder /build/hello /build/hello
CMD ["./shifu-demo"]
