FROM golang:1.26-alpine AS build

WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download

ENV CGO_ENABLED=0

COPY . ./
RUN go build -ldflags="-s -w" -trimpath -o unms-exporter main.go

FROM alpine:3.21

RUN apk add --no-cache tzdata ca-certificates && \
    addgroup -S exporter && \
    adduser -S -G exporter exporter

COPY --from=build /src/unms-exporter /usr/local/bin/

USER exporter

EXPOSE 9806

CMD ["/usr/local/bin/unms-exporter"]
