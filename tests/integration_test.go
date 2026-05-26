//go:build integration

package tests

import (
	"io"
	"net/http"
	"strings"
	"testing"
	"time"
)

func TestServerIntegration(t *testing.T) {
	serverURL := "http://localhost:8080/contacto/optimizacion-web"

	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	resp, err := client.Get(serverURL)
	if err != nil {
		t.Fatalf("Falló la conexión de integración con el servidor: %v. ¿Te aseguraste de ejecutar 'make dev' en otra terminal?", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Código de estado inesperado: obtenido %d, se esperaba 200", resp.StatusCode)
	}

	contentType := resp.Header.Get("Content-Type")
	if !strings.Contains(contentType, "text/html") {
		t.Errorf("Header Content-Type incorrecto: obtenido %s", contentType)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("No se pudo leer el cuerpo de la respuesta: %v", err)
	}

	// Sincronización del número de teléfono en el layout real
	if !strings.Contains(string(body), "+51935137875") {
		t.Error("La integración falló: El número de teléfono no se encuentra en el HTML final renderizado por la red.")
	}
}
