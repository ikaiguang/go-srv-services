# 版本信息
GO_VERSION:=$(shell go env GOPATH)
GOPATH:=$(shell go env GOPATH)
GIT_BRANCH=$(shell git branch | grep '*' | awk '{print $$2}')
GIT_VERSION=$(shell git describe --tags --always)

.PHONY: info
info:
	@echo "==> GO_VERSION: $(GO_VERSION)"
	@echo "==> GOPATH: $(GOPATH)"
	@echo "==> GIT_BRANCH: $(GIT_BRANCH)"
	@echo "==> GIT_VERSION: $(GIT_VERSION)"

# make run app=xxx
.PHONY: run
run:
	go run ./app/${app}-service/cmd/main/... -conf=./app/${app}-service/configs

# make migrate app=xxx
.PHONY: migrate
migrate:
	go run ./app/${app}-service/cmd/migration/... -conf=./app/${app}-service/configs

# make api path=xxx
.PHONY: api
api:
	go run ./cmd/proto/... -path=./api/${path}

# make config service=xxx
.PHONY: config
config:
	go run ./cmd/proto/... -path=./app/${service}-service/internal/conf

.PHONY: ping
ping:
	curl http://127.0.0.1:8081/api/v1/ping/hello && \
    echo "\n" && \
    curl http://127.0.0.1:8081/api/v1/ping/error && \
    echo "\n"
