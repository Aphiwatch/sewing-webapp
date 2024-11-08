# Dockerfile สำหรับ Backend
FROM golang:1.22.5
WORKDIR /app

# ติดตั้ง dependencies และสร้างแอป
COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main .

EXPOSE 8080

CMD ["./main"]