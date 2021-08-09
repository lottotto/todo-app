
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


# circle CI でのEC2にSCPでデプロイする方法
- add_ssh_keysを利用し、CircleCIにSSHの鍵(秘密鍵)を登録しておく.その際hostnameが聞かれるので、IPアドレスかhostnameを入力すること
- fingerprintsに登録したSSHの鍵を指定することでcircleCIが回っているコンテナに秘密鍵がダウンロードされる。それと同時にsshのconfigもダウンロードされる
- `scp -o StrictHostKeyChecking=no <SRC_FILE_PATH> <SSH_USER>@<SSH_ADRESS>:<DST_FILE_PATH> `を実行する