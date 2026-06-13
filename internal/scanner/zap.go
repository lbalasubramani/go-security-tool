package scanner

import (
	"fmt"
	"os/exec"
)

// ZAPScanner provides OWASP ZAP integration for dynamic application security testing (DAST)

// RunZAPScan executes a ZAP scan against a target (web app, API, K8s service)
// Production: Use ZAP Docker image or ZAP API (zapv2 Go client recommended)
func RunZAPScan(target string, scanType string) error {
	fmt.Printf("[ZAP] Starting %s scan on %s\n", scanType, target)

	// Example using ZAP Docker (recommended for K8s/CI)
	cmd := exec.Command("docker", "run", "--rm", "-t", "owasp/zap2docker-stable",
		"zap-baseline.py", "-t", target, "-r", "zap-report.html")

	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("ZAP scan failed: %w\nOutput: %s", err, output)
	}

	fmt.Println("[ZAP] Scan completed. Report: zap-report.html")
	// TODO: Parse report, store in DB, send to Grafana / Security Dashboard
	// TODO: Integrate with Trivy results for full picture
	return nil
}

// Recommended ZAP configurations for K8s environments
func GetZAPBestPractices() []string {
	return []string{
		"Use ZAP in headless mode inside cluster",
		"Scan internal services via ClusterIP or port-forward",
		"Integrate with CI (GitHub Actions / Tekton)",
		"Combine with Trivy for image + runtime + DAST coverage",
		"Export results to SARIF for unified vulnerability management",
	}
}