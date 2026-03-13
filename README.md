# Go Clean Architecture Backend Service

Production-style backend service built with **Golang** using Clean Architecture, PostgreSQL, MongoDB, Redis, RabbitMQ, and Grafana stack.

This project demonstrates how to build a **scalable, event-driven, production-ready backend service** with proper architecture, caching, messaging, and logging.

This repository is part of my backend engineering portfolio.

---

## 🚀 Features

- Clean Architecture
- REST API with Fiber
- JWT Authentication
- Redis caching
- RabbitMQ event messaging
- MongoDB transaction logs
- PostgreSQL main database
- Docker Compose infrastructure
- Loki + Promtail + Grafana logging
- Event-driven backend flow
- Cache invalidation strategy

---

## 🧠 Architecture

This project follows **Clean Architecture** principles.


Handler → Usecase → Repository(interface) → Infrastructure
↑
Middleware / Router


Rules:

- Handler contains HTTP logic only
- Usecase contains business logic
- Repository is interface only
- Infrastructure implements repository
- Entity has no framework dependency

---

## 📂 Project Structure


cmd/
api/

internal/
domain/
entity/
repository/
usecase/
handler/http/
routes/
middleware/

infrastructure/
postgres/
mongo/
redis/
redis_cache/
rabbitmq/
logger/
config/

pkg/
jwt/
utils/

docker/
docker-compose.yml
loki-config.yml
promtail-config.yml

Dockerfile
.env


Goal:

- maintainable
- scalable
- production ready
- easy to extend

---

## 🛠 Tech Stack

- Golang + Fiber
- PostgreSQL (main DB)
- MongoDB (transaction logs)
- Redis (cache)
- RabbitMQ (event bus)
- Grafana + Loki + Promtail (logging)
- Docker Compose

---

## 🔐 Authentication

- Register
- Login (JWT)
- Refresh token stored in DB
- Logout (invalidate refresh token)
- JWT middleware

---

## 📦 Categories API

- Create
- List
- Update
- Delete

---

## 📦 Items API

- Create
- List
- Detail
- Update
- Delete

Rules:

- Item belongs to Category
- Foreign key enforced
- Cache used for list endpoints

---

## ⚡ Redis Cache

Cached endpoints:

- GET /categories
- GET /items

TTL: 60s

Invalidated on:

- create
- update
- delete

---

## 📡 RabbitMQ Events

Published events:

- category.created
- category.updated
- category.deleted
- item.created
- item.updated
- item.deleted

---

## 🗄 MongoDB Logs

Consumer writes logs to:


clean_arch_logs.transaction_logs


Fields:

- id
- topic
- payload
- created_at

---

## 📊 Logging Stack

Structured logging using zap.

Flow:


App → stdout → promtail → loki → grafana


---

## ⚙️ Requirements

Install:

- Go 1.22+
- Docker
- Docker Compose

Optional:

- Postman
- MongoDB Compass

---

## 🔧 Environment

Create `.env`


APP_PORT=8080

POSTGRES_HOST=localhost
POSTGRES_PORT=5432
POSTGRES_USER=postgres
POSTGRES_PASSWORD=postgres
POSTGRES_DB=clean_arch

MONGO_URI=mongodb://admin:admin123@localhost:27017/?authSource=admin
MONGO_DB=clean_arch_logs

REDIS_ADDR=localhost:6379
REDIS_PASSWORD=
REDIS_DB=0

RABBITMQ_URL=amqp://guest:guest@localhost:5672/
RABBITMQ_EXCHANGE=events

JWT_SECRET=supersecret


When using docker, change host to service name.

---

## 🐳 Run with Docker


cd docker
docker compose up -d


Services:

- postgres
- mongo
- redis
- rabbitmq
- grafana
- loki
- promtail
- app

Check:


http://localhost:8080

http://localhost:15672

http://localhost:3000


Grafana login:


admin / admin


---

## ▶ Run without Docker

Start manually:

- postgres
- mongo
- redis
- rabbitmq

Then:


go run cmd/api/main.go


---

## 🔄 Event Flow


Handler
→ Usecase
→ Publish event
→ RabbitMQ
→ Consumer
→ MongoDB logs


---

## 💾 Cache Flow


Usecase
→ Cache decorator
→ Redis
→ PostgreSQL


---

## 📈 Logging Flow


Zap
→ stdout
→ promtail
→ loki
→ grafana


---

## 🎯 Purpose

This project demonstrates backend engineering skills including:

- Clean Architecture
- Event-driven backend
- Redis caching strategy
- Message queue integration
- Structured logging
- Docker infrastructure
- Production-style backend design

This repository is used as part of my backend engineering portfolio for global / remote backend roles.