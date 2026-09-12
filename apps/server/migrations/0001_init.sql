-- 0001_init.sql: Core Identity & Authentication Schema
-- TODO(contributor): [Issue #4] Execute via migration runner

CREATE TABLE IF NOT EXISTS users (
    id TEXT PRIMARY KEY,
    email TEXT NOT NULL UNIQUE,
    role TEXT NOT NULL DEFAULT 'user', -- 'user' | 'admin'
    status TEXT NOT NULL DEFAULT 'active', -- 'pending' | 'active' | 'suspended'
    name TEXT NOT NULL DEFAULT '',
    max_tunnels INTEGER NOT NULL DEFAULT 3,
    max_subdomains INTEGER NOT NULL DEFAULT 3,
    created_at TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS otp_challenges (
    id TEXT PRIMARY KEY,
    email TEXT NOT NULL,
    otp_hash TEXT NOT NULL,
    expires_at TEXT NOT NULL,
    consumed_at TEXT,
    failed_attempts INTEGER NOT NULL DEFAULT 0,
    created_at TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_otp_email ON otp_challenges(email);
CREATE INDEX IF NOT EXISTS idx_otp_expires ON otp_challenges(expires_at);

CREATE TABLE IF NOT EXISTS sessions (
    id TEXT PRIMARY KEY,
    user_id TEXT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    expires_at TEXT NOT NULL,
    revoked_at TEXT,
    created_at TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_sessions_user ON sessions(user_id);
