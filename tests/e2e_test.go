//go:build integration

package tests

import (
	"io"
	"net/http"
	"strings"
	"testing"
)

func TestGooglebotScrapingE2E(t *testing.T) {
	serverURL := "http://localhost:8080/contacto/desarrollo-software"

	req, _ := http.NewRequest(http.MethodGet, serverURL, nil)

	// Clonamos exactamente la identidad del indexador automático de Google
	req.Header.Set("User-Agent", "Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)")
	req.Header.Set("Accept", "text/html,application/xhtml+xml")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("Error E2E: El bot no pudo acceder al servidor: %v", err)
	}
	defer resp.Body.Close()

	bodyBytes, _ := io.ReadAll(resp.Body)
	htmlContent := string(bodyBytes)

	// Validaciones críticas del crawler de Google:

	// 1. Verificar inyección JSON-LD estructural
	if !strings.Contains(htmlContent, `type="application/ld+json"`) {
		t.Error("E2E FAILED: Falta la etiqueta script de datos estructurados para el robot.")
	}

	// 2. Verificar marcado telefónico Schema.org
	if !strings.Contains(htmlContent, `"telephone": "+51935137875"`) {
		t.Error("E2E FAILED: El bot de Google no pudo encontrar el nodo 'telephone' formateado en el JSON de la organización.")
	}

	// 3. Verificar Meta Descripción para las SERPs (Resultados de búsqueda)
	if !strings.Contains(htmlContent, `name="description"`) {
		t.Error("E2E FAILED: Falta la meta descripción. Google penalizará el snippet del resultado.")
	}
}
