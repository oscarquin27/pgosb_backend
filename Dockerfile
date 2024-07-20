FROM golang:1.22.1

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o cmd/main cmd/main.go

EXPOSE 5000

RUN cd ./cmd/
CMD ["main"]