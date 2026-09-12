import { describe, it, expect } from "bun:test";
import { app } from "./app";

describe("Control Plane Starter Tests", () => {
  it("GET /health returns healthy status", async () => {
    const res = await app.request("/health");
    expect(res.status).toBe(200);
    const body = (await res.json()) as { status: string };
    expect(body.status).toBe("healthy");
  });

  // TODO(contributor): [Issue #5] Add tests for /v1/auth/request-otp and /v1/auth/verify-otp
  // TODO(contributor): [Issue #6] Add tests for /v1/agent-tokens
  // TODO(contributor): [Issue #7] Add tests for /v1/subdomains
  // TODO(contributor): [Issue #8] Add tests for /v1/tunnel/sessions
});
