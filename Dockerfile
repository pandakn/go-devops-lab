FROM golang:1.21-alpine as builder

WORKDIR /app

COPY go.mod ./

RUN go mod download && go mod tidy

COPY . .

RUN CGO_ENABLED=0 go build -o /bin/app

FROM alpine:latest

RUN apk add --no-cache ca-certificates

COPY --from=builder /bin/app /bin/app

EXPOSE 8080

CMD ["/bin/app"]