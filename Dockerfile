# Builder
FROM --platform=$BUILDPLATFORM golang:1.26.2-alpine3.23 AS builder

RUN apk add --no-cache git

WORKDIR /usr/src/demo-container

COPY go.mod go.sum ./

RUN go mod download

COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o /usr/bin/demo-container

# Binary
FROM --platform=$BUILDPLATFORM alpine:3.23

WORKDIR /usr/bin

COPY --from=builder /usr/bin/demo-container ./

EXPOSE 3000

CMD ["./demo-container"]
