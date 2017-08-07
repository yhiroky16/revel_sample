FROM golang:1.8.3

ENV GOPATH $GOPATH:/go/src

RUN apt-get update && \
    apt-get upgrade -y

# install revel and revel-cli
RUN go get github.com/revel/revel && \
    go get github.com/revel/cmd/revel

# change current dir
RUN mkdir /go/src/AttendanceManage

EXPOSE 9000