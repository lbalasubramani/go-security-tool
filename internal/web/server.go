package web

import (
	"fmt"
	"net/http"
)

// Simple HTMX-ready web server for security dashboard
func StartServer() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `<!DOCTYPE html>
<html>
<head><title>Security Dashboard</title></head>
<body>
<h1>go-security-tool Dashboard</h1>
<div hx-get="/metrics" hx-trigger="load">Loading metrics...</div>
<script src="https://unpkg.com/htmx.org@1.9.12"></script>
</body>
</html>`)
	})
	http.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "<p>Trivy: 12 vulns | Falco: 3 events | Kube-bench: Compliant</p>")
	})
	http.ListenAndServe(":8080", nil)
}