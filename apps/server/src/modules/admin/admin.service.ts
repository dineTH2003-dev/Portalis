/**
 * Admin Governance Service
 * Handles platform KPIs, user management, and audit log inspection.
 */

export interface PlatformStats {
  totalUsers: number;
  activeTunnels: number;
  activeTokens: number;
}

/**
 * Retrieve high-level platform KPI statistics.
 * TODO(contributor): [Issue #10] Aggregate counts from users, tunnel_sessions, and agent_tokens tables.
 */
export async function getPlatformStats(): Promise<PlatformStats> {
  return {
    totalUsers: 1,
    activeTunnels: 0,
    activeTokens: 1,
  };
}
