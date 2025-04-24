FROM golang:1.23-alpine AS install
WORKDIR /app
COPY ./src/go.mod ./src/go.sum ./
RUN go mod download
RUN go install gotest.tools/gotestsum@latest

FROM install AS development
WORKDIR /app
COPY ./src ./
