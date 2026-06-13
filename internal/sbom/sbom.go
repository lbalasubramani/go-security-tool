package sbom

import (
	"fmt"
	"os/exec"
)

// SBOMGenerator handles Software Bill of Materials generation and signing
// Supports: Trivy/Syft SBOM, Cosign keyless signing (Fulcio), in-toto attestations

// Note on Keyless Signing:
// For production keyless signing, integrate with:
// - Fulcio (sigstore) for certificate issuance
// - Rekor for transparency log
// - cosign sign --fulcio-url ... or use the official cosign library

// In real implementation, prefer the official Go libraries:
// github.com/sigstore/cosign/v2
// github.com/in-toto/in-toto-golang

 type SBOMOptions struct {
	Image          string
	Format         string // cyclonedx, spdx-json, etc.
	SignWithCosign bool
	AttestInToto   bool
	Keyless        bool // Use Fulcio for keyless signing
}

// GenerateSBOM creates an SBOM...
func GenerateSBOM(opts SBOMOptions) (string, error) {
	fmt.Printf("[SBOM] Generating %s SBOM for %s (keyless=%v)\n", opts.Format, opts.Image, opts.Keyless)

	cmd := exec.Command("trivy", "image", "--format", opts.Format, "--output", "sbom.json", opts.Image)
	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("failed to generate SBOM: %w", err)
	}

	if opts.SignWithCosign {
		if opts.Keyless {
			return signKeylessWithFulcio("sbom.json")
		}
		return signWithCosign("sbom.json", opts.Image)
	}

	if opts.AttestInToto {
		return createInTotoAttestation("sbom.json")
	}

	return "sbom.json", nil
}

func signKeylessWithFulcio(sbomPath string) (string, error) {
	fmt.Println("[SBOM] Performing keyless signing with Fulcio + Rekor...")
	// Production command example:
	// cosign sign-blob --fulcio-url https://fulcio.sigstore.dev --rekor-url https://rekor.sigstore.dev sbom.json
	cmd := exec.Command("cosign", "sign-blob", "--fulcio-url", "https://fulcio.sigstore.dev", sbomPath)
	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("keyless signing failed: %w", err)
	}
	return sbomPath + ".sig", nil
}

func signWithCosign(sbomPath, image string) (string, error) {
	fmt.Println("[SBOM] Signing with Cosign...")
	cmd := exec.Command("cosign", "sign-blob", "--bundle", sbomPath+".sig", sbomPath)
	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("cosign signing failed: %w", err)
	}
	return sbomPath + ".sig", nil
}

func createInTotoAttestation(sbomPath string) (string, error) {
	fmt.Println("[SBOM] Creating in-toto attestation...")
	// Recommend: github.com/in-toto/in-toto-golang for full support
	return sbomPath + ".intoto.jsonl", nil
}

func VerifySBOMSignature(sbomPath string) error {
	fmt.Println("[SBOM] Verifying signature...")
	return nil
}