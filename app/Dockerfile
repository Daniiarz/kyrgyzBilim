FROM golang:1.16-alpine as builder

WORKDIR /go/src/app

ENV GO111MODULE=on

RUN apk add upx
RUN go get github.com/cespare/reflex
RUN go get github.com/json-iterator/go

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -tags=jsointer -o run .

RUN upx ./run

FROM alpine:latest as prod
RUN apk --no-cache add ca-certificates

WORKDIR /usr/src/app

#Copy executable from builder
COPY --from=builder /go/src/app/run /usr/src/app/run

ENV MEDIA_URL="http://164.90.234.21/media"
ENV GIN_MODE=release

CMD ["./run"]