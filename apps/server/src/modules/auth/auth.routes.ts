import { Hono } from "hono";
import * as authService from "./auth.service";

export const authRoutes = new Hono();

/**
 * POST /v1/auth/request-otp
 * Request an email verification code.
 */
authRoutes.post("/request-otp", async (c) => {
  const body = await c.req.json().catch(() => ({}));
  const email = body.email;

  if (!email || typeof email !== "string") {
    return c.json({ error: "Email is required" }, 400);
  }

  const result = await authService.requestOTP(email);
  return c.json(result, 201);
});

/**
 * POST /v1/auth/verify-otp
 * Verify email OTP and receive JWT access/refresh tokens.
 */
authRoutes.post("/verify-otp", async (c) => {
  const body = await c.req.json().catch(() => ({}));
  const { challengeId, otp } = body;

  if (!challengeId || !otp) {
    return c.json({ error: "challengeId and otp are required" }, 400);
  }

  const tokens = await authService.verifyOTP(challengeId, otp);
  return c.json(tokens, 200);
});
