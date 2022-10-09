FROM golang:latest

WORKDIR /go-blog

ADD . .

RUN go mod download

RUN go install github.com/githubnemo/CompileDaemon@latest

ENTRYPOINT CompileDaemon -command="./go-blog"