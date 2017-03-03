.DEFAULT_GOAL := help

test: ## Run all unit tests
	@go test -race -v $(shell go list ./... | grep -v /vendor/)

lint: ## Lint all files, golint is concerned with coding style: gives suggestions
	go list ./... | grep -v /vendor/ | xargs -L1 golint -set_exit_status

vet: ## Run the go vet tool, Vet examines Go source code and reports suspicious constructs
	go et $(shell go list ./... | grep -v /vendor/)

help: ## Display this help message
	@cat $(MAKEFILE_LIST) | grep -e "^[a-zA-Z_\-]*: *.*## *" | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.SILENT: build test lint vet clean docker-build docker-push help
