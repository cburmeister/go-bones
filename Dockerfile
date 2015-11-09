FROM golang:onbuild
RUN go get github.com/cburmeister/go-bones
EXPOSE 8000
