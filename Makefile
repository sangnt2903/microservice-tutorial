
ifndef env
	env = dev
endif

port=$(shell ./port.sh $(service) $(env))

service:
ifdef service
	@echo "Building service: $(service) port: $(port)"
	@cd $(service) &&  CGO_ENABLED=0 GOOS=linux go build -o bin
	@docker build -t $(service) -f Dockerfile --build-arg SERVICE=$(service) --build-arg ENV=$(env) --build-arg PORT=$(port) .
endif

.PHONY: service

