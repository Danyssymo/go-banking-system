# ROADMAP.md — go-banking-system

## Фаза 1: Foundation
- [x] Структура монорепозитория (user-service, payment-service, notification-service)
- [x] go.work в корне, связывающий все три модуля
- [x] Базовая конфигурация (env, config-файлы) — см. PROGRESS.md
- [x] Каркас CI (линтер, тесты) — без деплоя пока — см. PROGRESS.md
- Статус: готово (первый прогон на GitHub Actions — все 3 job'а зелёные)

## Фаза 2: user-service — core
- Регистрация, аутентификация пользователей
- Hexagonal architecture: domain / ports / adapters
- PostgreSQL через pgx v5
- Базовые unit-тесты
- Статус: не начато

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