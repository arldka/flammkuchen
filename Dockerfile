FROM golang:1.24 AS builder

WORKDIR /cmd/server

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -o /go/bin/flammkuchen ./cmd/server

FROM scratch

COPY --from=builder /go/bin/flammkuchen /

EXPOSE 8080

CMD ["/flammkuchen"]
