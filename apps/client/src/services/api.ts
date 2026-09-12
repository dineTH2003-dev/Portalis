/**
 * API Service Client
 * Typed API helper functions communicating with the Portalis Control Plane.
 */

const API_BASE = import.meta.env.VITE_API_URL || "http://localhost:4310";

function getAuthHeader(): Record<string, string> {
  const token = localStorage.getItem("portalis_token");
  return token ? { Authorization: `Bearer ${token}` } : {};
}

// TODO(contributor): [Issue #25] Request email OTP
export async function requestOTP(email: string): Promise<{ challengeId: string; expiresInSeconds: number }> {
  const res = await fetch(`${API_BASE}/v1/auth/request-otp`, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ email }),
  });
  if (!res.ok) throw new Error("Failed to request OTP challenge");
  return res.json();
}

// TODO(contributor): [Issue #25] Verify OTP and return tokens
export async function verifyOTP(challengeId: string, otp: string): Promise<{ accessToken: string; refreshToken: string; user: any }> {
  const res = await fetch(`${API_BASE}/v1/auth/verify-otp`, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ challengeId, otp }),
  });
  if (!res.ok) throw new Error("Verification failed: invalid or expired code");
  return res.json();
}

// TODO(contributor): [Issue #26] Fetch user active tunnels
export async function fetchActiveTunnels(): Promise<{ tunnels: any[] }> {
  const res = await fetch(`${API_BASE}/v1/tunnel/sessions`, {
    headers: getAuthHeader(),
  });
  if (!res.ok) throw new Error("Failed to fetch active tunnels");
  return res.json();
}

// TODO(contributor): [Issue #27] Create a new agent token
export async function createAgentToken(name: string): Promise<{ plaintextToken: string }> {
  const res = await fetch(`${API_BASE}/v1/agent-tokens`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
      ...getAuthHeader(),
    },
    body: JSON.stringify({ name }),
  });
  if (!res.ok) throw new Error("Failed to create agent token");
  return res.json();
}

// TODO(contributor): [Issue #28] Reserve a vanity subdomain
export async function reserveSubdomain(subdomain: string): Promise<any> {
  const res = await fetch(`${API_BASE}/v1/subdomains`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
      ...getAuthHeader(),
    },
    body: JSON.stringify({ subdomain }),
  });
  if (!res.ok) throw new Error("Failed to reserve subdomain");
  return res.json();
}
