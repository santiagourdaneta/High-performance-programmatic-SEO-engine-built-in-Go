# Etapa 1: Compilación (Se ejecuta en los runners de GitHub)
FROM golang:1.22-alpine AS builder

WORKDIR /app

# Descarga eficiente de dependencias
COPY go.mod ./
# COPY go.sum ./ # Descomentar si usas dependencias externas
RUN go mod download

COPY . .

# Compilación estática optimizada eliminando símbolos de depuración (-s -w)
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /motor-seo main.go

# Etapa 2: Imagen final minimalista de producción
FROM scratch

# Copiar certificados SSL si tu app consume APIs externas HTTPS
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Copiar el binario puro compilado en la etapa anterior
COPY --from=builder /motor-seo /motor-seo

EXPOSE 8080

ENTRYPOINT ["/motor-seo"]