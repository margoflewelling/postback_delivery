FROM golang:1.14

COPY . /app/go_app

RUN apt-get -y update
RUN go get github.com/go-redis/redis

ENTRYPOINT /app/go_app/deliver.go

EXPOSE 8080
