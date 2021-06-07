.PHONY: wordlist-mongo
wordlist-mongo:
	python3 data/prenasec_mongo.py data/raw/wordlist-1.txt $(SMART_PASSWD_MONGODB_URI)

.PHONY: wordlist-sql
wordlist-sql:
	touch data/parsed/wordlist-1.db
	python3 data/prenasec_sqlite.py data/raw/wordlist-1.txt data/parsed/wordlist-1.db

.PHONY: test-wordlist-sql
test-wordlist-sql:
	python3 data/prenasec_sqlite.py data/raw/wordlist-1.txt pkg/data/test/wordlist.db
	python3 data/prenasec_sqlite.py data/raw/wordlist-1.txt pkg/engine/test/wordlist.db
	python3 data/prenasec_sqlite.py data/raw/wordlist-1.txt pkg/server/test/wordlist.db
	python3 data/prenasec_sqlite.py data/raw/wordlist-1.txt public/db/wordlist.db

.PHONY: test
test:
	GOOS=linux GOARCH=amd64 go test -v ./pkg/...
	GOOS=js GOARCH=wasm go test -v ./wasm/...

.PHONY: build
build:
	DOCKER_BUILDKIT=1 docker build --target export-stage --output bin --file bin/Dockerfile .

.PHONY: wasm
wasm:
	DOCKER_BUILDKIT=1 docker build --target export-stage --output public --file wasm/Dockerfile .

.PHONY: docker
docker:
	DOCKER_BUILDKIT=1 docker build --file Dockerfile -t smart-passwd .
	docker run -it -p 8080:8080 smart-passwd

.PHONY: npmi
npmi:
	npm install --prefix ./public --only=production

.PHONY: npmu
npmu:
	npm update --prefix ./public

BG_IMAGES?=public/assets/styles/images/background

.PHONY: bg-build
bg-build:
	triangula run --image $(BG_IMAGES)/bg-orig.jpg \
				  --output $(BG_IMAGES)/bg \
				  --points 4000

.PHONY: bg-render
bg-render:
	triangula render --input $(BG_IMAGES)/bg.json \
					 --image $(BG_IMAGES)/bg-orig.jpg \
					 --output $(BG_IMAGES)/bg
	inkscape --export-filename $(BG_IMAGES)/bg \
			 --export-type=png -w 2560 -h 1600 \
			 $(BG_IMAGES)/bg.svg

FAVICONS?=public/images

.PHONY: favicon
favicon:
	inkscape --export-filename $(FAVICONS)/16 \
			 --export-type=png -w 16 -h 16 \
			 $(FAVICONS)/favicon.svg
	inkscape --export-filename $(FAVICONS)/32 \
			 --export-type=png -w 32 -h 32 \
			 $(FAVICONS)/favicon.svg
	inkscape --export-filename $(FAVICONS)/48 \
			 --export-type=png -w 48 -h 48 \
			 $(FAVICONS)/favicon.svg
	convert $(FAVICONS)/16.png $(FAVICONS)/32.png $(FAVICONS)/48.png $(FAVICONS)/favicon.ico