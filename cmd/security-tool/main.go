package main

import (
	"flag"
	"fmt"
	"os"

	"go-security-tool/internal/dashboard"
	"go-security-tool/internal/sbom"
)

func main() {
	fmt.Println("Go Security Tool - Enterprise Security Platform")

	// Simple subcommand parsing
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "sbom":
			handleSBOM(os.Args[2:])
			return
		case "dashboard":
			dashboard.StartMetricsServer(dashboard.NewGrafanaClient("http://localhost:3000", ""))
			return
		}
	}

	// Default: show help
	fmt.Println("Usage:")
	fmt.Println("  security-tool dashboard          - Start web dashboard")
	fmt.Println("  security-tool sbom generate <image> [--sign]   - Generate & optionally sign SBOM")
}

func handleSBOM(args []string) {
	if len(args) == 0 {
		fmt.Println("sbom generate <image> [--sign]")
		return
	}

	cmd := args[0]
	switch cmd {
	case "generate":
		if len(args) < 2 {
			fmt.Println("Usage: security-tool sbom generate <image> [--sign]")
			return
		}
		image := args[1]
		sign := false
		for _, a := range args {
			if a == "--sign" {
				sign = true
			}
		}

		result, err := sbom.GenerateSBOM(sbom.SBOMOptions{
			Image:          image,
			Format:         "cyclonedx",
			SignWithCosign: sign,
			AttestInToto:   sign,
		})
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("SBOM generated: %s\n", result)
	default:
		fmt.Println("Unknown sbom command:", cmd)
	}
}