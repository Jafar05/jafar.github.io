FROM golang:1.23-slim

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o /app/main .

CMD ["go", "run", "main.go"]