# go-security-tool

**Enterprise Security Tooling Platform** — built like Spotify & Netflix internal tools.

A comprehensive Go-based security platform for Kubernetes environments, featuring:

- Vulnerability scanning (Trivy)
- Runtime threat detection (Falco)
- Dynamic scanning (OWASP ZAP)
- **SBOM generation + Cosign + in-toto signing**
- Grafana observability integration
- **Backstage plugin** for Spotify-style developer portal experience
- Full threat modeling (STRIDE + DFD)

## Quick Start

```bash
go run cmd/security-tool/main.go
```

## Key Modules

- `internal/dashboard/grafana.go` — PromQL queries for Trivy & Falco
- `internal/runtime/falco.go` — Runtime security event processing
- `internal/sbom/sbom.go` — SBOM + Cosign/in-toto signing
- `internal/scanner/zap.go` — DAST integration
- `backstage-plugin/` — Ready-to-use Backstage frontend plugin

## Backstage Integration (Recommended)

Copy `backstage-plugin/` into your Backstage monorepo and register it. This brings security visibility directly into the developer portal — the exact pattern used by mature platform teams.

## Interview-Ready

This project demonstrates:
- End-to-end threat modeling
- Shift-left + runtime security
- Supply chain security (SBOM + signing)
- Platform engineering mindset (Backstage + golden paths)

Perfect for Senior Application Security, Product Security, and Security Software Engineer roles.