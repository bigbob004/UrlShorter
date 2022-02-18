# UrlShorter
# Укорачиватель ссылок

### Требования

- Ссылка должна быть уникальной и на один оригинальный URL должна ссылаться только одна сокращенная ссылка.
- Ссылка должна быть длинной 10 символов.
- Ссылка должна состоять из символов латинского алфавита в нижнем и верхнем регистре, цифр и символа _ (подчеркивание).

Сервис должен принимать следующие запросы по http:
1. Метод Post, который будет сохранять оригинальный URL в базе и возвращать сокращённый.
2. Метод Get, который будет принимать сокращённый URL и возвращать оригинальный URL.

- Сервис должен быть распространён в виде Docker-образа.
- В качестве хранилища ожидается использовать in-memory решение и postgresql. Какое хранилище использовать указывается параметром при запуске сервиса.
- Покрыть реализованный функционал Unit-тестами.

### Для запуска приложения:

Для запуска в памяти приложения:

```
make run_in_memory
```

Для запуска в БД:
Ели docker-контейнер и миграции не были применены раннее:

```
make docker_and_migrate
```
Если у вас нет образ postgres для docker, то необходимо его скачать с помощью команды:

```
docker pull postgres
```

Если у вас нет утилиты migrate, то вам необходимо её также установить... 
Об этом можно почитать здесь: https://github.com/golang-migrate/migrate/blob/master/cmd/migrate/README.md


Ели docker-контейнер и миграции были применены раннее:

```
make run_in_db
```



# To do:
- запуск приложения в docker-compose.
- unit-тесты.
