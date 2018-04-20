FROM golang:1.10-alpine3.7 AS build

# Install tools required to build the project
# We need to run `docker build --no-cache .` to update those dependencies
RUN apk add --no-cache git  && \
    go get -u github.com/tools/godep


ENV APP_NAME gozh
ENV APP_DIR $GOPATH/src/$APP_NAME

WORKDIR $APP_DIR
COPY .  $APP_DIR

EXPOSE 80

RUN godep go build -o $APP_NAME . 

CMD  ./$APP_NAME conf/cf.json >> logs/stdout_`date +%Y%m%d-%H:%M:%d`.log 2>&1
