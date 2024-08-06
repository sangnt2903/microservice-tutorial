protoc:
	docker build -t protoc -f Dockerfile .

gen:
ifdef service
	@docker run -it --rm \
	-e SERVICE=$(service) \
	-v ./:/defs \
	-v ./generated/:/generated \
	protoc:latest ./entrypoint.sh

	@cp generated/$(service)/service.swagger.json ../services/$(service)/swagger.json
else
	@echo "service is required, make gen service=ecosystem/document"
endif