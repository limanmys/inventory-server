FROM golang:1.21-bookworm

LABEL maintainer="Zeki Ahmet Bayar <zeki@liman.dev>"

RUN apt update && apt install libpopt-dev libhdb9-heimdal libgssapi3-heimdal -y

WORKDIR /opt/build

COPY go.mod .

COPY go.sum .

RUN go mod download

COPY . .

RUN GOOS=linux CGO_ENABLED=1 go build -ldflags="-s -w" -o /opt/build/inventory-server cmd/server/main.go

RUN mkdir reports

COPY scripts/start.sh /tmp/start.sh

RUN sed -e 's/$/ -type=test/' /tmp/start.sh

RUN ["chmod", "755", "/tmp/start.sh"]

RUN ["chmod", "+x", "/tmp/start.sh"]

ENTRYPOINT ["/tmp/start.sh"]