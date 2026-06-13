package dashboard

import (
	"fmt"
	"net/http"
	"time"
)

// GrafanaClient handles integration with Grafana for security metrics
// Supports PromQL queries for Trivy Operator, Falco (via Loki), Kube-bench

type GrafanaClient struct {
	BaseURL string
	APIKey  string
}

func NewGrafanaClient(baseURL, apiKey string) *GrafanaClient {
	return &GrafanaClient{BaseURL: baseURL, APIKey: apiKey}
}

// QueryTrivyMetrics fetches vulnerability counts from Trivy Operator via Prometheus/Grafana
func (c *GrafanaClient) QueryTrivyMetrics() (string, error) {
	// Example PromQL: sum(trivy_image_vulnerabilities) by (severity)
	query := `sum(trivy_image_vulnerabilities{severity="CRITICAL"})`
	// In production: use Grafana API or Prometheus client
	fmt.Printf("[Grafana] Querying Trivy: %s\n", query)
	return "Critical vulnerabilities: 47 (last 24h)", nil
}

// QueryFalcoEvents pulls runtime security events (Falco → Loki → Grafana)
func (c *GrafanaClient) QueryFalcoEvents() (string, error) {
	// Loki query example via Grafana Explore
	query := `{job="falco"} |= "alert"`
	fmt.Printf("[Grafana] Querying Falco via Loki: %s\n", query)
	return "Falco alerts last hour: 12 (3 high severity)", nil
}

// GetDashboardURL returns embeddable Grafana dashboard URL
func (c *GrafanaClient) GetDashboardURL(dashboardUID string) string {
	return fmt.Sprintf("%s/d/%s?orgId=1&kiosk", c.BaseURL, dashboardUID)
}

// StartMetricsServer exposes /metrics endpoint for the web dashboard
func StartMetricsServer(client *GrafanaClient) {
	http.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) {
		trivy, _ := client.QueryTrivyMetrics()
		falco, _ := client.QueryFalcoEvents()
		fmt.Fprintf(w, "<div class='metrics'>\n<h3>Security Overview</h3>\n<p>%s</p>\n<p>%s</p>\n</div>", trivy, falco)
	})
}