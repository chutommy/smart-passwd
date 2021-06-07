FROM golang:1.16 AS build-stage
LABEL maintainer="Tommy Chu <chutommy101@gmail.com>"
LABEL project="SmartPasswd"

# set enviroment
ENV GO111MODULE="on"
ENV GOARCH="amd64"
ENV GOOS="linux"
ENV CGO_ENABLED="1"

# prepare workspace
WORKDIR /build
COPY go.mod .
COPY go.sum .
COPY pkg/ pkg/
COPY main.go .
RUN go mod download
RUN go mod tidy

# build a binary
RUN go build -o main main.go


FROM golang:1.16 AS launch-stage

# set enviroment
ENV GO111MODULE="on"
ENV GOARCH="amd64"
ENV GOOS="linux"
ENV CGO_ENABLED="1"

# prepare workspace
WORKDIR /launch
COPY public/ public/
COPY data/ data/
COPY config.yaml .
COPY --from=build-stage /build/main .

# launch
EXPOSE 8080
ENTRYPOINT ["/launch/main"]