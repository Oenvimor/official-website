FROM golang:alpine AS builder

ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /build

COPY . .

RUN go mod tidy

RUN go build -o cqupt-hub .

FROM scratch

COPY ./config.yaml /config.yaml

COPY --from=builder /build/cqupt-hub /

ENTRYPOINT ["/cqupt-hub", "/config.yaml"]