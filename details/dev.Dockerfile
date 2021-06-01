# Please keep up to date with the new-version of Golang docker for builder
FROM golang:latest

RUN apt update && apt upgrade -y && \
    apt install -y git \
    make openssh-client

RUN go get -u github.com/a-h/templ/cmd@v0.0.88

WORKDIR /app 

RUN curl -fLo install.sh https://raw.githubusercontent.com/cosmtrek/air/master/install.sh \
    && chmod +x install.sh && sh install.sh && cp ./bin/air /bin/air

CMD air