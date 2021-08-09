
[![CircleCI](https://circleci.com/gh/lottotto/todo-app.svg?style=svg)](https://circleci.com/gh/lottotto/todo-app)
[![codecov](https://codecov.io/gh/lottotto/todo-app/branch/main/graph/badge.svg?token=HPYQ70GGG3)](https://codecov.io/gh/lottotto/todo-app)
[![Docker Automated build](https://img.shields.io/docker/automated/danish9966/todo-app)](https://hub.docker.com/repository/docker/danish9966/todo-app)
# ToDo api
[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Flottotto%2Ftodo-app.svg?type=shield)](https://app.fossa.com/projects/git%2Bgithub.com%2Flottotto%2Ftodo-app?ref=badge_shield)

簡単なtodo アプリを実装します.

## License
[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Flottotto%2Ftodo-app.svg?type=large)](https://app.fossa.com/projects/git%2Bgithub.com%2Flottotto%2Ftodo-app?ref=badge_large)
簡単なtodo アプリを実装します.


# elastic APM の設定方法メモ

- elasticsearch と kibanaとapmserver を起動する
- ローカルで`go get go.elastic.co/apm`を実行しapm agent をインストールする
- 環境変数`ELASTIC_APM_SERVER_URL`にapmserver のURLを設定する.
  - ex) `http://x.x.x.x:8200`
- elastic APMでDBの性能などを確認できるようにするためには、`db.QueryContext`メソッドを利用し、`c.Request().Context()`を引数に追加すること

# circle CI でのEC2にSCPでデプロイする方法
- add_ssh_keysを利用し、CircleCIにSSHの鍵(秘密鍵)を登録しておく.その際hostnameが聞かれるので、IPアドレスかhostnameを入力すること
- fingerprintsに登録したSSHの鍵を指定することでcircleCIが回っているコンテナに秘密鍵がダウンロードされる。それと同時にsshのconfigもダウンロードされる
- `scp -o StrictHostKeyChecking=no <SRC_FILE_PATH> <SSH_USER>@<SSH_ADRESS>:<DST_FILE_PATH> `を実行する
