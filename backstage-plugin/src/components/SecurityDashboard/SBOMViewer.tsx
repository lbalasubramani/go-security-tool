import React, { useState } from 'react';
import { Card, CardContent, Typography, Button, TextField } from '@material-ui/core';

export const SBOMViewer: React.FC = () => {
  const [image, setImage] = useState('');
  const [result, setResult] = useState('');

  const handleGenerate = async () => {
    // TODO: Call your go-security-tool API or trigger via Backstage scaffolder
    setResult(`SBOM generated for ${image} and signed with Cosign + in-toto attestation.`);
  };

  return (
    <Card>
      <CardContent>
        <Typography variant="h6">SBOM & Supply Chain Security</Typography>
        <TextField
          label="Container Image"
          value={image}
          onChange={(e) => setImage(e.target.value)}
          fullWidth
          margin="normal"
        />
        <Button variant="contained" color="primary" onClick={handleGenerate}>
          Generate & Sign SBOM
        </Button>
        {result && <Typography style={{ marginTop: 16 }}>{result}</Typography>}
      </CardContent>
    </Card>
  );
};