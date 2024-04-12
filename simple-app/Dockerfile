FROM golang:alpine as builder

WORKDIR /app
COPY . .

RUN go build -o main .

FROM alpine
COPY --from=builder /app/main /app/main

EXPOSE 8080
CMD ["/app/main"]
