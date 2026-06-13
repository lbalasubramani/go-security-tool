# Backstage Security Plugin (go-security-tool)

**Spotify-style internal security portal plugin** for Backstage.

This plugin integrates **go-security-tool** (Trivy, Falco, Grafana, SBOM, ZAP) directly into your Backstage developer portal.

## Why Backstage?
- Created by Spotify, now CNCF project
- Perfect for enterprise Internal Developer Platforms (IDPs)
- Single pane of glass for security + developer experience

## Plugin Structure
```
backstage-plugin/
├── package.json
├── src/
│   ├── plugin.ts                 # Plugin registration
│   ├── components/
│   │   ├── SecurityDashboard/
│   │       SecurityDashboard.tsx
│   │       SecurityOverviewCard.tsx
│   └── api/
│       securityApi.ts          # Calls go-security-tool API or Grafana
└── README.md
```

## Quick Start
1. Copy this folder into your Backstage app under `plugins/security-tool`
2. Run `yarn install` in the plugin
3. Register the plugin in `app-config.yaml` and `packages/app/src/App.tsx`
4. Add the `<SecurityDashboard />` card to entity pages or a dedicated Security tab

## Features
- Live security metrics from Grafana (Trivy, Falco)
- SBOM viewer + verification status
- Trigger scans from Backstage UI
- Links to detailed go-security-tool dashboard

## Integration with go-security-tool
The plugin can call:
- Your Go service REST API (Fiber/Huma)
- Grafana directly (embedded dashboards)
- Or proxy through your Platform API (PAPI)

## Next Steps
- Add real API client using `@backstage/core-plugin-api`
- Embed Grafana panels via iframe or Grafana React plugin
- Add RBAC for who can trigger scans
- Surface SBOM + signing status per service

This turns your internal security tooling into a first-class citizen in the developer portal — exactly how Spotify and other mature platforms do it.