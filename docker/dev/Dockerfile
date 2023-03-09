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

# ホストのファイルをコンテナの作業ディレクトリに移行
COPY ../../. /go/src/app/

# 依存関係のダウンロードとビルド
RUN go mod download && go build -o /go/bin/app ./cmd/api

EXPOSE 5000

# アプリケーションの起動
CMD ["/go/bin/app"]
