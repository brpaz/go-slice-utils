.SILENT: ;               # no need for @
.ONESHELL: ;             # recipes execute in same shell
.NOTPARALLEL: ;          # wait for this target to finish
.EXPORT_ALL_VARIABLES: ; # send all vars to shell

.PHONY: all # All targets are accessible for user
.DEFAULT: help # Running Make will run the help target

help: ## Show Help
    grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

format: ## Run Unit tests
	@gofmt -l -w .

test: ## Run Unit tests
	go test -v ./...

test-cover: ## Run tests with test coverage
	go test -v ./... -cover	-coverprofile=cover.out

test-cover-html: test-cover ## Runs test coverage and export result to htmml
	go tool cover -html=cover.out

release: ## Creates a new release
	git tag $(TAG)
	git push --tags

docs: ## Launcher godoc server
	godoc -http=:6060
