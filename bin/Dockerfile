FROM golang:1.16 AS build-stage
LABEL maintainer="Tommy Chu <tommychu2256@gmail.com>"
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
RUN GOOS=windows GOARCH=amd64 go build -o bin/windows/SmartPasswd-amd64.exe main.go
#RUN GOOS=windows GOARCH=386 go build -o bin/windows/SmartPasswd-386.exe main.go

RUN GOOS=darwin GOARCH=amd64 go build -o bin/darwin/SmartPasswd-amd64 main.go
RUN GOOS=darwin GOARCH=arm64 go build -o bin/darwin/SmartPasswd-arm64 main.go

RUN GOOS=linux GOARCH=amd64 go build -o bin/linux/SmartPasswd-amd64 main.go
#RUN GOOS=linux GOARCH=386 go build -o bin/linux/SmartPasswd-386 main.go
#RUN GOOS=linux GOARCH=arm go build -o bin/linux/SmartPasswd-arm main.go
RUN GOOS=linux GOARCH=arm64 go build -o bin/linux/SmartPasswd-arm64 main.go

RUN GOOS=android GOARCH=arm64 go build -o bin/android/SmartPasswd-arm64 main.go


FROM scratch AS export-stage
# store binaries
COPY --from=build-stage /build/bin /