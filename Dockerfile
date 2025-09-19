
FROM golang:1.24.5-alpine AS builder

# Install git for Goose
RUN apk add --no-cache git

WORKDIR /app


COPY go.mod go.sum ./


RUN go mod download

# Install Goose
RUN go install github.com/pressly/goose/v3/cmd/goose@latest

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o trinity-palette .


FROM alpine:latest


RUN apk --no-cache add ca-certificates


RUN addgroup -g 1001 -S appgroup && \
    adduser -u 1001 -S appuser -G appgroup


WORKDIR /app

COPY --from=builder /app/trinity-palette .


COPY --chown=appuser:appgroup templates/ ./templates/
COPY --chown=appuser:appgroup static/ ./static/


RUN chown -R appuser:appgroup /app


USER appuser


EXPOSE 8080

CMD ["./trinity-palette"]