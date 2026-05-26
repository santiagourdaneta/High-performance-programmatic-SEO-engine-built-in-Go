package main

import (
	"fmt"
	"net/http"
	"strings"
)

// Datos que queremos que se propaguen por internet
const (
	MiNumero   = "+51933391624"
	MiNombre   = "Servicios Técnicos Automatizados"
	MiSitioWeb = "https://santiagourdaneta.github.io"
)

// Lista de nichos o palabras clave para generar páginas masivas
var palabrasClave = []string{
	"desarrollo-software",
	"arquitectura-sistemas",
	"optimizacion-web",
	"consultoria-tecnica",
	"soporte-remoto",
}

func main() {
	// 1. Manejador para las páginas dinámicas
	http.HandleFunc("/contacto/", func(w http.ResponseWriter, r *http.Request) {
		// Extraer la palabra clave de la URL (ej: /contacto/optimizacion-web)
		slug := strings.TrimPrefix(r.URL.Path, "/contacto/")
		
		// Formatear el título para que se vea limpio
		titulo := strings.Title(strings.ReplaceAll(slug, "-", " "))

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		
		// Renderizado directo en HTML (Sin Node, sin dependencias pesadas)
		fmt.Fprintf(w, `<!DOCTYPE html>
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

	// 2. El mapa del sitio (Sitemap XML) que le dice a los robots qué páginas leer
	http.HandleFunc("/sitemap.xml", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/xml")
		fmt.Fprintf(w, `<?xml version="1.0" encoding="UTF-8"?>
<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">`)
		
		for _, kw := range palabrasClave {
			fmt.Fprintf(w, `
   <url>
      <loc>%s/contacto/%s</loc>
      <changefreq>daily</changefreq>
      <priority>0.8</priority>
   </url>`, MiSitioWeb, kw)
		}
		fmt.Fprintf(w, `
</urlset>`)
	})

	fmt.Println("🚀 Motor de publicación orgánica corriendo en http://localhost:8080")
	fmt.Println("📍 Revisa tu mapa de indexación en: http://localhost:8080/sitemap.xml")
	_ = http.ListenAndServe(":8080", nil)
}