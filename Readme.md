
[![CircleCI](https://circleci.com/gh/lottotto/todo-app.svg?style=svg)](https://circleci.com/gh/lottotto/todo-app)
[![codecov](https://codecov.io/gh/lottotto/todo-app/branch/main/graph/badge.svg?token=HPYQ70GGG3)](https://codecov.io/gh/lottotto/todo-app)
[![Docker Automated build](https://img.shields.io/docker/automated/danish9966/todo-app)](https://hub.docker.com/repository/docker/danish9966/todo-app)
# ToDo app
簡単なtodo アプリを実装します.


# in order not to forget how to setting eleastic apm

- elasticsearch と kibanaとapmserver を起動する
- ローカルで`go get go.elastic.co/apm`を実行しapm agent をインストールする
- 環境変数`ELASTIC_APM_SERVER_URL`にapmserver のURLを設定する.
  - ex) `http://x.x.x.x:8200`
- elastic APMでDBの性能などを確認できるようにするためには、`db.QueryContext`メソッドを利用し、`c.Request().Context()`を引数に追加すること