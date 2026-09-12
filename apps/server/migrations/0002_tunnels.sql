-- 0002_tunnels.sql: Agent Tokens, Tunnel Sessions, and Subdomains
-- TODO(contributor): [Issue #4] Execute via migration runner

CREATE TABLE IF NOT EXISTS agent_tokens (
    id TEXT PRIMARY KEY,
    user_id TEXT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    name TEXT NOT NULL,
    token_hash TEXT NOT NULL UNIQUE,
    token_prefix TEXT NOT NULL,
    last_used_at TEXT,
    revoked_at TEXT,
    created_at TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_agent_tokens_prefix ON agent_tokens(token_prefix);

CREATE TABLE IF NOT EXISTS tunnel_sessions (
    id TEXT PRIMARY KEY,
    user_id TEXT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    agent_token_id TEXT REFERENCES agent_tokens(id),
    subdomain TEXT NOT NULL,
    local_port INTEGER NOT NULL,
    public_url TEXT,
    status TEXT NOT NULL DEFAULT 'pending', -- 'pending' | 'active' | 'disconnected' | 'revoked'
    grant_jti TEXT UNIQUE,
    connected_at TEXT,
    disconnected_at TEXT,
    last_heartbeat_at TEXT,
    created_at TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_tunnel_subdomain ON tunnel_sessions(subdomain);

CREATE TABLE IF NOT EXISTS reserved_subdomains (
    id TEXT PRIMARY KEY,
    subdomain TEXT NOT NULL UNIQUE,
    user_id TEXT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    status TEXT NOT NULL DEFAULT 'active',
    created_at TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS blocked_subdomains (
    id TEXT PRIMARY KEY,
    subdomain TEXT NOT NULL UNIQUE,
    reason TEXT,
    created_at TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS audit_logs (
    id TEXT PRIMARY KEY,
    actor_user_id TEXT REFERENCES users(id),
    action TEXT NOT NULL,
    entity_type TEXT,
    entity_id TEXT,
    ip_address TEXT,
    metadata_json TEXT,
    created_at TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_audit_action ON audit_logs(action);
