FROM golang:latest
WORKDIR /go/src/MANOMASH
COPY . /go/src/MANOMASH
RUN go mod tidy
#MySQL設定ファイルをイメージ内にコピー
# ADD ./my.cnf /etc/mysql/conf.d/my.cnf
# Airをインストール
RUN go install github.com/cosmtrek/air@latest
