FROM golang:alpine as builder

ARG APP_VERSION=0.1.0

COPY ./gowebsrv /go/src/gowebsrv

WORKDIR /go/src/gowebsrv

RUN go build gowebsrv

ENTRYPOINT ["/usr/local/bin/go", "run", "gowebsrv"]

FROM scratch

COPY --from=builder /go/src/gowebsrv/gowebsrv /bin/

EXPOSE 8000

CMD ["gowebsrv"]