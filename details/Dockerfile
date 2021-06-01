FROM golang:latest


ENV GO111MODULE=on
ENV GOFLAGS=-mod=vendor

RUN go get -u github.com/a-h/templ/cmd

WORKDIR /src
COPY go.mod /src/.
COPY go.sum /src/.

RUN go mod download

COPY . /src/.

RUN $(go env GOPATH)/bin/cmd generate

RUN go mod verify
RUN go mod vendor

RUN go build -o /src/tmp/cmd ./cmd
EXPOSE 80
CMD ["/src/tmp/cmd"]
