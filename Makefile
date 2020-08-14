.PHONY: build run

build:
	docker build -t smartpasswd --rm .

run:
	docker run --network=host -it smartpasswd
