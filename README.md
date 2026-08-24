# photoalbum

Каркас REST API-сервиса на Go (модуль `photoalbum`). На данный момент реализованы:

- HTTP API-сервер (`net/http` + `gorilla/mux`) с единственным эндпоинтом `/hello`;
- слой доступа к данным (`internal/app/store`) поверх PostgreSQL (`lib/pq`);
- модель пользователя с хешированием пароля через `bcrypt` (`internal/app/model`);
- миграция для таблицы `users`.

## Структура проекта

```
cmd/apiserver/            точка входа, парсинг флагов и конфига
internal/app/apiserver/   HTTP-сервер: роутинг, обработчики, конфиг сервера
internal/app/store/       доступ к БД: Store, UserRepository, тестовые хелперы
internal/app/model/       доменные модели (User)
internal/pkg/logger/      минимальный консольный логгер с уровнями
configs/apiserver.toml    конфиг сервера по умолчанию
migrations/               SQL-миграции (up/down)
```

## Требования

- Go 1.17+
- PostgreSQL (для работы с БД и для тестов пакета `store`)

## Конфигурация

Настройки читаются из TOML-файла (по умолчанию `./configs/apiserver.toml`):

```toml
bind_addr = ":8080"
log_level = "debug"

[store]
database_url = "host=localhost dbname=photoalbum user=dbadmin password=hitomi sslmode=disable"
```

Путь к конфигу можно переопределить флагом:

```
./apiserver --config-path=./configs/apiserver.toml
```

## База данных

Создайте базу и накатите миграцию `migrations/20210829201206_create_users.up.sql`, например:

```bash
createdb photoalbum
psql -d photoalbum -f migrations/20210829201206_create_users.up.sql
```

(структура файлов миграций совместима с [sql-migrate](https://github.com/rubenv/sql-migrate), при желании можно использовать его вместо ручного накатывания)

## Сборка и запуск

```bash
make build          # go build -v ./cmd/apiserver
./apiserver
```

Проверка:

```bash
curl http://localhost:8080/hello
```

## Тесты

```bash
make test
```

Тесты пакета `internal/app/store` обращаются к реальной БД: адрес берётся из переменной окружения `DATABASE_URL`, а если она не задана — используется `host=localhost dbname=photoalbum user=dbadmin password=hitomi sslmode=disable` (см. `internal/app/store/store_test.go`). Перед их запуском база должна существовать и иметь применённую миграцию.
