

.PHONY: setup-tools c-service adr commit

setup-tools:
	@if [ -z `command -v air` ]; then go install github.com/cosmtrek/air@latest ; fi

c-service:
	chmod +x ./scripts/create_service.sh
	./scripts/create_service.sh ${DIR}

ADR_COUNT:=$(shell find docs/adr -type f | wc -l | tr -d ' ') 
adr:
	npx scaffdog generate ADR --output 'docs/adr' --answer 'number:${ADR_COUNT}'

commit:
	npx git-cz

