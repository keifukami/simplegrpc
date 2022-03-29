FROM golang:1.17-alpine3.15 AS builder

WORKDIR /simplegrpc/server

COPY . .

RUN go build -o /server ./cmd/server

FROM alpine:3.15 AS runtime

COPY --from=builder /server /server

CMD /server

