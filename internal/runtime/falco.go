package runtime

import (
	"fmt"
	"time"
)

// FalcoClient provides integration with Falco for runtime threat detection
// Supports: Falco sidekick, gRPC output, or direct event ingestion

type FalcoEvent struct {
	Time      time.Time
	Priority  string
	Rule      string
	Output    string
	Source    string // k8s_audit, syscall, etc.
}

// ProcessFalcoEvents ingests and processes runtime security events
// Production: Connect to Falco sidekick HTTP/gRPC or read from Loki
func ProcessFalcoEvents(events []FalcoEvent) {
	for _, e := range events {
		if e.Priority == "CRITICAL" || e.Priority == "ERROR" {
			fmt.Printf("[Falco ALERT] %s: %s (Source: %s)\n", e.Rule, e.Output, e.Source)
			// TODO: Forward to PagerDuty, Slack, or Security Hub
			// TODO: Trigger automated response (e.g., isolate pod via Kubernetes API)
		}
	}
}

// Example: Simulate receiving events from Falco sidekick webhook
func StartFalcoWebhookListener() {
	// In real impl: http server listening on /falco-webhook
	fmt.Println("[Falco] Starting webhook listener for sidekick events...")
}

// GetFalcoRules returns recommended CNCF Falco rules for K8s + containers
func GetFalcoRules() []string {
	return []string{
		"Terminal shell in container",
		"Write to /etc/hosts",
		"Unexpected outbound connection",
		"K8s service account token mounted",
		"Privileged container launched",
	}
}