FROM ubuntu:16.04

RUN mkdir /myapp
WORKDIR /myapp

COPY . /myapp
RUN  chmod +x /myapp/start.sh
CMD  /myapp/start.sh
