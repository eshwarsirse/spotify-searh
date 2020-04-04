FROM golang:1.12 AS builder

WORKDIR /src

COPY ./ ./

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /app .

FROM alpine:latest

RUN apk --no-cache add ca-certificates tzdata

WORKDIR /app

COPY --from=builder  /app .

EXPOSE 8080

ENTRYPOINT ["./app"]
