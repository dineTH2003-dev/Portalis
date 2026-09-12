import { Hono } from "hono";
import * as tunnelService from "./tunnel.service";

export const tunnelRoutes = new Hono();

// GET /v1/tunnel/sessions - List active user tunnels
tunnelRoutes.get("/sessions", async (c) => {
  // TODO(contributor): [Issue #8] Extract userId from authContext
  const tunnels = await tunnelService.listUserTunnels("user-stub-id");
  return c.json({ tunnels });
});

// POST /v1/tunnel/sessions - Request tunnel grant
tunnelRoutes.post("/sessions", async (c) => {
  const body = await c.req.json().catch(() => ({}));
  const { subdomain, localPort } = body;

  if (!subdomain || !localPort) {
    return c.json({ error: "subdomain and localPort are required" }, 400);
  }

  // TODO(contributor): [Issue #8] Validate agent token from header
  const grant = await tunnelService.createTunnelGrant("user-stub-id", subdomain, Number(localPort));
  return c.json(grant, 201);
});
