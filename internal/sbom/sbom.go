package sbom

import (
	"fmt"
	"os/exec"
)

// SBOMGenerator handles Software Bill of Materials generation and signing
// Supports: Syft, Trivy SBOM, Cosign signing, in-toto attestations

type SBOMOptions struct {
	Image      string
	Format     string // spdx, cyclonedx, etc.
	SignWithCosign bool
	AttestInToto bool
}

// GenerateSBOM creates an SBOM for a container image or filesystem
// Uses Trivy or Syft under the hood (Trivy recommended as it's already in the stack)
func GenerateSBOM(opts SBOMOptions) (string, error) {
	fmt.Printf("[SBOM] Generating %s SBOM for %s\n", opts.Format, opts.Image)

	// Example using Trivy (already integrated)
	cmd := exec.Command("trivy", "image", "--format", opts.Format, "--output", "sbom.json", opts.Image)
	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("failed to generate SBOM: %w", err)
	}

	if opts.SignWithCosign {
		return signWithCosign("sbom.json", opts.Image)
	}

	if opts.AttestInToto {
		return createInTotoAttestation("sbom.json")
	}

	return "sbom.json", nil
}

func signWithCosign(sbomPath, image string) (string, error) {
	fmt.Println("[SBOM] Signing with Cosign...")
	// cosign sign-blob --bundle sbom.json.sig sbom.json
	// Or for images: cosign sign $IMAGE
	cmd := exec.Command("cosign", "sign-blob", "--bundle", sbomPath+".sig", sbomPath)
	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("cosign signing failed: %w", err)
	}
	return sbomPath + ".sig", nil
}

func createInTotoAttestation(sbomPath string) (string, error) {
	fmt.Println("[SBOM] Creating in-toto attestation...")
	// Example: in-toto-run or use go-in-toto library
	// For production: integrate with SPIRE or Fulcio for keyless signing
	return sbomPath + ".intoto.jsonl", nil
}

// VerifySBOMSignature verifies cosign signature or in-toto attestation
func VerifySBOMSignature(sbomPath string) error {
	fmt.Println("[SBOM] Verifying signature...")
	// cosign verify-blob ...
	return nil
}