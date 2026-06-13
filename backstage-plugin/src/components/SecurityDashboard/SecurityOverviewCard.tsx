import React from 'react';
import { Card, CardContent, Typography, Chip } from '@material-ui/core';

interface Props {
  title: string;
  value: string;
  status: 'success' | 'warning' | 'error';
  link?: string;
}

export const SecurityOverviewCard: React.FC<Props> = ({ title, value, status, link }) => {
  const color = status === 'success' ? 'primary' : status === 'warning' ? 'secondary' : 'error';
  return (
    <Card>
      <CardContent>
        <Typography variant="h6">{title}</Typography>
        <Typography variant="h4" color={color}>
          {value}
        </Typography>
        <Chip label={status.toUpperCase()} color={color} size="small" />
        {link && <a href={link}>View in Grafana →</a>}
      </CardContent>
    </Card>
  );
};