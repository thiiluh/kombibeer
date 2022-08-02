FROM golang

ADD . /go/src/kombibeer

WORKDIR /go/src/kombibeer

RUN go mod tidy

RUN go build -o kombibeer

ENTRYPOINT ["./kombibeer"]

EXPOSE 8080