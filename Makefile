NAME = vincentvanderweele/wwweeklies-presentations
PKGS = $(shell go list ./... | grep -v /vendor/)

update-deps:
	godep save $(PKGS)

generate-server:
	go-server-generator ./swagger.yaml

build:
	make generate-server && docker build -t $(NAME) .

run:
	docker run --rm -it \
	-p 9999:9999 \
	-v $(shell pwd):/go/src/github.com/$(NAME) \
	$(NAME) \
	fresh

run-swagger-ui:
	docker run --rm -it \
	-p 80:8080 \
	-e "API_URL=http://localhost:9999/swagger" \
	swaggerapi/swagger-ui

.PHONY: update-deps generate-server build run run-swagger-ui
