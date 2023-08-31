FROM golang:1.19 as builder
LABEL authors="evgeniy.ershov"

# Установка рабочей директории внутри контейнера
WORKDIR /go/src/my_grades

# Копирование go mod
COPY go.mod .
COPY go.sum .

RUN go mod download

# Копирование файлов проекта внутрь контейнера
COPY . .

RUN go build -o /bin/app ./cmd/main.go

FROM alpine:3.11.5
WORKDIR /app

COPY --from=builder /bin/app ./

ENTRYPOINT ["./app"]