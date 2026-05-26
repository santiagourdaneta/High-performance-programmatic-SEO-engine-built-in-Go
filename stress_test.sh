#!/bin/bash
echo "=== Lanzando Prueba de Carga y Estrés Masivo ==="
echo "Enviando 50,000 peticiones totales con 200 conexiones concurrentes..."

# Ejecuta el benchmark contra la ruta de SEO dinámico
hey -n 50000 -c 200 http://localhost:8080/contacto/arquitectura-sistemas

echo "=== Prueba completada ==="