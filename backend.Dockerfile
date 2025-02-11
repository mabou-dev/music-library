FROM golang:1.23-alpine as builder

WORKDIR /workspace

COPY . .

RUN go mod tidy
RUN go build -o bin/server cmd/server.go

#FROM alpine:latest
FROM scratch

COPY --from=builder /workspace/bin/server server

ARG DEFAULT_PORT=8080
EXPOSE $DEFAULT_PORT

CMD ["./server", "--help"]
