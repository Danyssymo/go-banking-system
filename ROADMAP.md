# ROADMAP.md — go-banking-system

## Фаза 1: Foundation
- [x] Структура монорепозитория (user-service, payment-service, notification-service)
- [x] go.work в корне, связывающий все три модуля
- [x] Базовая конфигурация (env, config-файлы) — см. PROGRESS.md
- [x] Каркас CI (линтер, тесты) — без деплоя пока — см. PROGRESS.md
- Статус: готово (первый прогон на GitHub Actions — все 3 job'а зелёные)

## Фаза 2: user-service — core
- [~] Регистрация, аутентификация пользователей (регистрация реализована полностью — domain → ports → usecase → adapters → HTTP-эндпоинт `POST /register`, проверена живьём через curl: 201/409/400; аутентификация — не начата)
- [~] Hexagonal architecture: domain / ports / adapters
  - [x] Доменная модель `User` (приватные поля + `NewUser` с валидацией + `ReconstituteUser` для восстановления из БД)
  - [x] Геттеры для полей `User`
  - [x] Ports (интерфейсы: `UserRepository`, `PasswordHasher`)
  - [~] Use cases (Register — готов; Authenticate — не начат)
  - [x] Adapters (`postgres.UserRepository`, `hasher.BcryptHasher`, HTTP-адаптер `register_handler`/`register_request`/`register_response`/`errors.go` — все готовы и проверены живьём)
- [x] PostgreSQL через pgx v5 (адаптер написан; миграция `000001_create_users_table` проверена живьём — `up`/`down` отработали на реальной БД в docker-compose)
- [x] `main.go` — сборка зависимостей (config → pgxpool → адаптеры → usecase → Gin router), запущен и проверен живьём
- [ ] Базовые unit-тесты
- Статус: в процессе — регистрация работает end-to-end и проверена вручную (curl + прямая проверка БД: bcrypt-хэш, 201/409/400 на всех кейсах); осталось: unit-тесты, use case аутентификации, вынести router/сборку зависимостей из `main.go` в отдельный слой (`internal/app` + `router.go`) перед тем как в Фазе 3 добавится ещё и gRPC-сервер

## Фаза 3: user-service — gRPC
- gRPC API поверх core-логики (для межсервисного взаимодействия)
- Protobuf-контракты
- Статус: не начато

## Фаза 4: payment-service + Outbox
- Переводы, транзакции, баланс
- Outbox pattern для консистентности БД ↔ Kafka
- Kafka (franz-go) — публикация событий о транзакциях
- Статус: не начато

## Фаза 5: notification-service
- Подписка на Kafka-события (payment, user)
- Отправка уведомлений (email/push — можно замокать внешний провайдер)
- Статус: не начато

## Фаза 6: Observability
- Prometheus метрики по всем сервисам
- OpenTelemetry — трейсинг между сервисами
- Дашборды (Grafana — опционально)
- Статус: не начато

## Фаза 7: Kubernetes / Helm
- Helm-чарты для всех трёх сервисов
- Локальный деплой (minikube/kind) для проверки
- Статус: не начато

## Фаза 8: CI/CD + полировка портфолио
- GitHub Actions: тесты, линт, сборка образов
- README с архитектурной схемой
- Финальная проверка кода на идиоматичность
- Статус: не начато