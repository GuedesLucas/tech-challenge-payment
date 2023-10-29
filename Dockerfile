FROM golang:1.21 AS builder

WORKDIR /go/src/app
COPY . .

COPY .env .env

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./main.go

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/
COPY --from=builder /go/src/app/main .

COPY --from=builder /go/src/app/.env .env

EXPOSE 8080

CMD ["./main"]