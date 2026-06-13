#!/bin/bash
set -e

echo "=== go-security-tool Demo ==="

echo "\n[1/4] Building the tool..."
make build

echo "\n[2/4] Running SBOM generation + signing..."
make sbom

echo "\n[3/4] Starting dashboard in background (for demo)..."
# In real demo you would run: make dashboard

echo "\n[4/4] Example Backstage plugin usage..."
echo "The Backstage plugin can now call the Go service via the SecurityClient API."

echo "\n=== Demo complete! ==="
echo "Next: Deploy the Go service, register the Backstage plugin, and enjoy unified security visibility."