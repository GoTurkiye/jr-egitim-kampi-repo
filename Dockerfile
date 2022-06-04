FROM golang:1.17-alpine3.15 as builder

WORKDIR /app

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=arm64

COPY . .

RUN go build main.go

FROM alpine

COPY --from=builder /app /app

CMD ["./app/main"]
