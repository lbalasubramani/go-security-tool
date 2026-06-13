# ADR-0001: Build Internal Security Tooling Platform as Go + Backstage

## Status
Accepted

## Context
We need a unified security platform for Kubernetes that combines scanning, runtime detection, supply chain security, and developer experience. Inspired by how Spotify and Netflix built internal tooling.

## Decision
Build **go-security-tool** in Go with:
- Core security engines (Trivy, Falco, ZAP)
- SBOM + Cosign/Fulcio + in-toto
- Backstage plugin for the developer portal
- Grafana for observability

## Consequences
**Positive:**
- Strong interview demonstration of full security lifecycle
- Platform engineering mindset
- Extensible and GitOps-friendly

**Negative / Risks:**
- Requires maintaining multiple integrations
- Backstage plugin needs frontend expertise

## Alternatives Considered
- Pure Python security CLI
- Only Grafana dashboards (no unified tool)
- Commercial tools only (Aqua, Prisma, etc.)

## Links
- Related: Threat model, DFDs, Backstage plugin skeleton