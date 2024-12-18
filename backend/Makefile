# Makefile for PostgreSQL Docker deployment

# 默认数据库配置
POSTGRES_USER=watamin
POSTGRES_PASSWORD=c
POSTGRES_DB=bank
POSTGRES_CONTAINER=bank-postgres
PG_PORT=5432

# Docker image for PostgreSQL
POSTGRES_IMAGE=postgres:12-alpine

# 启动 PostgreSQL 容器
.PHONY: start-postgres
start-postgres:
	sudo docker run --name $(POSTGRES_CONTAINER) \
		-e POSTGRES_USER=$(POSTGRES_USER) \
		-e POSTGRES_PASSWORD=$(POSTGRES_PASSWORD) \
		-e POSTGRES_DB=$(POSTGRES_DB) \
		-p $(PG_PORT):5432 \
		-v pgdata:/var/lib/postgresql/data \
		-d $(POSTGRES_IMAGE)

# 停止 PostgreSQL 容器
.PHONY: stop-postgres
stop-postgres:
	sudo docker stop $(POSTGRES_CONTAINER) && sudo docker rm $(POSTGRES_CONTAINER)

# 清理 Docker 卷
.PHONY: clean-volumes
clean-volumes:
	sudo docker volume rm pgdata

# 查看 PostgreSQL 容器状态
.PHONY: status
status:
	sudo docker ps -f "name=$(POSTGRES_CONTAINER)"

# 连接到 PostgreSQL 容器
.PHONY: connect
connect:
	sudo docker exec -it $(POSTGRES_CONTAINER) psql -U $(POSTGRES_USER) -d $(POSTGRES_DB)

# 查看日志
.PHONY: logs
logs:
	sudo docker logs -f $(POSTGRES_CONTAINER)

