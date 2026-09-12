import React from "react";
import { Link } from "react-router-dom";

/**
 * LandingPage Component
 * TODO(contributor): [Issue #24]
 *   - Design responsive hero banner with Portalis branding
 *   - Add feature highlight cards (Multiplexing, Security, Subdomains, CLI)
 *   - Add quickstart terminal snippet
 */
export const LandingPage: React.FC = () => {
  return (
    <div style={{ maxWidth: 800, margin: "60px auto", padding: "0 20px", textAlign: "center" }}>
      <h1>Portalis 🚀</h1>
      <p style={{ color: "#64748b", fontSize: "1.2rem" }}>
        Self-hosted developer connectivity and tunneling platform.
      </p>

      <div style={{ margin: "30px 0" }}>
        <Link
          to="/auth"
          style={{
            padding: "10px 24px",
            background: "#2563eb",
            color: "#fff",
            borderRadius: 6,
            textDecoration: "none",
            fontWeight: 600,
          }}
        >
          Sign In
        </Link>
      </div>

      <div
        style={{
          padding: 40,
          border: "2px dashed #cbd5e1",
          borderRadius: 8,
          background: "#fff",
          color: "#94a3b8",
        }}
      >
        <code>[Issue #24] Feature cards and hero introduction to be built here</code>
      </div>
    </div>
  );
};
