.DEFAULT_GOAL:=help

TAG=dev

.PHONY: build-api
build-api:  ## build the backend/api project for Raspberry Pi
	env GOOS=linux GOARCH=arm go build -a -o ../marcel ../cmd/marcel && \
	docker image build -t zenika/marcel:$(TAG) ../

.PHONY: build-front
build-front: ## build the front project
	cd ../pkg/frontend && \
	yarn build && \
	docker image build -t zenika/marcel-frontend:$(TAG) .

.PHONY: build-back
build-back: ## build the backoffice project
	cd ../pkg/backoffice && \
	yarn && yarn build && \
	docker image build -t zenika/marcel-backoffice:$(TAG) .

.PHONY: build
build: build-front build-back build-api ## build all projects for Raspberry Pi

.PHONY: help
help:
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage: \033[36m\033[0m\n"} /^[$$()% a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)
