FROM golang:1.20.0-bullseye

ADD . /project

WORKDIR /project

RUN go mod download

ENTRYPOINT ["go", "run", "./code/todo-app/..."]
