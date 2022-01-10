DEPENDENCIES = yq
K := $(foreach exec,$(DEPENDENCIES),\
        $(if $(shell which $(exec)),some string,$(error "No $(exec) in PATH. Please install $(exec)")))
env_file = .env
env := $(shell yq r $(env_file) -p pv "**" | sed 's/ /|/g' | sed 's/:|/:/g')
env_path_value = $(subst :, ,$(env_element))
env_path = $(word 1,$(env_path_value))
env_value = $(subst |, ,$(word 2,$(env_path_value)))
env_variable = $(eval $(env_path) := $(env_value))
$(foreach env_element,$(env),$(env_variable))

# DOCKER TASKS
# Build the container
build-container: ## Build the container
	docker build --build-arg PORT=$(APP_PORT) --build-arg APP_NAME=$(APP_NAME) -t $(APP_NAME) .
run-container: ## Run container on port
	docker run --rm -d -p $(APP_PORT):$(APP_PORT) $(APP_NAME)
stop-container: ## Stop and remove a running container
	docker stop $(APP_NAME); docker rm $(APP_NAME)
.DEFAULT_GOAL := default

default: build

build:
	@go get -u github.com/go-swagger/go-swagger/cmd/swagger
	@swagger generate server -A provider-service -f schema/swagger.yml
	@swagger generate client -A provider-service -f schema/swagger.yml
	@go get -u -f ./...
	@go build ./cmd/provider-service-server

.PHONY: build