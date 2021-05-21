FROM golang:1.16 AS build-stage
LABEL maintainer="Tommy Chu <chutommy101@gmail.com>"
LABEL project="SmartPasswd"

# set enviroment
ENV GO111MODULE="on"
ENV GOARCH="amd64"
ENV GOOS="linux"
ENV CGO_ENABLED="0"

# prepare workspace
WORKDIR /build
COPY go.mod .
COPY go.sum .
COPY pkg/ pkg/
COPY main.go .
RUN go mod download
RUN go mod tidy

# build binaries
RUN GOOS=windows GOARCH=amd64 go build -o bin/windows/smart-passwd-amd64.exe main.go
RUN GOOS=windows GOARCH=386 go build -o bin/windows/smart-passwd-386.exe main.go

RUN GOOS=darwin GOARCH=amd64 go build -o bin/darwin/smart-passwd-amd64 main.go
RUN GOOS=darwin GOARCH=arm64 go build -o bin/darwin/smart-passwd-arm64 main.go

RUN GOOS=linux GOARCH=amd64 go build -o bin/linux/smart-passwd-amd64 main.go
RUN GOOS=linux GOARCH=386 go build -o bin/linux/smart-passwd-386 main.go
RUN GOOS=linux GOARCH=arm go build -o bin/linux/smart-passwd-arm main.go
RUN GOOS=linux GOARCH=arm64 go build -o bin/linux/smart-passwd-arm64 main.go


FROM scratch AS export-stage
# store binaries
COPY --from=build-stage /build/bin /