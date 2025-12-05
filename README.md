# Calculator API

Веб-сервис на Go для арифметических вычислений через REST API.

### Требования

- Git
- Go 1.21+

## Быстрый запуск

### 1. Установка
```bash
git clone https://github.com/vnchk1/CalculatorAPI.git <your-directory> 
cd <your-directory> 
go mod download
```

### 2. Копирование (либо создание собственной) конфигурации
```bash
cp .env.example .env
```

### 3. Запуск
```bash
go run cmd/api/main.go
```

API будет доступен на `http://localhost:8080` (или на другом порту, указанном в конфигурации)

## API Endpoints

### POST `/sum`
Складывает числа и сохраняет результат по токену.

**Запрос:**
```json
{
  "numbers": [1, 2, 3],
  "token": "user-token-123"
}
```

**Ответ:**
```json
{
  "sum": 6
}
```

### POST `/multiply`
Умножает числа.

**Запрос:**
```json
{
  "numbers": [2, 3, 4],
  "token": "user-token-456"
}
```

**Ответ:**
```json
{
  "multiply": 24
}
```

### Ошибки (400):
```json
{"error": "Invalid request body"}
```
```json
{"error": "Empty request body"}
```

## Документация

Swagger UI: `http://localhost:8080/swagger/index.html`  (или на другом порту, указанном в конфигурации)

## Тестирование
```bash
go test ./...
```