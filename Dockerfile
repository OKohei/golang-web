# ベースとなるDockerイメージ
FROM golang:latest

# ソース格納先
ENV SRC_DIR=/go/src/github.com/

# ワーキングディレクトリの指定
WORKDIR $SRC_DIR

# 依存モジュールをインストール
RUN go get github.com/beego/bee
RUN go get -u github.com/astaxie/beego