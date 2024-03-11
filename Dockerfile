FROM golang:latest

RUN apt-get update && apt-get install ffmpeg -y
RUN apt-get install python3-opencv -y
RUN apt-get install libpng-dev -y

WORKDIR /usr/src/app

RUN go install github.com/cosmtrek/air@latest

COPY ./src/ .

WORKDIR /usr/src/libs
COPY ./libs/ .

WORKDIR /usr/src/libs/pngloss
RUN make
RUN make install

WORKDIR /usr/src/app
RUN go mod tidy
