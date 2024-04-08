FROM golang:1.20-alpine as builder

ENV CGO_ENABLED=0

WORKDIR /build

COPY . ./

RUN go mod tidy

RUN go build -o go-example

FROM alpine:3.16

COPY --from=builder /build/go-example /bin

EXPOSE 3333

CMD ["/bin/go-example"]