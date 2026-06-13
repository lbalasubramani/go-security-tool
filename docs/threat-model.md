# Threat Model for go-security-tool

## 1. Assets
- Kubernetes clusters & workloads
- Container images & SBOMs
- Runtime processes (nodes, pods)
- Security scan results & dashboards
- User/service accounts (OIDC/EntraID)

## 2. Data Flow Diagram (DFD)
See `plantuml-dfd.puml` for visual.

High-level flow:
User/Dev → CLI/Web UI → Scanner (Trivy/ZAP) → Storage (DB/Metrics) → Dashboard (Grafana)
Runtime (Falco) → Alerts → Response (Automation)

## 3. STRIDE Threat Analysis

| Threat Category | Example Threat | Mitigation in go-security-tool |
|-----------------|----------------|--------------------------------|
| **Spoofing**    | Attacker impersonates user or service | OIDC/EntraID + mTLS between components; signed requests |
| **Tampering**   | Modify scan results or SBOMs | Immutable storage, signed SBOMs (cosign/in-toto), checksums |
| **Repudiation** | Deny performing a scan or action | Full audit logging (to Loki/Elastic), immutable logs, digital signatures |
| **Information Disclosure** | Leak vulnerability data or cluster topology | RBAC + least privilege, encrypted at rest/transit (Keyvault), redaction in reports |
| **Denial of Service** | Flood scanner or overwhelm dashboard | Rate limiting, circuit breakers, resource quotas, caching |
| **Elevation of Privilege** | Escalate to cluster-admin via compromised tool | Strict RBAC, just-in-time access, network policies (Cilium), pod security standards |

## 4. Recommended Controls
- Shift-left: Trivy + ZAP in CI + admission controllers
- Runtime: Falco + eBPF rules + automated response
- Observability: Grafana + Loki + Thanos for metrics/logs/traces
- Supply Chain: SBOM + signing + policy (OPA/Gatekeeper)

## 5. Next Steps for Deep Threat Modeling
- Create detailed DFD per component
- Run threat modeling workshop (e.g., using Threat Dragon or IriusRisk)
- Map to MITRE ATT&CK for Kubernetes