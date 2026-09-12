import React from "react";
import { BrowserRouter, Routes, Route, Link } from "react-router-dom";
import { AuthProvider, useAuth } from "./store/authContext";
import { LandingPage } from "./pages/LandingPage";
import { AuthPage } from "./pages/AuthPage";
import { ActiveTunnelsPage } from "./pages/ActiveTunnelsPage";
import { AgentTokensPage } from "./pages/AgentTokensPage";
import { SubdomainsPage } from "./pages/SubdomainsPage";
import { AdminDashboardPage } from "./pages/AdminDashboardPage";

const Navigation: React.FC = () => {
  const { isAuthenticated, logout } = useAuth();
  return (
    <nav style={{ display: "flex", justifyContent: "space-between", padding: "16px 24px", background: "#fff", borderBottom: "1px solid #e2e8f0" }}>
      <Link to="/" style={{ fontSize: 20, fontWeight: 700, textDecoration: "none", color: "#0f172a" }}>
        Portalis
      </Link>
      <div style={{ display: "flex", gap: 20 }}>
        {isAuthenticated ? (
          <>
            <Link to="/dashboard" style={{ textDecoration: "none", color: "#334155" }}>Tunnels</Link>
            <Link to="/tokens" style={{ textDecoration: "none", color: "#334155" }}>Tokens</Link>
            <Link to="/subdomains" style={{ textDecoration: "none", color: "#334155" }}>Subdomains</Link>
            <Link to="/admin" style={{ textDecoration: "none", color: "#334155" }}>Admin</Link>
            <button onClick={logout} style={{ border: "none", background: "none", color: "#ef4444" }}>Logout</button>
          </>
        ) : (
          <Link to="/auth" style={{ textDecoration: "none", color: "#2563eb", fontWeight: 600 }}>Sign In</Link>
        )}
      </div>
    </nav>
  );
};

export const App: React.FC = () => {
  return (
    <AuthProvider>
      <BrowserRouter>
        <Navigation />
        <Routes>
          <Route path="/" element={<LandingPage />} />
          <Route path="/auth" element={<AuthPage />} />
          <Route path="/dashboard" element={<ActiveTunnelsPage />} />
          <Route path="/tokens" element={<AgentTokensPage />} />
          <Route path="/subdomains" element={<SubdomainsPage />} />
          <Route path="/admin" element={<AdminDashboardPage />} />
        </Routes>
      </BrowserRouter>
    </AuthProvider>
  );
};

export default App;
