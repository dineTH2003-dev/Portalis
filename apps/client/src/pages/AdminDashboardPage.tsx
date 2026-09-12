import React from "react";

/**
 * AdminDashboardPage Component
 * TODO(contributor): [Issue #29]
 *   1. Protect route verifying user.role === 'admin'
 *   2. Fetch KPI metrics from GET /v1/admin/stats
 *   3. Build user review queue table (approve pending, suspend, edit quotas)
 *   4. Render cluster-wide active tunnels with force-disconnect action
 *   5. Build paginated audit log explorer
 */
export const AdminDashboardPage: React.FC = () => {
  return (
    <div style={{ maxWidth: 960, margin: "40px auto", padding: "0 20px" }}>
      <h2>Operator Admin Console</h2>
      <p style={{ color: "#64748b" }}>Platform metrics, user governance, and audit logging.</p>

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
        <code>[Issue #29] Admin KPI cards, user review table, and audit log explorer to be built here</code>
      </div>
    </div>
  );
};
