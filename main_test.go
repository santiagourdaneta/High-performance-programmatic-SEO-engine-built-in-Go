package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestContactoHandler(t *testing.T) {
	tests := []struct {
		name           string
		urlPath        string
		expectedTitle  string
		expectedStatus int
	}{
		{
			name:           "Slug valido - Optimizacion Web",
			urlPath:        "/contacto/optimizacion-web",
			expectedTitle:  "Optimizacion Web | Contacto Directo",
			expectedStatus: http.StatusOK,
		},
		{
			name:           "Slug valido - Desarrollo Software",
			urlPath:        "/contacto/desarrollo-software",
			expectedTitle:  "Desarrollo Software | Contacto Directo",
			expectedStatus: http.StatusOK,
		},
	}

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		slug := strings.TrimPrefix(r.URL.Path, "/contacto/")
		titulo := capitalizarTexto(slug) // Usamos el helper corregido

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`
			<!DOCTYPE html>
			<html>
			<head>
				<title>` + titulo + ` | Contacto Directo ` + MiNumero + `</title>
				<meta name="description" content="Contacta con expertos en ` + titulo + `.">
				<script type="application/ld+json">{"telephone": "` + MiNumero + `"}</script>
			</head>
			<body>
				<a href="tel:` + MiNumero + `">` + MiNumero + `</a>
			</body>
			</html>
		`))
	})

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest(http.MethodGet, tt.urlPath, nil)
			if err != nil {
				t.Fatalf("No se pudo crear la peticion: %v", err)
			}

			rr := httptest.NewRecorder()
			handler.ServeHTTP(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("Se esperaba status %d, se obtuvo %d", tt.expectedStatus, rr.Code)
			}

			body := rr.Body.String()

			if !strings.Contains(body, tt.expectedTitle) {
				t.Errorf("El cuerpo no contiene el titulo esperado: %s", tt.expectedTitle)
			}
		})
	}
}

func TestSitemapHandler(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "/sitemap.xml", nil)
	if err != nil {
		t.Fatalf("No se pudo crear la peticion del sitemap: %v", err)
	}

	rr := httptest.NewRecorder()

	sitemapHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/xml")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`<?xml version="1.0" encoding="UTF-8"?><urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">`))
		for _, kw := range palabrasClave {
			_, _ = w.Write([]byte(`<url><loc>` + MiSitioWeb + `/contacto/` + kw + `</loc></url>`))
		}
		_, _ = w.Write([]byte(`</urlset>`))
	})

	sitemapHandler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Se esperaba status 200, se obtuvo %d", rr.Code)
	}

	body := rr.Body.String()
	for _, kw := range palabrasClave {
		expectedURL := MiSitioWeb + "/contacto/" + kw
		if !strings.Contains(body, expectedURL) {
			t.Errorf("El sitemap no incluye la URL esperada: %s", expectedURL)
		}
	}
}
