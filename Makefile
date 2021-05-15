.PHONY: wordlist
wordlist:
	python3 data/prenasec.py data/raw/wordlist-1.txt data/parsed/wordlist-1.db

.PHONY: mock-wordlist
test-wordlist:
	python3 data/prenasec.py data/raw/wordlist-1.txt pkg/data/test/wordlist.db

.PHONY: test
test:
	go test -v ./...

.PHONY: build
build:
	GOOS=windows GOARCH=amd64 go build -o bin/windows/smart-passwd-amd64.exe main.go
	GOOS=windows GOARCH=386 go build -o bin/windows/smart-passwd-386.exe main.go
	GOOS=windows GOARCH=arm go build -o bin/windows/smart-passwd-arm.exe main.go

	GOOS=darwin GOARCH=amd64 go build -o bin/darwin/smart-passwd-amd64 main.go
	GOOS=darwin GOARCH=arm64 go build -o bin/darwin/smart-passwd-arm64 main.go

	GOOS=linux GOARCH=amd64 go build -o bin/linux/smart-passwd-amd64 main.go
	GOOS=linux GOARCH=386 go build -o bin/linux/smart-passwd-386 main.go
	GOOS=linux GOARCH=arm go build -o bin/linux/smart-passwd-arm main.go
	GOOS=linux GOARCH=arm64 go build -o bin/linux/smart-passwd-arm64 main.go

BG_IMAGES?=templates/assets/styles/images/background

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

FAVICONS?=templates/images

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