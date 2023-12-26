FROM golang:1.21-bullseye

LABEL maintainer="Zeki Ahmet Bayar <zeki@liman.dev>"

WORKDIR /opt/build

COPY go.mod .

COPY go.sum .

RUN go mod download

COPY . .

RUN GOOS=linux CGO_ENABLED=1 go build -ldflags="-s -w" -o /opt/build/inventory-server cmd/server/main.go

COPY scripts/start.sh /tmp/start.sh

RUN ["chmod", "755", "/tmp/start.sh"]

RUN ["chmod", "+x", "/tmp/start.sh"]

ENTRYPOINT ["/tmp/start.sh"]