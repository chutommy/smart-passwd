.PHONY: build run

.PHONY: wordlist
wordlist:
	python3 data/prenasec.py data/raw/wordlist-1.txt data/parsed/wordlist-1.db

.PHONY: test-wordlist
test-wordlist:
	python3 data/prenasec.py data/raw/wordlist-1.txt pkg/data/test/wordlist.db

build:
	docker build -t smartpasswd --rm .

run:
	docker run --network=host -it smartpasswd
