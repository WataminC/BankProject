# Makefile for PostgreSQL and Redis Docker deployment

# 默认数据库配置
POSTGRES_USER=watamin
POSTGRES_PASSWORD=c
POSTGRES_DB=bank
POSTGRES_CONTAINER=bank-postgres
PG_PORT=5432

# Docker image for PostgreSQL
POSTGRES_IMAGE=postgres:12-alpine

# Redis 配置
REDIS_CONTAINER=bank-redis
REDIS_PORT=6379
REDIS_IMAGE=redis:7-alpine

# 启动 PostgreSQL 容器
.PHONY: start-postgres
start-postgres:
	docker run --name $(POSTGRES_CONTAINER) \
		-e POSTGRES_USER=$(POSTGRES_USER) \
		-e POSTGRES_PASSWORD=$(POSTGRES_PASSWORD) \
		-e POSTGRES_DB=$(POSTGRES_DB) \
		-p $(PG_PORT):5432 \
		-v pgdata:/var/lib/postgresql/data \
		-d $(POSTGRES_IMAGE)

# 停止 PostgreSQL 容器
.PHONY: stop-postgres
stop-postgres:
	docker stop $(POSTGRES_CONTAINER)

# 启动 Redis 容器
.PHONY: start-redis
start-redis:
	docker run --name $(REDIS_CONTAINER) \
		-p $(REDIS_PORT):6379 \
		-d $(REDIS_IMAGE)

# 停止 Redis 容器
.PHONY: stop-redis
stop-redis:
	docker stop $(REDIS_CONTAINER)

# 启动已经创建的 PostgreSQL 和 Redis 容器
.PHONY: run
run:
	docker start $(POSTGRES_CONTAINER) $(REDIS_CONTAINER)

# 清理 Docker 卷
.PHONY: clean-volumes
clean-volumes:
	docker volume rm pgdata

# 查看 PostgreSQL 容器状态
.PHONY: status-postgres
status-postgres:
	docker ps -f "name=$(POSTGRES_CONTAINER)"

# 查看 Redis 容器状态
.PHONY: status-redis
status-redis:
	docker ps -f "name=$(REDIS_CONTAINER)"

# 连接到 PostgreSQL 容器
.PHONY: connect-postgres
connect-postgres:
	docker exec -it $(POSTGRES_CONTAINER) psql -U $(POSTGRES_USER) -d $(POSTGRES_DB)

# 查看日志
.PHONY: logs-postgres
logs-postgres:
	docker logs -f $(POSTGRES_CONTAINER)

# 查看 Redis 日志
.PHONY: logs-redis
logs-redis:
	docker logs -f $(REDIS_CONTAINER)