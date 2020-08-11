FROM golang:alpine AS builder
LABEL maintainer="tommychu2256@gmail.com"

# set enviroment
ENV GO111MODULE="on"
ENV GOARCH="amd64"
ENV GOOS="linux"
ENV CGO_ENABLED="0"

# prepare workdir
WORKDIR /webapp
COPY go.mod .
COPY go.sum .
RUN go mod download

# build binary
COPY . .
RUN go mod tidy
RUN go build -o main .


FROM alpine:latest AS product

# set enviroment
ENV PORT=80

# prepare workdir
RUN apk update && apk add ca-certificates
WORKDIR /webapp
COPY --from=builder /webapp/config.yaml .
COPY --from=builder /webapp/main .

# run service
EXPOSE $PORT
ENTRYPOINT ["./main"]
