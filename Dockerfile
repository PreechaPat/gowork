FROM golang:1.25-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o gowork .

FROM debian:trixie-slim
ARG SLIM=false
COPY --from=builder /app/gowork /gowork
COPY --from=builder /app/dist /dist

RUN if [ "$SLIM" = "false" ]; then \
    apt-get update && \
    apt-get install -y --no-install-recommends curl procps && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*; \
    fi

EXPOSE 8080
ENTRYPOINT ["/gowork"]
