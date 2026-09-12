/**
 * Agent Token Service
 * Manages persistent CLI authentication tokens hashed via bcrypt.
 */

export interface AgentToken {
  id: string;
  name: string;
  tokenPrefix: string;
  lastUsedAt?: string;
  createdAt: string;
}

export interface CreatedTokenResponse extends AgentToken {
  plaintextToken: string; // Shown only once upon creation
}

/**
 * Generate a new named agent token.
 * TODO(contributor): [Issue #6]
 *   1. Generate 32-character cryptographically secure hex string.
 *   2. Form token: `portalis_${randomHex}`.
 *   3. Extract prefix: token.slice(0, 16).
 *   4. Hash token using bcrypt.hash(token, 10).
 *   5. Insert into agent_tokens table.
 *   6. Return CreatedTokenResponse containing plaintextToken.
 */
export async function createAgentToken(userId: string, name: string): Promise<CreatedTokenResponse> {
  console.log(`[AGENT STUB] Creating token ${name} for user ${userId}...`);
  return {
    id: "token-stub-uuid",
    name,
    tokenPrefix: "portalis_stub1234",
    plaintextToken: "portalis_stub1234567890abcdef12345678",
    createdAt: new Date().toISOString(),
  };
}

/**
 * List active agent tokens for user.
 * TODO(contributor): [Issue #6] Query agent_tokens WHERE user_id = ? AND revoked_at IS NULL.
 */
export async function listAgentTokens(userId: string): Promise<AgentToken[]> {
  return [
    {
      id: "token-stub-uuid",
      name: "laptop-cli",
      tokenPrefix: "portalis_stub1234",
      createdAt: new Date().toISOString(),
    },
  ];
}

/**
 * Revoke an agent token.
 * TODO(contributor): [Issue #6] UPDATE agent_tokens SET revoked_at = CURRENT_TIMESTAMP WHERE id = ? AND user_id = ?.
 */
export async function revokeAgentToken(userId: string, tokenId: string): Promise<boolean> {
  console.log(`[AGENT STUB] Revoking token ${tokenId} for user ${userId}...`);
  return true;
}
