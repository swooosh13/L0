# L0

## Первый запуск
Создание бд
```bash
$ docker-compose up -d
$ docker-compose exec pgdb sh
$ psql -U admin
admin=# create database orders; 
```
Поднятие миграций с помощью [goose](https://github.com/pressly/goose)
```bash
$ make migrate-up
```

## Работа приложения

[service](cmd/app/main.go) - основной сервис  
[publisher](cmd/pub/main.go) - сервис для генерации и публикации в канал валидных/невалидных сообщений

### Логи запуска приложения

Параллельно запустим publisher
```bash
$ make pub
```

```bash
$ go run cmd/app/main.go
{"level":"info","ts":1654768720.9499724,"caller":"logger/zap.go:25","msg":"Config has been loaded"}
{"level":"info","ts":1654768721.0336752,"caller":"logger/zap.go:25","msg":"PostgresComposite has been created successfully"}
{"level":"info","ts":1654768721.082607,"caller":"logger/zap.go:25","msg":"Data has been recovered from db to in-memory cache"}
{"level":"info","ts":1654768721.0831296,"caller":"logger/zap.go:25","msg":"OrderComposite has been created successfully"}
{"level":"info","ts":1654768721.1146371,"caller":"logger/zap.go:25","msg":"Nats connection successfully established"}
{"level":"info","ts":1654768721.121121,"caller":"logger/zap.go:25","msg":"Server has been started"}
{"level":"info","ts":1654768960.8691394,"caller":"logger/zap.go:25","msg":"The order was successfully written to the DB and in-memory cache"}
{"level":"info","ts":1654768963.871098,"caller":"logger/zap.go:25","msg":"The order was successfully written to the DB and in-memory cache"}
{"level":"error","ts":1654768966.8756394,"caller":"logger/zap.go:29","msg":"Received error json format"}
{"level":"info","ts":1654768966.8914194,"caller":"logger/zap.go:25","msg":"The order was successfully written to the DB and in-memory cache"}
{"level":"info","ts":1654768969.9020646,"caller":"logger/zap.go:25","msg":"The order was successfully written to the DB and in-memory cache"}
```
