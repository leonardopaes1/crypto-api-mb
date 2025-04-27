# Build stage
FROM golang:1.21 as builder

WORKDIR /app
COPY go.mod .
COPY main.go .

RUN CGO_ENABLED=0 GOOS=linux go build -o crypto-api main.go

# Final image
FROM gcr.io/distroless/static:nonroot

WORKDIR /
COPY --from=builder /app/crypto-api .
EXPOSE 8080

USER nonroot:nonroot
ENTRYPOINT ["/crypto-api"]