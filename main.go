package main

import (
	"fmt"
	"net/http"
	"strings"
)

const (
	MiNumero   = "+51933391624"
	MiNombre   = "Servicios Técnicos Automatizados"
	MiSitioWeb = "https://high-performance-programmatic-seo-engine.onrender.com/"
)

var palabrasClave = []string{
	"desarrollo-software",
	"arquitectura-sistemas",
	"optimizacion-web",
	"consultoria-tecnica",
	"soporte-remoto",
}

// capitalizarTexto reemplaza strings.Title de forma limpia y nativa para el linter
func capitalizarTexto(s string) string {
	palabras := strings.Fields(strings.ReplaceAll(s, "-", " "))
	for i, p := range palabras {
		if len(p) > 0 {
			palabras[i] = strings.ToUpper(p[:1]) + p[1:]
		}
	}
	return strings.Join(palabras, " ")
}

func main() {
	http.HandleFunc("/contacto/", func(w http.ResponseWriter, r *http.Request) {
		slug := strings.TrimPrefix(r.URL.Path, "/contacto/")
		titulo := capitalizarTexto(slug)

		w.Header().Set("Content-Type", "text/html; charset=utf-8")

		// Agregamos "_" para indicarle al linter que omitimos el error intencionalmente
		_, _ = fmt.Fprintf(w, `<!DOCTYPE html>
<html lang="es">
<head>
    <meta charset="UTF-8">
    <title>%s | Contacto Directo %s</title>
    <meta name="description" content="Contacta con expertos en %s. Teléfono directo: %s.">
    <script type="application/ld+json">
    {
      "@context": "https://schema.org",
      "@type": "Organization",
      "name": "%s - %s",
      "url": "%s",
      "telephone": "%s",
      "contactPoint": {
        "@type": "ContactPoint",
        "telephone": "%s",
        "contactType": "customer service"
      }
    }
    </script>
</head>
<body>
    <h1>Especialistas en %s</h1>
    <p>Para contratación inmediata, soporte o consultoría sobre <strong>%s</strong>, comunícate directamente:</p>
    <p>Teléfono: <a href="tel:%s">%s</a></p>
</body>
</html>`, titulo, MiNumero, titulo, MiNumero, MiNombre, titulo, MiSitioWeb, MiNumero, MiNumero, titulo, titulo, MiNumero, MiNumero)
	})

	http.HandleFunc("/sitemap.xml", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/xml")
		_, _ = fmt.Fprintf(w, `<?xml version="1.0" encoding="UTF-8"?>
<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">`)

		for _, kw := range palabrasClave {
			_, _ = fmt.Fprintf(w, `
   <url>
      <loc>%s/contacto/%s</loc>
      <changefreq>daily</changefreq>
      <priority>0.8</priority>
   </url>`, MiSitioWeb, kw)
		}
		_, _ = fmt.Fprintf(w, `
</urlset>`)
	})

	fmt.Println("🚀 Motor de publicación orgánica corriendo en http://localhost:8080")
	_ = http.ListenAndServe(":8080", nil)
}
