
FROM golang:1.18.3-alpine3.16

WORKDIR /app

COPY go.mod ./
COPY *.go ./
COPY static ./static

RUN go build -o /go-ci-cd

EXPOSE 3000

CMD ["/go-ci-cd"]