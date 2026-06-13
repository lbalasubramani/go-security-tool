import React from 'react';

interface GrafanaEmbedProps {
  dashboardUrl: string;
  title?: string;
}

export const GrafanaEmbed: React.FC<GrafanaEmbedProps> = ({ dashboardUrl, title = 'Security Metrics' }) => {
  return (
    <div style={{ width: '100%', height: '600px', border: '1px solid #ddd' }}>
      <iframe
        src={dashboardUrl}
        title={title}
        width="100%"
        height="100%"
        frameBorder="0"
      />
    </div>
  );
};