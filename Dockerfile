FROM golang:1.15 AS build

WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download

COPY . ./
RUN go build -o unms-exporter main.go

FROM ubuntu:20.04

RUN apt-get -q update && apt-get -yq install ca-certificates
COPY --from=build /src/unms-exporter /usr/local/bin/

CMD ["/usr/local/bin/unms-exporter"]
