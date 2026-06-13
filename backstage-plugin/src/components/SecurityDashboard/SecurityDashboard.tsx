import React from 'react';
import { Typography, Grid, Card, CardContent } from '@material-ui/core';
import { SecurityOverviewCard } from './SecurityOverviewCard';

export const SecurityDashboard = () => {
  return (
    <>
      <Typography variant="h4">Security Overview</Typography>
      <Grid container spacing={3}>
        <Grid item xs={12} md={6}>
          <SecurityOverviewCard
            title="Vulnerabilities (Trivy)"
            value="47 Critical"
            status="warning"
            link="/grafana/trivy"
          />
        </Grid>
        <Grid item xs={12} md={6}>
          <SecurityOverviewCard
            title="Runtime Alerts (Falco)"
            value="12 High Severity"
            status="error"
            link="/grafana/falco"
          />
        </Grid>
        <Grid item xs={12}>
          <Card>
            <CardContent>
              <Typography variant="h6">SBOM & Supply Chain</Typography>
              <Typography>Latest SBOM signed with Cosign + in-toto attestation verified.</Typography>
              {/* TODO: Call go-security-tool API for real SBOM status */}
            </CardContent>
          </Card>
        </Grid>
      </Grid>
    </>
  );
};