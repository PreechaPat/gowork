FROM golang:1.25-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o gowork .

FROM debian:trixie-slim
COPY --from=builder /app/gowork /gowork
COPY --from=builder /app/dist /dist
EXPOSE 8080
ENTRYPOINT ["/gowork"]
