FROM golang:1.19-alpine as builder

WORKDIR /app
COPY .. .
RUN go mod download
RUN go build -o /out/app ./cmd/workers

FROM alpine:latest
WORKDIR /workers
COPY --from=builder /out/app /workers/app
ENTRYPOINT ["/workers/app"]