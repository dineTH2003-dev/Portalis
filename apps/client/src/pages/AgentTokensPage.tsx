import React from "react";

/**
 * AgentTokensPage Component
 * TODO(contributor): [Issue #27]
 *   1. Fetch user's active tokens from GET /v1/agent-tokens
 *   2. Build "Generate Token" modal form with label input
 *   3. Display newly created plaintext token once with click-to-copy
 *   4. Render table of tokens with masked prefixes and revocation button
 */
export const AgentTokensPage: React.FC = () => {
  return (
    <div style={{ maxWidth: 860, margin: "40px auto", padding: "0 20px" }}>
      <h2>Agent Tokens</h2>
      <p style={{ color: "#64748b" }}>Persistent authentication tokens for the Portalis CLI.</p>

      <div
        style={{
          padding: 50,
          textAlign: "center",
          border: "2px dashed #cbd5e1",
          borderRadius: 8,
          background: "#fff",
          color: "#94a3b8",
        }}
      >
        <code>[Issue #27] Token generation modal, copy button, and active token list to be built here</code>
      </div>
    </div>
  );
};
