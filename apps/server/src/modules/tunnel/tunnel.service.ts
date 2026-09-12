/**
 * Tunnel Session Service
 * Manages active developer tunnel sessions and Tunnel Grant JWT generation.
 */

export interface TunnelSession {
  id: string;
  subdomain: string;
  localPort: number;
  publicUrl: string;
  status: "pending" | "active" | "disconnected" | "revoked";
  connectedAt?: string;
  lastHeartbeatAt?: string;
}

export interface TunnelGrant {
  sessionId: string;
  grantToken: string; // Cryptographic JWT presented to Gateway WebSocket
  publicUrl: string;
}

/**
 * Issue a new tunnel grant for an authenticated agent.
 * TODO(contributor): [Issue #8]
 *   1. Validate agent token and verify user active status.
 *   2. Check user's max_tunnels quota.
 *   3. Check subdomain collision:
 *      - Ensure not in blocked_subdomains.
 *      - Ensure not reserved by another user.
 *   4. Create tunnel_sessions table record with status 'pending'.
 *   5. Sign HS256 Tunnel Grant JWT (30m TTL) with payload:
 *      { jti, sid, uid, sdn: subdomain, prt: localPort }.
 */
export async function createTunnelGrant(
  userId: string,
  subdomain: string,
  localPort: number
): Promise<TunnelGrant> {
  console.log(`[TUNNEL STUB] Creating grant for ${subdomain} -> localhost:${localPort}...`);
  return {
    sessionId: "session-stub-uuid",
    grantToken: "stub.jwt.grant_token_for_gateway",
    publicUrl: `http://${subdomain}.localhost:8080`,
  };
}

/**
 * List active tunnels for user.
 * TODO(contributor): [Issue #8] Query tunnel_sessions WHERE user_id = ? AND status = 'active'.
 */
export async function listUserTunnels(userId: string): Promise<TunnelSession[]> {
  return [
    {
      id: "session-stub-uuid",
      subdomain: "demo",
      localPort: 3000,
      publicUrl: "http://demo.localhost:8080",
      status: "active",
      connectedAt: new Date().toISOString(),
      lastHeartbeatAt: new Date().toISOString(),
    },
  ];
}
