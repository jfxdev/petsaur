FROM golang:1.17.7-alpine3.15 as builder

WORKDIR /bin

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm64 \
   go build -a -ldflags "-s -w" \
  -o app .

FROM alpine:3.15.0 AS bin

COPY --from=builder /bin/app /app

COPY scripts/run.sh /run.sh

RUN chmod +x /run.sh
RUN apk --no-cache add ca-certificates

EXPOSE 8080

ENTRYPOINT ["./run.sh"]