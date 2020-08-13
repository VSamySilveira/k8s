FROM golang:alpine as builder

ARG APP_VERSION=0.1.0

COPY ./gowebsrv /go/src/gowebsrv

WORKDIR /go/src/gowebsrv

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o gowebsrv .

FROM scratch

COPY --from=builder /go/src/gowebsrv/gowebsrv /usr/bin/gowebsrv

EXPOSE 8000

CMD ["gowebsrv"]