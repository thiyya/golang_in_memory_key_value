FROM golang:1.16-alpine
RUN apk update && apk upgrade && apk add --no-cache bash git
RUN apk add --update curl && rm -rf /var/cache/apk/*

WORKDIR /golang_in_memory_key_value
COPY . /golang_in_memory_key_value
RUN go mod init golang_in_memory_key_value
RUN go mod vendor
RUN go build

EXPOSE 8080

CMD [ "go run main.go" ]