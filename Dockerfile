FROM golang:1.20.3-alpine AS builder

WORKDIR /link_shorter

COPY . .

RUN go mod tidy

RUN CGO_ENABLED=0 go build -o main ./cmd/main.go

FROM alpine

COPY --from=builder /link_shorter/main .
COPY --from=builder /link_shorter/.env .

EXPOSE 8080

ENTRYPOINT ["/main"]
