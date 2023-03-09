FROM golang:1.18 AS builder

# ログに出力する時間をJSTにするため、タイムゾーンを設定
ENV TZ /usr/share/zoneinfo/Asia/Tokyo

# ワーキングディレクトリの設定
WORKDIR /go/src/app

# ModuleモードをON
ENV GO111MODULE=on

# 必要な依存関係をインストール
RUN apt-get update && apt-get install -y \
    git \
    && rm -rf /var/lib/apt/lists/*

# go.modを参照し、go.sumファイルの更新を行う
COPY go.mod go.sum ./
RUN go mod download

# ホストのファイルをコンテナの作業ディレクトリに移行
COPY . .

EXPOSE 5000

# ビルド
RUN go build -o /go/bin/app

CMD ["/go/bin/app"]
