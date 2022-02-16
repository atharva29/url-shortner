FROM golang:1.17.1-alpine3.14
WORKDIR /build
COPY . .

WORKDIR /build

RUN GOOS=linux CGO_ENABLED=0 go build -a -ldflags '-s -w -extldflags "-static"' -o app .

FROM alpine:3.12
RUN apk --no-cache add ca-certificates
RUN chown 1001:1001 /
USER 1001
WORKDIR /app
COPY --from=0 /build/app .
CMD ["./app"]