import { Hono } from "hono";
import { authRoutes } from "./modules/auth/auth.routes";
import { agentTokenRoutes } from "./modules/agent/agent-token.routes";
import { tunnelRoutes } from "./modules/tunnel/tunnel.routes";
import { subdomainRoutes } from "./modules/subdomain/subdomain.routes";
import { adminRoutes } from "./modules/admin/admin.routes";

export const app = new Hono();

// Health Check Endpoint
app.get("/health", (c) => c.json({ status: "healthy", timestamp: new Date().toISOString() }));

// Mount Versioned API Route Groups
// TODO(contributor): [Issue #3] Add global error handling and request logging middleware
app.route("/v1/auth", authRoutes);
app.route("/v1/agent-tokens", agentTokenRoutes);
app.route("/v1/tunnel", tunnelRoutes);
app.route("/v1/subdomains", subdomainRoutes);
app.route("/v1/admin", adminRoutes);
