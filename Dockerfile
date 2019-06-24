FROM golang:1.12-alpine AS builder
RUN apk update && apk add --no-cache git make ca-certificates && update-ca-certificates
WORKDIR /srv
COPY . .
ENV GO111MODULE=on
ARG CMD=server
ARG BIN=server
RUN go build -o $BIN -ldflags "-s -w -X main.version={{.Version}} -X main.commit={{.ShortCommit}} -X main.date={{.Date}}" ./cmd/$CMD

FROM scratch

ARG BIN=server
COPY --from=builder /svc/$BIN /usr/local/bin/
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd

USER nobody
EXPOSE 8080
ENTRYPOINT [$BIN, "--port", "8080"]