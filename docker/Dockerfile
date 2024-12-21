FROM golang:1.22-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

# Install CompileDaemon
RUN go install github.com/githubnemo/CompileDaemon@latest

COPY . .

EXPOSE 3005

# --polling flag == instructs CompileDaemon to periodically check for file changes instead of relying on filesystem events
ENTRYPOINT ["CompileDaemon", "--build=go build -o build/goapp", "--command=./build/goapp", "--polling"]