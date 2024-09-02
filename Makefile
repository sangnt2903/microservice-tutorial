registry=sangnguyenitp

GIT_COMMIT := $(shell git rev-parse --short HEAD)

ifndef env
	env = dev
endif

port=$(shell ./port.sh $(service) $(env))

service_name=$(shell echo $(service) | sed 's/\//-/g')

build:
ifdef service
	@echo "Building service: $(service) port: $(port)"
	@cd $(service) &&  CGO_ENABLED=0 GOOS=linux go build -o bin
	@docker build -t $(service_name) -f Dockerfile --build-arg SERVICE=$(service) --build-arg ENV=$(env) --build-arg PORT=$(port) .
endif

push: build
ifdef service
	@echo "Pushing image to registry: $(registry)/sai/$(service_name):$(GIT_COMMIT)"
	@docker tag $(service_name) $(registry)/sai:$(GIT_COMMIT)
	@docker push $(registry)/sai:$(GIT_COMMIT)

ifeq ($(env),dev)
	@echo "Pushing image to registry: $(registry)/sai/$(service_name):latest"
	@docker tag $(service_name) $(registry)/sai:latest
	@docker push $(registry)/sai:latest
endif
else
	@echo "Error: 'service' variable is not set"
endif


