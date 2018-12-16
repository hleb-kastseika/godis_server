FROM golang

ADD . /go/src/godis_server

RUN go install godis_server

ENTRYPOINT /go/bin/godis_server

EXPOSE 9090