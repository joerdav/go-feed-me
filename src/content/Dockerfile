FROM golang:latest


ENV GO111MODULE=on

WORKDIR /src
COPY go.mod /src/.
COPY go.sum /src/.

RUN go mod download

COPY . /src/.

RUN go mod verify
RUN go mod vendor

RUN go build -o /src/app ./cmd
EXPOSE 80
CMD ["/src/app"]
