import { Hono } from "hono";
import * as adminService from "./admin.service";

export const adminRoutes = new Hono();

// GET /v1/admin/stats - Retrieve platform KPIs
adminRoutes.get("/stats", async (c) => {
  // TODO(contributor): [Issue #10] Add adminGuard middleware verifying role === 'admin'
  const stats = await adminService.getPlatformStats();
  return c.json(stats);
});
