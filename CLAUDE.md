# CLAUDE.md — go-banking-system

## О проекте
Продакшен-грейд портфолио-проект для норвежского рынка труда.
Монорепозиторий с тремя микросервисами:
- `user-service` — регистрация, аутентификация, профили пользователей
- `payment-service` — переводы, транзакции, баланс
- `notification-service` — уведомления (email/push) о событиях

## Стек
- HTTP: Gin
- Межсервисное взаимодействие: gRPC
- Очереди: Kafka (franz-go)
- БД: PostgreSQL (pgx v5)
- Кэш: Redis
- Наблюдаемость: Prometheus + OpenTelemetry
- Инфраструктура: Kubernetes, Helm
- CI/CD: GitHub Actions
- Модули: go.work (workspace на все три сервиса)

## Учебный режим — ГЛАВНОЕ ПРАВИЛО
Я (Dany) учу Go параллельно с этим проектом, до этого писал на Java.
При любой задаче:
1. Сначала объясни ПОЧЕМУ так делается в Go (особенно если это отличается от Java-подхода)
2. Покажи 2-3 варианта, если они есть, с плюсами/минусами
3. Только потом пиши код — с комментариями, поясняющими нетривиальные места
4. Не пиши сразу "идиоматичный продакшен-код" без объяснений — цель не просто рабочий код, а понимание
5. Если я прошу просто "напиши X" без контекста — всё равно кратко поясни ключевое решение
6. Говорить о существовании альтернативных технологиях (актульаных на рынке)

Полезны параллели с Java (goroutines vs threads, interfaces vs Java interfaces,
error handling vs exceptions, defer vs try-finally и т.д.) — используй их
как точку опоры для объяснений.

## Статус и план
Смотри ROADMAP.md — 8 фаз:
1. Foundation (структура репо, go.work, базовая конфигурация)
2. user-service — core логика
3. user-service — gRPC
4. payment-service + Outbox pattern
5. notification-service
6. Observability (Prometheus, OpenTelemetry)
7. Kubernetes / Helm
8. CI/CD + финальная полировка портфолио

Текущая фаза: Фаза 2 — user-service, core-логика
Что сделано: доменная модель `User` (+ `ReconstituteUser`), ports (`UserRepository`, `PasswordHasher`), `usecase.RegisterUseCase`, адаптеры `adapters/postgres.UserRepository` и `adapters/hasher.BcryptHasher`; первая миграция (`000001_create_users_table`) проверена живьём (up/down на реальной БД в docker-compose); HTTP-слой на Gin (`register_handler`, `register_request`, `register_response`, `errors.go`) и `main.go` (сборка зависимостей: config → pgxpool → адаптеры → usecase → router) — эндпоинт `POST /register` поднят и проверен живьём через curl (201 Created, 409 при дубликате email, 400 при невалидном email/коротком пароле; в БД подтверждён bcrypt-хэш пароля)
Что дальше: базовые unit-тесты, use case аутентификации (потребует метод Verify/Compare в `ports.PasswordHasher`), вынести регистрацию роутов и сборку зависимостей из `main.go` в отдельный слой (`internal/app` + `adapters/http/router.go`) — задел на Фазу 3, где добавится ещё и gRPC-сервер

## Архитектурные соглашения
- Hexagonal architecture (по аналогии с тем, что я использую на текущей Java-работе)
- Outbox pattern для консистентности между БД и Kafka в payment-service
- Ошибки — через явный error handling, без паники в бизнес-логике
- [дополнять по ходу проекта конкретными решениями, которые примем]

## Стиль кода
- gofmt / golangci-lint обязательны перед коммитом
- Тесты — table-driven, где уместно
- [дополнять по мере выработки конвенций]

## Workflow
- Работаем поэтапно по ROADMAP.md, не перескакивая фазы
- Diff tool: auto (правки — через diff-вьювер GoLand)
- /compact — на границах логических кусков задачи
- /clear — при переходе между несвязанными задачами
- Важные решения фиксировать здесь, в CLAUDE.md, а не только в истории чата