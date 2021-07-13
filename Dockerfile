FROM golang:alpine
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN apk add git
RUN apk add libc-dev
RUN apk add gcc
RUN apk add vim
RUN go mod tidy
RUN go build -o main .
CMD ["/app/main"]
