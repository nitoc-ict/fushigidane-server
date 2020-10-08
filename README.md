# fushigidane-server

## Description
CCC2020に出場した、チームﾌｼｷﾞﾀﾞﾈのサーバーです

## Usage
以下の説明は本リポジトリをクローン後、そのディレクトリまで移動していることを前提としています。
### DB
テストのdocker-composeでいい感じに立てます
```
$ sudo docker-compose up -d
```

### start App Server
```
$ go build

$ ./fushigidane-server
```
