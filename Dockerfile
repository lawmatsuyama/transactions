# --- Stage 1:
FROM golang:1.18-alpine as builder
ENV BUILD_PATH=/go/src/github.com/lawmatsuyama/transactions
RUN apk update && apk add --no-cache curl gcc git libc-dev
WORKDIR ${BUILD_PATH}
COPY . .
RUN go mod download
RUN go test ./... -v
RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 \
    go build -a -ldflags '-linkmode=external' -installsuffix cgo -o transactions
RUN ls -lah ${BUILD_PATH}/transactions
RUN cp ${BUILD_PATH}/transactions /bin/transactions

FROM alpine:3.11
RUN apk update && apk add --no-cache ca-certificates tzdata libc6-compat
COPY --from=builder /bin/transactions /transactions
CMD ["/transactions"]
EXPOSE 8080
