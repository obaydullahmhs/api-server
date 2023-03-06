FROM golang:alpine AS builder
RUN mkdir /api-server
COPY . /api-server
WORKDIR /api-server
RUN go build .

FROM alpine 
WORKDIR /api-server
COPY --from=builder /api-server/ /api-server/

ENTRYPOINT ["./api-server"]
