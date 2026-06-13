import {
  createPlugin,
  createRoutableExtension,
} from '@backstage/core-plugin-api';
import { rootRouteRef } from './routes';

export const securityToolPlugin = createPlugin({
  id: 'security-tool',
  routes: {
    root: rootRouteRef,
  },
});

export const SecurityDashboardPage = securityToolPlugin.provide(
  createRoutableExtension({
    name: 'SecurityDashboardPage',
    component: () =>
      import('./components/SecurityDashboard/SecurityDashboard').then(m => m.SecurityDashboard),
    mountPoint: rootRouteRef,
  }),
);