
## Создать файл миграции
```
migrate create -ext sql -dir migrations first_migration
```

## Накатить миграции
```
migrate -path migrations -database "postgres://dbuser:dbpassword@localhost:5432/dbgo?sslmode=disable" up

migrate -path migrations -database "postgres://dbuser:dbpassword@localhost:5432/dbgo?sslmode=disable" up 1
```

## Откатить миграции
```
migrate -path migrations -database "postgres://dbuser:dbpassword@localhost:5432/dbgo?sslmode=disable" down

migrate -path migrations -database "postgres://dbuser:dbpassword@localhost:5432/dbgo?sslmode=disable" down 1
```

## Переход к указанной версии миграции
```
migrate -path migrations -database "postgres://dbuser:dbpassword@localhost:5432/dbgo?sslmode=disable" goto 20240703054100
```

## Команда починки

Команда переводит состояние в указанную версию, однако сама миграция не применяется:
```
migrate -path migrations -database "postgres://dbuser:dbpassword@localhost:5432/dbgo?sslmode=disable" force 20240703054100
```

## Очистка базы данных:
```
migrate -path migrations -database "postgres://dbuser:dbpassword@localhost:5432/dbgo?sslmode=disable" drop
```