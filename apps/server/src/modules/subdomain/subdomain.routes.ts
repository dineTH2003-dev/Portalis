import { Hono } from "hono";
import * as subdomainService from "./subdomain.service";

export const subdomainRoutes = new Hono();

// GET /v1/subdomains - List reserved subdomains
subdomainRoutes.get("/", async (c) => {
  // TODO(contributor): [Issue #7] Extract userId from authContext
  const subdomains = await subdomainService.listUserSubdomains("user-stub-id");
  return c.json({ subdomains });
});

// POST /v1/subdomains - Reserve a subdomain
subdomainRoutes.post("/", async (c) => {
  const body = await c.req.json().catch(() => ({}));
  const subdomain = body.subdomain;

  if (!subdomain) {
    return c.json({ error: "subdomain is required" }, 400);
  }

  // TODO(contributor): [Issue #7] Extract userId from authContext
  const reserved = await subdomainService.reserveSubdomain("user-stub-id", subdomain);
  return c.json(reserved, 201);
});
