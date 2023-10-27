FROM golang:1.21

WORKDIR /app

COPY . .

RUN go build -o main

FROM scratch

COPY --from=builder /app/main /main

EXPOSE 7575

CMD ["/main"]