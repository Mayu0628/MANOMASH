# MANOMASH

## 使用ライブラリ
- gorilla: Routing, Middleware,Session

## 起動方法
### 初回 or Dockerfileに変更があった場合
- `docker-compose up --build`
### 2回目以降
- `docker-compose up`
### DBの初期化をするときは
- `docker-compose down --volume`
### DockerのDBにアクセスするときは
- `docker exec -it db-for-techread bash`
- `mysql -u ユーザー名 -p`