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
	GOOS=windows go build -o smart-passwd-windows.exe main.go
	GOOS=darwin go build -o smart-passwd-darwin main.go
	GOOS=linux go build -o smart-passwd-linux main.go

	GOOS=windows GOARCH=amd64 go build -o bin/windows/smart-passwd-amd64.exe main.go
	GOOS=windows GOARCH=386 go build -o bin/windows/smart-passwd-386.exe main.go
	GOOS=windows GOARCH=arm go build -o bin/windows/smart-passwd-arm.exe main.go

	GOOS=darwin GOARCH=amd64 go build -o bin/darwin/smart-passwd-amd64 main.go
	GOOS=darwin GOARCH=arm64 go build -o bin/darwin/smart-passwd-arm64 main.go

	GOOS=linux GOARCH=amd64 go build -o bin/linux/smart-passwd-amd64 main.go
	GOOS=linux GOARCH=386 go build -o bin/linux/smart-passwd-386 main.go
	GOOS=linux GOARCH=arm go build -o bin/linux/smart-passwd-arm main.go
	GOOS=linux GOARCH=arm64 go build -o bin/linux/smart-passwd-arm64 main.go