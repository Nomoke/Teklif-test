# TeklifHub

## 🚀 **Запуск проекта**

Проект можно запустить **в Docker** или **локально**.

### Запуск в Docker

1. Запустите сервисы в контейнерах:

   ```bash
   make docker-up

2. Примените миграции базы данных:

   ```bash
   make migrate
   

# **Доступные команды**
- make run – Запуск приложения локально.
- make build – Сборка бинарника проекта.
- make migrate – Применение миграций для базы данных.
- make docker-up – Запуск проекта в Docker.
- make docker-down – Остановка и удаление контейнеров.


## 🌐 **API**

После выполнения всех шагов **API** будет доступно для использования.

**Swagger UI** доступен по следующему адресу:

[**Swagger UI**](http://localhost:8080/swagger/index.html#/offers/post_offers)
