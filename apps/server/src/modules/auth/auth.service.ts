/**
 * Authentication Service
 * Handles passwordless OTP challenges, verification, and JWT session issuance.
 */

export interface OTPRequestResult {
  challengeId: string;
  expiresInSeconds: number;
}

export interface AuthTokens {
  accessToken: string;
  refreshToken: string;
  user: {
    id: string;
    email: string;
    role: string;
  };
}

/**
 * Request an email OTP challenge.
 * TODO(contributor): [Issue #5]
 *   1. Validate email syntax.
 *   2. Generate random 6-digit code: Math.floor(100000 + Math.random() * 900000).
 *   3. Hash code using SHA-256.
 *   4. Insert challenge into otp_challenges table with 10-minute expiry.
 *   5. Send OTP via email provider (or log to console in dev mode).
 */
export async function requestOTP(email: string): Promise<OTPRequestResult> {
  console.log(`[AUTH STUB] Generating OTP for ${email}...`);
  return {
    challengeId: "stub-challenge-uuid",
    expiresInSeconds: 600,
  };
}

/**
 * Verify OTP code and issue JWT session tokens.
 * TODO(contributor): [Issue #5]
 *   1. Lookup challenge in otp_challenges table.
 *   2. Verify not expired and not already consumed.
 *   3. Increment failed_attempts on mismatch; lockout after 5 failures.
 *   4. Upsert user record (first user registered gets 'admin' role).
 *   5. Sign Access JWT (15m TTL) and Refresh JWT (30d TTL) using 'jose'.
 *   6. Create session record in database.
 */
export async function verifyOTP(challengeId: string, otp: string): Promise<AuthTokens> {
  console.log(`[AUTH STUB] Verifying challenge ${challengeId} with code ${otp}...`);
  return {
    accessToken: "stub_access_token_jwt",
    refreshToken: "stub_refresh_token_jwt",
    user: {
      id: "user-uuid-stub",
      email: "developer@portalis.dev",
      role: "admin",
    },
  };
}
