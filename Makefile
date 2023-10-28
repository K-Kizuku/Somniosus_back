DOCKER_COMPOSE := docker compose
APPS := account 
SHARED_NETWORK := microservices-shared

.PHONY: setup-tools c-service c-proto adr commit evans lint psql migrate/status migrate/new migrate/up buf-gen

export GOPRIVATE=github.com/K-Kizuku/Somniosus_back

setup-tools:
	@if [ -z `command -v air` ]; then go install github.com/cosmtrek/air@latest ; fi
	@if ! [ -x ${GOPATH}/bin/buf ]; then go install github.com/bufbuild/buf/cmd/buf@latest  ; fi

##### scaffing

c-service:
	chmod +x ./scripts/create_service.sh
	./scripts/create_service.sh ${DIR}

ADR_COUNT:=$(shell find docs/adr -type f | wc -l | tr -d ' ') 
adr:
	npx scaffdog generate ADR --output 'docs/adr' --answer 'number:${ADR_COUNT}'


##### git

commit:
	npx git-cz

pull-sub:
	git submodule update --remote


##### tools

evans:
	evans --proto ./protobuf/**/* --port 8080

lint:
	@if ! [ -x /usr/local/bin/golangci-lint ]; then \
	    brew install golangci/tap/golangci-lint ;\
	fi
	golangci-lint run --concurrency 2 ./...


###### DB

POSTGRES_USER := root
DATABASE      := treasure_app
psql:
	$(DOCKER_COMPOSE) exec db psql -U $(POSTGRES_USER) -d $(DATABASE)

GOOSE_DRIVER   := postgres
GOOSE_DBSTRING ?= host=db user=root dbname=treasure_app password=p@ssword sslmode=disable
migrate/status:
	$(DOCKER_COMPOSE) run --rm migration status

NAME?=
SERVICE?=
migrate/new:
	echo '-- +goose Up' > apps/${SERVICE}/app/infra/db/migrations/$(shell ls apps/${SERVICE}/app/infra/db/migrations | awk -F"_*.sql" 'BEGIN {max=0} {split($$1, a, "_"); if(a[1]>max){max = a[1]}}END{print max+1}')_${NAME}.sql

migrate/up:
	$(DOCKER_COMPOSE) run --rm migration up

##### buf

buf-gen: pull-sub
	sh ./scripts/protoc_generate.sh     


##### docker 
.PHONY: network
network: ## Gatewayと通信用のdocker network作成
	if ! docker network ls | grep ${SHARED_NETWORK} > /dev/null; then docker network create ${SHARED_NETWORK}; fi

.PHONY: network-clean
network-clean: ## Gatewayと通信用のdocker network削除
	if docker network ls | grep ${SHARED_NETWORK} > /dev/null; then docker network rm ${SHARED_NETWORK}; fi

# build-<appname>の生成
.PHONY: $(addprefix build-,$(APPS))
$(foreach app,$(APPS), \
	$(eval \
	build-$(app):;\
		CGO_ENABLED=0 go build -o $(OUTPUT)/ ./apps/$(app)/cmd/$(app) \
	) \
)
# run-<appname>の生成
.PHONY: $(addprefix run-,$(APPS))
$(foreach app,$(APPS), \
	$(eval \
	run-$(app):;\
		$(OUTPUT)/$(app) \
	) \
)
# watch-<appname>の生成
.PHONY: $(addprefix watch-,$(APPS))
$(foreach app,$(APPS), \
	$(eval \
	watch-$(app):;\
		@air --build.cmd "make build-$(app)" \
			--build.bin "$(OUTPUT)/$(app)" \
			--build.include_dir=["apps/$(app)/...", "lib/..."] \
			--root ./apps/$(app) \
			-c .air.toml \
	) \
)