FROM golang:latest
WORKDIR /go/src/MANOMASH
COPY . /go/src/MANOMASH
RUN go mod tidy
# Airをインストール
RUN go install github.com/cosmtrek/air@latest
