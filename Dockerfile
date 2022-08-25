# FROM node:alpine AS frontend
# COPY frontend/tsconfig.json frontend/package.json frontend/package-lock.json /src/
# WORKDIR /src
# COPY schema /schema

FROM golang:1.17 as backend
COPY go.mod go.sum /src/
WORKDIR /src
RUN go mod download
COPY ./cmd ./pkg /src/
RUN go build ./...

FROM scratch
COPY --from=backend /src/backend /backend
ENTRYPOINT ["/backend"]