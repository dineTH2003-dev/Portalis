import React from "react";

/**
 * SubdomainsPage Component
 * TODO(contributor): [Issue #28]
 *   1. Fetch user's reserved subdomains from GET /v1/subdomains
 *   2. Build reservation input form with client-side validation (3-63 lowercase alphanumeric)
 *   3. Handle error states: 409 Conflict (taken) or 422 Unprocessable (blocked)
 *   4. Render table of reserved subdomains with release/delete action
 */
export const SubdomainsPage: React.FC = () => {
  return (
    <div style={{ maxWidth: 860, margin: "40px auto", padding: "0 20px" }}>
      <h2>Custom Subdomains</h2>
      <p style={{ color: "#64748b" }}>Reserve vanity subdomains for your local developer services.</p>

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
        <code>[Issue #28] Subdomain reservation input and reservations list to be built here</code>
      </div>
    </div>
  );
};
