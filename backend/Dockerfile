# Build stage
FROM golang:1.22-alpine AS build

RUN apk add --no-cache gcc musl-dev

WORKDIR /app

# Copy dependency files first for better caching
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build binary with CGO enabled (SQLite support)
RUN CGO_ENABLED=1 go build -o skillpath ./cmd/server/

# Runtime stage
FROM alpine:3.19

RUN apk add --no-cache ca-certificates tzdata

# Create non-root user
RUN addgroup -S app && adduser -S app -G app

WORKDIR /app

COPY --from=build /app/skillpath .

USER app

EXPOSE 3005

HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
  CMD wget -qO- http://localhost:3005/health || exit 1

CMD ["./skillpath"]
