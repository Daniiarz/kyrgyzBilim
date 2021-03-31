FROM golang:1.16 as builder

RUN mkdir /app
ADD . /app
WORKDIR /app

RUN go clean --modcache
RUN go mod download
RUN CGO_ENABLE=0 GOOS=liux go build -a -installsuffix cgo -o main .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
RUN apk add --no-cache git make musl-dev go
COPY --from=builder /app/main .

ENV GOROOT /user/lib/go
ENV GOPATH /go
ENV PATH /go/bin:$PATH

RUN mkdir -p ${GOPATH}/src ${GOPATH}/bin
EXPOSE 80
CMD ["./main"]