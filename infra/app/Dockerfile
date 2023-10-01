FROM golang:1.20-alpine as builder

WORKDIR /app
COPY .. .
RUN go mod download
RUN go build -o /out/app ./cmd/app

FROM alpine:latest
WORKDIR /app
COPY --from=builder /out/app /app/app
EXPOSE 8001 9001
ENTRYPOINT ["/app/app"]