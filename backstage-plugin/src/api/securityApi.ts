import { createApiRef, DiscoveryApi, FetchApi } from '@backstage/core-plugin-api';

export const securityApiRef = createApiRef<SecurityApi>({
  id: 'plugin.security-tool.service',
});

export interface SecurityApi {
  getSecurityOverview(): Promise<SecurityOverview>;
  generateSBOM(image: string, sign: boolean): Promise<string>;
}

interface SecurityOverview {
  vulnerabilities: number;
  falcoAlerts: number;
  sbomStatus: string;
}

export class SecurityClient implements SecurityApi {
  private readonly discoveryApi: DiscoveryApi;
  private readonly fetchApi: FetchApi;

  constructor(options: { discoveryApi: DiscoveryApi; fetchApi: FetchApi }) {
    this.discoveryApi = options.discoveryApi;
    this.fetchApi = options.fetchApi;
  }

  private async getBaseUrl() {
    // Configure this to point to your deployed go-security-tool service
    // or use Backstage proxy for /go-security-tool
    return await this.discoveryApi.getBaseUrl('go-security-tool');
  }

  async getSecurityOverview(): Promise<SecurityOverview> {
    const baseUrl = await this.getBaseUrl();
    const response = await this.fetchApi.fetch(`${baseUrl}/metrics`);
    // In real impl: parse JSON from your Go service
    return {
      vulnerabilities: 47,
      falcoAlerts: 12,
      sbomStatus: 'Signed with Cosign + in-toto',
    };
  }

  async generateSBOM(image: string, sign: boolean): Promise<string> {
    const baseUrl = await this.getBaseUrl();
    const response = await this.fetchApi.fetch(`${baseUrl}/sbom`, {
      method: 'POST',
      body: JSON.stringify({ image, sign }),
    });
    return response.text();
  }
}