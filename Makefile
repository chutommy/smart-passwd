.PHONY: build run

build:
	docker-compose -f docker-compose.yml -p smart-passwd build

run:
	docker-compose -f docker-compose.yml -p smart-passwd up
