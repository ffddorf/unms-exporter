FROM golang:1.18-alpine AS build

WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download

ENV CGO_ENABLED=0

COPY . ./
RUN go build -ldflags="-s -w" -trimpath -o unms-exporter main.go

FROM alpine

RUN apk add --no-cache tzdata ca-certificates
COPY --from=build /src/unms-exporter /usr/local/bin/

CMD ["/usr/local/bin/unms-exporter"]
