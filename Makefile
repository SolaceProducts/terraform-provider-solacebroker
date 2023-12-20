#make
PKG_LIST := $(shell go list ./... | grep -v /vendor/)


.PHONY:
dep: ## Get the dependencies
	@go mod tidy

.PHONY:
vet: ## Run go vet
	@go vet ${PKG_LIST}

.PHONY:
fmt: ## Run gofmt
	@gofmt -w -l .

.PHONY:
test: ## Run unit tests
	@go test -short ${PKG_LIST}

.PHONY:
test-coverage: ## Run tests with coverage
	@mkdir -p reports
	@go test -short -coverprofile reports/cover.out ${PKG_LIST}
	@go tool cover -html reports/cover.out -o reports/cover.html

.PHONY: testacc
testacc: ## Run acceptance tests
	@TF_ACC=1 go test ./... -v $(TESTARGS) -timeout 120m

.PHONY:
generate-docs: dep ## Build the binary file
	@go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs generate --provider-name solacebroker

.PHONY:
build: dep ## Build the binary file
	@go build -a -ldflags '-s -w -extldflags "-static"' -o terraform-provider-solacebroker

.PHONY:
install: dep ## Install the provider for dev use
	@go install -a

.PHONY:
clean: ## Remove previous build
	@rm -f reports/cover.html reports/cover.out terraform-provider-solacebroker

.PHONY:
help: ## Display this help screen
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
