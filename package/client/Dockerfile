FROM golang:1.17-alpine3.15 AS builder

WORKDIR /simplegrpc/client

COPY . .

RUN go build -o /client ./cmd/client

FROM alpine:3.15 AS runtime

COPY --from=builder /client /client

CMD /client

