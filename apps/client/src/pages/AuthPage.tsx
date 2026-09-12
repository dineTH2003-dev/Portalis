import React from "react";

/**
 * AuthPage Component
 * TODO(contributor): [Issue #25]
 *   1. Build email input form with Turnstile bot protection
 *   2. Build 6-digit OTP verification code inputs
 *   3. Connect to api.requestOTP() and api.verifyOTP()
 *   4. Store JWT in authContext upon successful verification
 *   5. Handle error messages and resend countdown timer
 */
export const AuthPage: React.FC = () => {
  return (
    <div style={{ maxWidth: 440, margin: "80px auto", padding: 24, textAlign: "center" }}>
      <h2>Sign In to Portalis</h2>
      <p style={{ color: "#64748b", marginBottom: 24 }}>Passwordless email authentication</p>

      <div
        style={{
          padding: 40,
          border: "2px dashed #cbd5e1",
          borderRadius: 8,
          background: "#fff",
          color: "#94a3b8",
        }}
      >
        <code>[Issue #25] OTP email input form and verification inputs to be implemented here</code>
      </div>
    </div>
  );
};
