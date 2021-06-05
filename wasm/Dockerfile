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
COPY wasm/ wasm/
RUN go mod download
RUN go mod tidy

# build binary
RUN GOOS=js GOARCH=wasm go build -o bin/sqlite-main.wasm wasm/sqlite/main.go
RUN GOOS=js GOARCH=wasm go build -o bin/mongo-main.wasm wasm/mongo/main.go
# get the js support file
RUN mkdir scripts
RUN cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" scripts/


FROM scratch AS export-stage
# store binaries
COPY --from=build-stage /build/bin/sqlite-main.wasm /wasm/
COPY --from=build-stage /build/bin/mongo-main.wasm /wasm/
COPY --from=build-stage /build/scripts/wasm_exec.js /scripts/
