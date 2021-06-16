FROM golang:latest

ENV GO111MODULE=on

RUN go install github.com/a-h/templ/cmd/templ@latest

WORKDIR /src
COPY go.mod /src/.
COPY go.sum /src/.

RUN go mod download

COPY . /src/.

RUN templ generate

RUN go mod verify
RUN go mod vendor

RUN go build -o /src/tmp/cmd ./cmd
EXPOSE 80
CMD ["/src/tmp/cmd"]
