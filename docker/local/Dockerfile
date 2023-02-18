# Dockerfile
FROM golang:1.18 AS builder

ARG GITHUB_TOKEN

# ログに出力する時間をJSTにするため、タイムゾーンを設定
ENV TZ /usr/share/zoneinfo/Asia/Tokyo

# ワーキングディレクトリの設定
WORKDIR /go/src/app

# ModuleモードをON
ENV GO111MODULE=on

RUN export GOPRIVATE=github.com/ficilcom

RUN git config --global url."https://${GITHUB_TOKEN}:x-oauth-basic@github.com/".insteadOf "https://github.com/"

# ホストのファイルをコンテナの作業ディレクトリに移行
COPY . .

# AIRのインストール
RUN go install github.com/cosmtrek/air@latest


# go.modを参照し、go.sumファイルの更新を行う
RUN go mod tidy

EXPOSE 5000

# localではホットリロードを有効にしたいのでairで起動する
CMD ["air", "-c", ".air.toml"]
