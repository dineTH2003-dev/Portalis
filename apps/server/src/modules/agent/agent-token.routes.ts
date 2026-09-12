import { Hono } from "hono";
import * as tokenService from "./agent-token.service";

export const agentTokenRoutes = new Hono();

// GET /v1/agent-tokens - List tokens
agentTokenRoutes.get("/", async (c) => {
  // TODO(contributor): [Issue #6] Extract userId from authContext
  const tokens = await tokenService.listAgentTokens("user-stub-id");
  return c.json({ tokens });
});

// POST /v1/agent-tokens - Create new token
agentTokenRoutes.post("/", async (c) => {
  const body = await c.req.json().catch(() => ({}));
  const name = body.name || "default-token";
  // TODO(contributor): [Issue #6] Extract userId from authContext
  const token = await tokenService.createAgentToken("user-stub-id", name);
  return c.json(token, 201);
});

// DELETE /v1/agent-tokens/:id - Revoke token
agentTokenRoutes.delete("/:id", async (c) => {
  const id = c.req.param("id");
  // TODO(contributor): [Issue #6] Extract userId from authContext
  await tokenService.revokeAgentToken("user-stub-id", id);
  return c.json({ success: true });
});
