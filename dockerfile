FROM golang:alpine AS builder

# 禁用CGO，减小镜像体积
ENV CGO_ENABLED 0
# go mod 镜像源修改
ENV GOPROXY https://goproxy.cn,direct

WORKDIR /build
COPY . .
RUN go build -ldflags="-s -w" -o shifu-demo main.go


FROM alpine
WORKDIR /build
COPY --from=builder /build/shifu-demo /build/shifu-demo
CMD ["./shifu-demo"]
