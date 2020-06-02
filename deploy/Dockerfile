FROM golang:1.12.12 as build
ENV CGO_ENABLED=0
ENV GOPATH=/go

WORKDIR /go/src/github.com/jsfong/sample-http-go-server
COPY src src
COPY pkg pkg

RUN go get $PWD/src
RUN go build -o /sample-http-go-server $PWD/src

# FROM gcr.io/distroless/static:debug
# FROM gcr.io/distroless/static:latest
FROM golang:alpine

COPY --from=build /sample-http-go-server /go/bin/sample-http-go-server

ENTRYPOINT ["/go/bin/sample-http-go-server"]