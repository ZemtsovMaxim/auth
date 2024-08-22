FROM golang:latest

# Установка wait-for-it для ожидания поднятия сервисов
ADD https://raw.githubusercontent.com/vishnubob/wait-for-it/master/wait-for-it.sh /usr/local/bin/wait-for-it
RUN chmod +x /usr/local/bin/wait-for-it

# Установка make
RUN apt-get update && apt-get install -y make

# Включаем модульную систему Go
RUN go env -w GO111MODULE=on

# Устанавливаем рабочую директорию и копируем зависимости
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

# Копируем весь проект
COPY . .

# # Устанавливаем рабочую директорию на auth
# WORKDIR /app

# Компилируем и запускаем приложение после применения миграций
CMD /usr/local/bin/wait-for-it db:5432 --timeout=30 -- \
    sh -c "make migrate && go run cmd/auth/main.go --config=configs/local.yaml & \
    /usr/local/bin/wait-for-it localhost:8080 --timeout=40 -- make test & \
    wait"