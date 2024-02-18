# リリース用のビルドを行うコンテナイメージを作成するステージ
# デプロイ用コンテナに含めるバイナリを作成するコンテナ
FROM golang:1.22-bullseye as deploy-builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -trimpath -ldflags "-w -s" -o app


# マネージドサービス上で動かすことを想定したリリース用のコンテナイメージを作成するステージ
# デプロイ用のコンテナ
FROM debian:bullseye-slim as deploy

RUN apt-get update

COPY --from=deploy-builder /app/app .

CMD ["./app"]