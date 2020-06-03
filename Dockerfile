
FROM golang:1.14 AS builder

WORKDIR /src/calculadora

COPY /src/calculadora .

RUN GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags="-w -s" -o /bin/calculadora

FROM scratch

COPY --from=builder /bin/calculadora /calculadora

ENTRYPOINT ["/calculadora"]