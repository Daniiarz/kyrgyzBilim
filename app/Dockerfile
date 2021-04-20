FROM golang:1.16-alpine as builder

WORKDIR /go/src/app

ENV GO111MODULE=on

RUN go get github.com/cespare/reflex
RUN go get github.com/json-iterator/go

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -tags=jsointer -o run .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/

#Copy executable from builder
COPY --from=builder /go/src/app/run .

ENV GIN_MODE=release

CMD ["./run"]