FROM golang:latest


ENV GO111MODULE=on
ENV GOFLAGS=-mod=vendor

WORKDIR /src
COPY go.mod /src/.
COPY go.sum /src/.

RUN go mod download

RUN go mod verify
RUN go mod vendor

COPY . /src/.

RUN go build -o /src/tmp/cmd ./cmd
EXPOSE 80
CMD ["/src/tmp/cmd"]
