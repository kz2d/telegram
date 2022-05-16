FROM golang:1.16 as base

#FROM base as dev

RUN mkdir /app
ADD . /app
WORKDIR /app

RUN go mod download

RUN go build -o main /app/cmd/telegramd/main.go


EXPOSE ["8080"]
CMD ["/app/main"]