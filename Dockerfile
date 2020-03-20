# ベースとなるDockerイメージ指定
FROM golang:latest

# コンテナ内に作業ディレクトリを作成
RUN mkdir /go/src/work

# コンテなログイン時のディレクトリ指定
WORKDIR /go/src/work

# ホストのファイルをコンテナディレクトリに以降
ADD . /go/src/work