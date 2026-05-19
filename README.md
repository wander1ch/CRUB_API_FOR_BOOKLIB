# CRUD API for Book Library

Простое REST API для управления библиотекой книг, написанное на Go с использованием PostgreSQL.  
Контейнеризировано с помощью Docker Compose.

## Возможности

- Создание книги (POST /books)
- Получение списка всех книг (GET /books)
- Получение книги по ID (GET /books/{id})
- Обновление книги (PUT /books/{id})
- Удаление книги (DELETE /books/{id})

## Технологии

- Go 1.21+
- PostgreSQL 18.3
- Docker & Docker Compose
- Стандартная библиотека Go (`net/http`, `database/sql`)

## Запуск проекта (Docker Compose)

1. **Клонируйте репозиторий**
   ```bash
   git clone https://github.com/wander1ch/CRUD_API_FOR_BOOKLIB.git
   cd CRUD_API_FOR_BOOKLIB
   ```

## Пример запросов сurl

- создание книги
  ```bash
  curl -X POST http://localhost:8080/books \
   -H "Content-Type: application/json" \
   -d '{"title":"1984","author":"George Orwell","year":1949}'
  ```
- получить все книги
  ```bash
  curl http://localhost:8080/books
  ```
- получение книги по ID
  ```bash
  curl http://localhost:8080/books/1
  ```
- обновить книгу
  ```bash
  curl -X PUT http://localhost:8080/books/1 \
   -H "Content-Type: application/json" \
   -d '{"title":"Nineteen Eighty-Four","author":"George Orwell","year":1949}'
  ```
- удалить книгу
  ````bash
  curl -X DELETE http://localhost:8080/books/1
  '```'
  ````

## полезные команды

- Перезапустить только веб‑сервис

  ```bash
  docker compose restart web
  ```

- Посмотреть логи

  ```bash
  docker compose logs web -f
  ```

- Остановить всё
  ```bash
  docker compose down
  ```
- Остановить и удалить данные БД
  ```bash
  docker compose down -v
  ```
