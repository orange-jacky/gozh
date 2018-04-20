FROM golang:1.10-alpine3.7 AS build

# Install tools required to build the project
# We need to run `docker build --no-cache .` to update those dependencies
RUN apk add --no-cache git  && \
    go get -u github.com/tools/godep

RUN mkdir /myapp
WORKDIR /myapp
COPY . /myapp


ENV APP_NAME gozh
RUN godep go build -ldflags '-d -w -s' -o $APP_NAME . 

CMD  ./$APP_NAME conf/cf.json  > stdout.log 2>&1
