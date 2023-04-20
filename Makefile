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

# make run service=xxx-service
.PHONY: run
run:
	go run ./app/${service}/cmd/main/... -conf=./app/${service}/configs

# make dev service=xxx-service
.PHONY: dev
dev:
	go run ./app/${service}/cmd/main/... -conf-consul=./app/${service}/configs/consul

# make migrate service=xxx-service
.PHONY: migrate
migrate:
	go run ./app/${service}/cmd/migration/... -conf=./app/${service}/configs

# make api service=xxx-service
.PHONY: api
api:
	go run ./cmd/proto/... -path=./api/${service}

# make config service=xxx-service
.PHONY: config
config:
	go run ./cmd/proto/... -path=./app/${service}/internal/conf

# make consul-config service=xxx-service
.PHONY: consul-config
consul-config:
	go run ./cmd/consul-config/... -conf=./app/${service}/configs/develop

# ping 请注意端口号
# 注意，请先运行服务
# make run service=ping-service
.PHONY: ping
ping:
	curl http://127.0.0.1:11001/api/v1/ping/hello && \
    echo "\n" && \
    curl http://127.0.0.1:11001/api/v1/ping/error && \
    echo "\n"
