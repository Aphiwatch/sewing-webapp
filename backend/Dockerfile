# Dockerfile สำหรับ Backend
FROM golang:1.20 AS builder
WORKDIR /app

# ติดตั้ง dependencies และสร้างแอป
COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .
RUN go build -o main ./cmd/main.go

FROM gcr.io/distroless/base-debian11
COPY --from=builder /app/main /app/main
CMD ["/app/main"]
