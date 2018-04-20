FROM golang:1.10 AS build




RUN mkdir /myapp
WORKDIR /myapp

COPY . /myapp

RUN go build -o gozh .

RUN  chmod +x /myapp/start.sh
CMD  /myapp/start.sh  gozh
