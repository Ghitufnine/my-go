# MyGo Clean Architecture Backend

Production-style backend service built with Golang using Clean Architecture, PostgreSQL, MongoDB, Redis, RabbitMQ, and Grafana stack.

This project demonstrates a scalable backend structure with:

- Clean Architecture
- Event-driven logging
- Redis caching
- JWT authentication
- RabbitMQ messaging
- MongoDB transaction logs
- Docker compose infrastructure
- Loki + Promtail + Grafana logging

---

# Architecture

This project follows Clean Architecture.


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

# Project Structure


cmd/
api/
main.go
container/
container.go
container_server.go

internal/
domain/entity
repository
usecase
handler/http
routes
middleware

infrastructure/
postgres
mongo
redis
redis_cache
rabbitmq
logger
config

pkg/
jwt
utils

docker/
docker-compose.yml
loki-config.yml
promtail-config.yml

Dockerfile
.env


---

# Tech Stack

- Golang + Fiber
- PostgreSQL (main DB)
- MongoDB (transaction logs)
- Redis (cache)
- RabbitMQ (event bus)
- Grafana + Loki + Promtail (logging)
- Docker Compose

---

# Features

## Authentication

- Register
- Login (JWT)
- Refresh token stored in DB
- Logout (invalidate refresh token)
- JWT middleware

## Categories

- Create
- List
- Update
- Delete

## Items

- Create
- List
- Get detail
- Update
- Delete

Rules:

- Item belongs to Category
- FK enforced
- Cache for list endpoints

## Redis Cache

Cached:


GET /categories
GET /items


TTL: 60s

Invalidated on:

- create
- update
- delete

## RabbitMQ Events

Published:


category.created
category.updated
category.deleted

item.created
item.updated
item.deleted


## Mongo Logs

Consumer writes to:


clean_arch_logs.transaction_logs


Fields:


id
topic
payload
created_at


## Logging

Structured JSON logs using zap.

Sent to:


stdout → promtail → loki → grafana


---

# Requirements

Install:

- Go 1.22+
- Docker
- Docker Compose
- Postman (optional)
- MongoDB Compass (optional)

---

# Environment

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

# Running with Docker


cd docker
docker compose up -d


Services:


postgres
mongo
redis
rabbitmq
grafana
loki
promtail
app


Check:


http://localhost:8080

http://localhost:15672

http://localhost:3000


Grafana login:


admin / admin


---

# Running without Docker

Start manually:


postgres
mongo
redis
rabbitmq


Then:


go run cmd/api/main.go


---

# Database Migration

Run manually in PostgreSQL:


CREATE TABLE users (
id UUID PRIMARY KEY,
email TEXT UNIQUE NOT NULL,
password TEXT NOT NULL,
created_at TIMESTAMP
);

CREATE TABLE refresh_tokens (
id UUID PRIMARY KEY,
user_id UUID REFERENCES users(id),
token TEXT,
expires_at TIMESTAMP
);

CREATE TABLE categories (
id UUID PRIMARY KEY,
name TEXT,
created_at TIMESTAMP
);

CREATE TABLE items (
id UUID PRIMARY KEY,
category_id UUID REFERENCES categories(id),
name TEXT,
price NUMERIC,
created_at TIMESTAMP
);


MongoDB auto creates collection.

Redis auto creates keys.

RabbitMQ auto creates exchange.

---

# API Endpoints

## Auth

POST /api/auth/register
POST /api/auth/login
POST /api/auth/logout

## Categories

POST /api/categories
GET /api/categories
PUT /api/categories/:id
DELETE /api/categories/:id

## Items

POST /api/items
GET /api/items
GET /api/items/:id
PUT /api/items/:id
DELETE /api/items/:id

Authorization:


Authorization: Bearer TOKEN


---

# Event Flow


Handler
→ Usecase
→ Publish event
→ RabbitMQ
→ Consumer
→ MongoDB logs


---

# Cache Flow


Usecase
→ CacheDecorator
→ Redis
→ Postgres


---

# Logging Flow


Zap → stdout → promtail → loki → grafana


---

# Notes

This project is for learning and demonstration of:

- Clean Architecture
- Distributed backend
- Event-driven design
- Production logging
- Docker infra