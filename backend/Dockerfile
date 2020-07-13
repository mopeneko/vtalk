FROM golang:1.14-alpine AS builder

WORKDIR /usr/local/src/app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build -o app

FROM alpine

COPY --from=builder /usr/local/src/app/app /usr/local/bin/
ENTRYPOINT [ "app" ]

