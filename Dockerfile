FROM golang:1.18.4-alpine3.15 as builder

WORKDIR /bin

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm64 \
   go build -a -ldflags "-s -w" \
  -o app .

FROM alpine:3.15 AS bin

COPY --from=builder /bin/app /app

COPY scripts/entrypoint.sh /entrypoint.sh

RUN chmod +x /entrypoint.sh
RUN apk --no-cache add ca-certificates

EXPOSE 8080

ENTRYPOINT ["./entrypoint.sh"]