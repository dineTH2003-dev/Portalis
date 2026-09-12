import React from "react";

/**
 * ActiveTunnelsPage Component
 * TODO(contributor): [Issue #26]
 *   1. Fetch user's active tunnels from GET /v1/tunnel/sessions
 *   2. Implement 5-second setInterval auto-polling
 *   3. Render table of active tunnels (subdomain link, localPort, heartbeat badge)
 *   4. Implement "Disconnect" action triggering session revocation
 *   5. Display empty state with CLI instructions when 0 tunnels exist
 */
export const ActiveTunnelsPage: React.FC = () => {
  return (
    <div style={{ maxWidth: 860, margin: "40px auto", padding: "0 20px" }}>
      <h2>Active Tunnels</h2>
      <p style={{ color: "#64748b" }}>Live tunnels connected from developer workstations.</p>

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
        <code>[Issue #26] Real-time tunnel session table and disconnect controls to be built here</code>
      </div>
    </div>
  );
};
