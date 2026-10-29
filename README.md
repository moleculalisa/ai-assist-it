# Todo API Service
REST API для управления пользователями и задачами (Todo).
Проект реализован на Go и содержит интерактивную документацию API с использованием Swagger.
### Требования
- Go 1.22 или выше

## Установка и запуск
Клонируйте репозиторий:
git clone https://github.com/moleculalisa/ai-assist-it.git
cd ai-assist-it
cd ai-assist-it
Установите зависимости проекта:
go mod download
Запустите приложение:
go run main.go
После запуска API будет доступно по адресу:
http://localhost:8080

## Зависимости
Основные зависимости проекта:
- github.com/go-chi/chi/v5 — маршрутизация HTTP-запросов;
- github.com/swaggo/http-swagger/v2 — отображение Swagger UI;
- github.com/swaggo/swag — генерация Swagger-документации.

### Swagger
Swagger-документация генерируется с помощью Swaggo.
Для повторной генерации документации выполните:
swag init -g main.go -o docs
В результате в директории docs/ создаются:
docs.go
swagger.json
swagger.yaml
Swagger UI

После запуска приложения интерактивная документация доступна по адресу:
http://localhost:8080/swagger/
Swagger UI позволяет просматривать и тестировать доступные API-методы непосредственно из браузера.

API

| Users  |                    |                               |      
|--------|--------------------|-------------------------------|
| Метод  | Endpoint           | Описание                      |   
| GET    | /api/v1/users/     | Получить список пользователей |   
| GET    | /api/v1/users/{id} | Получить пользователя по ID   |   
| POST   | /api/v1/users/     | Создать пользователя          |   
| PUT    | /api/v1/users/{id} | Обновить пользователя         |   
| DELETE | /api/v1/users/{id} | Удалить пользователя          |   

| Todos  |                    |                        |   
|--------|--------------------|------------------------|
| Метод  | Endpoint           | Описание               |   
| GET    | /api/v1/todos/     | Получить список задач  |   
| GET    | /api/v1/todos/{id} | Получить задачу по ID  |  
| POST   | /api/v1/todos/     | Создать задачу         |  
| PUT    | /api/v1/todos/{id} | Обновить задачу        |   
| DELETE | /api/v1/todos/{id} | Удалить задачу         |   
|        |                    |                        |   
|        |                    |                        |   
|        |                    |                        |   



## Примеры запросов и ответов
Создание пользователя
Запрос:
```bash
curl -X POST http://localhost:8080/api/v1/users/ \
  -H "Content-Type: application/json" \
  -d '{"email":"user@example.com","name":"Ivan"}'
```

Пример ответа:
```json
{
  "success": true,
  "data": {
    "id": 1,
    "email": "user@example.com",
    "name": "Ivan"
  }
}
```

Получение пользователя
Запрос:
```bash
curl http://localhost:8080/api/v1/users/1
```

Пример ответа:
```json
{
  "success": true,
  "data": {
    "id": 1,
    "email": "user@example.com",
    "name": "Ivan"
  }
}
```

Создание задачи
Запрос:
```bash
curl -X POST http://localhost:8080/api/v1/todos/ \
  -H "Content-Type: application/json" \
  -d '{"user_id":1,"title":"Изучить Swagger","description":"Добавить документацию API","done":false}'
```
Пример ответа:
```json
{
  "success": true,
  "data": {
    "id": 1,
    "user_id": 1,
    "title": "Изучить Swagger",
    "description": "Добавить документацию API",
    "done": false
  }
}
```

## Структура проекта
- ☑ main.go — точка входа приложения;
- ☑ internal/handlers/ — HTTP-обработчики пользователей и задач;
- ☑ internal/models/ — модели и структуры запросов/ответов;
- ☑ docs/ — сгенерированная Swagger-документация;
- ☑ README.md — документация проекта;
- ☑ go.mod — описание Go-модуля и основных зависимостей;
- ☑ go.sum — контрольные суммы зависимостей.
Структура директорий:
.
├── main.go
├── go.mod
├── go.sum
├── README.md
├── docs/
│   ├── docs.go
│   ├── swagger.json
│   └── swagger.yaml
└── internal/
    ├── handlers/
    │   ├── todo.go
    │   └── user.go
    └── models/
        ├── todo.go
        └── user.go

## Лицензия
Учебный проект. Используется в рамках образовательного задания.
Контакты

Автор проекта: [GitHub-профиль](https://github.com/moleculalisa)

