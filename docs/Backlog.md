# Portalis Engineering Backlog: From-Scratch Development Roadmap

Welcome to the official **Portalis** engineering backlog. This document lays out the step-by-step master plan to build Portalis from scratch as an open-source platform. 

Every item below is an atomic, self-contained task designed to be created as a **GitHub Issue** that university batchmates and contributors can assign themselves, build, test, and submit as a Pull Request.

---

## Roadmap Overview

```
Phase 1: Protocol Foundation ──► Phase 2: Control Plane ──► Phase 3: Go Gateway
                                                                    │
Phase 5: React Dashboard  ◄─── Phase 6: E2E Testing   ◄─── Phase 4: Go Agent CLI
```

---

## Phase 1: Shared Protocol Foundation

### Issue #1: `[Protocol] Define canonical Go frame protocol structures`
- **Area**: `area:protocol` | **Type**: `type:feature` | **Complexity**: `XS` | **Label**: `good first issue`
- **Scope**: `frames/protocol.go`
- **Description**: Create the canonical wire format structs for communication between the Gateway and Agent.
- **Tasks**:
  - Define `FrameType` constants (`request`, `response`, `ping`, `pong`, `error`).
  - Define `HeaderMap map[string][]string`.
  - Define `RequestFrame` (`Type`, `RequestID`, `Method`, `Path`, `Headers`, `Body []byte`).
  - Define `ResponseFrame` (`Type`, `RequestID`, `StatusCode`, `Headers`, `Body []byte`, `Error string`).
  - Define `ControlFrame` (`Type`, `Error string`).
- **Acceptance Criteria**: Frames compile cleanly with `go 1.23`.

---

### Issue #2: `[Protocol] Implement frame validation and serialization unit tests`
- **Area**: `area:protocol` | **Type**: `type:test` | **Complexity**: `S` | **Label**: `good first issue`
- **Scope**: `frames/validation.go`, `frames/protocol_test.go`
- **Description**: Add validation helper methods (`Validate()`) to verify non-empty `RequestID`, valid HTTP methods, and valid status codes (100-599). Add JSON serialization tests.
- **Acceptance Criteria**: `go test -v ./...` in `frames/` passes with 100% code coverage.

---

## Phase 2: Control Plane (Bun + Hono + SQLite)

### Issue #3: `[Server] Initialize Bun + Hono project and SQLite database wrapper`
- **Area**: `area:control-plane` | **Type**: `type:feature` | **Complexity**: `S`
- **Scope**: `apps/server/src/core/db/`
- **Description**: Setup `package.json` for Bun, configure Hono, and implement SQLite initialization using `bun:sqlite` with mandatory PRAGMAs:
  - `PRAGMA journal_mode = WAL;`
  - `PRAGMA foreign_keys = ON;`
  - `PRAGMA busy_timeout = 5000;`
  - `PRAGMA synchronous = NORMAL;`
- **Acceptance Criteria**: Database connection opens safely with WAL mode enabled.

---

### Issue #4: `[Server] Create database migration runner and initial SQL schema`
- **Area**: `area:control-plane` | **Type**: `type:feature` | **Complexity**: `M`
- **Scope**: `apps/server/migrations/`, `apps/server/scripts/migrate.ts`
- **Description**: Implement a forward-only migration runner that executes `.sql` files in sequence and tracks applied migrations in a `schema_migrations` table.
- **Schema to create**:
  - `users` (id, email, role, status, max_tunnels, max_subdomains)
  - `otp_challenges` (id, email, otp_hash, expires_at, failed_attempts)
  - `sessions` (id, user_id, expires_at, revoked_at)
  - `agent_tokens` (id, user_id, name, token_hash, token_prefix, last_used_at, revoked_at)
  - `tunnel_sessions` (id, user_id, subdomain, local_port, status, grant_jti, connected_at, last_heartbeat_at)
  - `reserved_subdomains` (id, subdomain, user_id, status)
  - `blocked_subdomains` (id, subdomain, reason)
  - `audit_logs` (id, actor_user_id, action, ip_address, metadata_json, created_at)
- **Acceptance Criteria**: `bun run migrate` applies all migrations cleanly.

---

### Issue #5: `[Server] Implement 2-step passwordless OTP authentication`
- **Area**: `area:control-plane` | **Type**: `type:feature` | **Complexity**: `M`
- **Scope**: `apps/server/src/modules/auth/`
- **Description**:
  - `POST /v1/auth/request-otp`: Accepts `{ email }`. Generates a random 6-digit OTP, saves SHA-256 hash in `otp_challenges` (10-min TTL), and logs or emails the code.
  - `POST /v1/auth/verify-otp`: Accepts `{ challengeId, otp }`. Verifies code, enforces max 5 failed attempts lockout, provisions user on first login, and returns Access & Refresh JWTs.
- **Acceptance Criteria**: Integration tests verify complete OTP issue and verification cycle.

---

### Issue #6: `[Server] Implement Agent Token management API`
- **Area**: `area:control-plane` | **Type**: `type:feature` | **Complexity**: `S`
- **Scope**: `apps/server/src/modules/agent/`
- **Description**:
  - `POST /v1/agent-tokens`: Generates `portalis_<32-hex-chars>`, hashes via bcrypt (10 rounds), returns plaintext token once.
  - `GET /v1/agent-tokens`: Lists active agent tokens for the authenticated user (with masked tokens).
  - `DELETE /v1/agent-tokens/:id`: Revokes an agent token.
- **Acceptance Criteria**: Token generation, bcrypt verification, and revocation pass tests.

---

### Issue #7: `[Server] Implement Subdomain reservation and collision validation`
- **Area**: `area:control-plane` | **Type**: `type:feature` | **Complexity**: `S`
- **Scope**: `apps/server/src/modules/subdomain/`
- **Description**: Allow developers to reserve custom vanity subdomains. Validate format (3-63 lowercase alphanumeric with hyphens), verify it is not blocked, and enforce per-user `max_subdomains` quota.
- **Acceptance Criteria**: Collision detection rejects duplicates and returns 409/422.

---

### Issue #8: `[Server] Implement Tunnel Session creation & short-lived Grant JWT issuance`
- **Area**: `area:control-plane` | **Type**: `type:feature` | **Complexity**: `M`
- **Scope**: `apps/server/src/modules/tunnel/`
- **Description**: When agent calls `POST /v1/tunnel/sessions` with an `agent_token` and desired `localPort` and `subdomain`:
  - Validate agent token and check user `max_tunnels` quota.
  - Create `tunnel_sessions` record (status: `pending`).
  - Sign short-lived **Tunnel Grant JWT** (30m TTL) containing `jti`, `sid`, `uid`, `sdn`, `prt`.
- **Acceptance Criteria**: CLI receives signed grant token to present to Gateway.

---

### Issue #9: `[Server] Implement internal Gateway introspection & session callbacks`
- **Area**: `area:control-plane` | **Type**: `type:feature` | **Complexity**: `S`
- **Scope**: `apps/server/src/modules/tunnel/tunnel-internal.routes.ts`
- **Description**: Endpoints protected by `x-gateway-secret`:
  - `GET /v1/internal/tunnel/grants/:jti/introspect`: Gateway verifies grant validity.
  - `POST /v1/internal/tunnel/sessions/:id/connected`: Transitions session to `active`.
  - `POST /v1/internal/tunnel/sessions/:id/disconnected`: Transitions session to `disconnected`.
  - `POST /v1/internal/tunnel/sessions/:id/heartbeat`: Updates `last_heartbeat_at`.
- **Acceptance Criteria**: Gateway can introspect grants and report state updates.

---

### Issue #10: `[Server] Implement Operator Admin APIs and Audit Trail`
- **Area**: `area:control-plane` | **Type**: `type:feature` | **Complexity**: `M`
- **Scope**: `apps/server/src/modules/admin/`, `apps/server/src/modules/audit/`
- **Description**: Admin-only endpoints (`role === 'admin'`) for listing users, approving accounts, toggling quotas, blocking subdomains, force-disconnecting tunnels, and inspecting audit logs.
- **Acceptance Criteria**: Non-admin users are rejected with 403 Forbidden.

---

## Phase 3: Go Ingress Gateway (Data Plane)

### Issue #11: `[Gateway] Scaffold Go Ingress Gateway & configuration loader`
- **Area**: `area:gateway` | **Type**: `type:feature` | **Complexity**: `S` | **Label**: `good first issue`
- **Scope**: `gateway/cmd/gateway/main.go`, `gateway/internal/config/`
- **Description**: Initialize Go module `gateway`, implement environment variable configuration loader (`HTTP_PORT=8080`, `WS_PORT=9000`, `SERVER_API_URL`, `INTERNAL_GATEWAY_SECRET`, `TUNNEL_GRANT_SECRET`).
- **Acceptance Criteria**: Gateway loads configuration cleanly and logs startup banner.

---

### Issue #12: `[Gateway] Implement Agent WebSocket listener & Grant verification`
- **Area**: `area:gateway` | **Type**: `type:feature` | **Complexity**: `M`
- **Scope**: `gateway/internal/auth/agent_server.go`
- **Description**: Listen on `:9000` at `/v1/tunnel/ws` for agent connections using `gorilla/websocket`. Extract `Bearer <jwt>` grant token, verify signature, and call Control Plane `/introspect`.
- **Acceptance Criteria**: Invalid tokens are rejected with 401/403; valid tokens establish WebSocket.

---

### Issue #13: `[Gateway] Implement thread-safe in-memory Subdomain Router`
- **Area**: `area:gateway` | **Type**: `type:feature` | **Complexity**: `M`
- **Scope**: `gateway/internal/router/`
- **Description**: Implement a thread-safe routing table using `sync.RWMutex` mapping subdomain strings to active `Tunnel` instances. Handle tunnel registration, lookup, eviction, and concurrent reads.
- **Acceptance Criteria**: Unit test with `-race` flag verifies zero race conditions under concurrent access.

---

### Issue #14: `[Gateway] Implement public wildcard HTTP ingress listener`
- **Area**: `area:gateway` | **Type**: `type:feature` | **Complexity**: `M`
- **Scope**: `gateway/internal/proxy/http_server.go`
- **Description**: HTTP server listening on `:8080`. Extracts subdomain from `Host` header (e.g. `myapp.localhost:8080` -> `myapp`), queries Router. Returns styled 404 HTML if no active tunnel found.
- **Acceptance Criteria**: Public HTTP requests correctly identify target tunnel.

---

### Issue #15: `[Gateway] Implement bidirectional request/response frame multiplexing`
- **Area**: `area:gateway` | **Type**: `type:feature` | **Complexity**: `L`
- **Scope**: `gateway/internal/proxy/http_server.go`
- **Description**: Convert incoming HTTP requests to `RequestFrame`, register a response channel keyed by `RequestID`, transmit frame over the agent's WebSocket. Listen for `ResponseFrame`, write headers and body back to public client. Handle 30s timeout with 504 Gateway Timeout.
- **Acceptance Criteria**: Concurrently proxies multiple HTTP requests across a single WebSocket connection.

---

### Issue #16: `[Gateway] Implement 30s Heartbeat Monitor & stale session reaping`
- **Area**: `area:gateway` | **Type**: `type:feature` | **Complexity**: `M`
- **Scope**: `gateway/internal/heartbeat/monitor.go`
- **Description**: Dedicated goroutine running every 30 seconds. Sends `ControlFrame{Type: "ping"}` to all active tunnels. Updates `LastHeartbeat` on `pong`. If no pong for > 120s, evicts tunnel and calls `/disconnected`.
- **Acceptance Criteria**: Stale/dead tunnels are automatically purged from memory.

---

### Issue #17: `[Gateway] Implement graceful shutdown to drain in-flight requests`
- **Area**: `area:gateway` | **Type**: `type:feature` | **Complexity**: `M`
- **Scope**: `gateway/cmd/gateway/main.go`
- **Description**: Listen for `SIGINT` and `SIGTERM`. Stop accepting new HTTP/WS connections, allow up to 15 seconds for pending request channels to complete, notify Control Plane, and exit cleanly.
- **Acceptance Criteria**: No abrupt socket termination on server restart.

---

## Phase 4: Go Agent CLI

### Issue #18: `[Agent] Scaffold Go Agent CLI and local configuration store`
- **Area**: `area:agent` | **Type**: `type:feature` | **Complexity**: `S` | **Label**: `good first issue`
- **Scope**: `agent/cmd/agent/main.go`, `agent/internal/config/`
- **Description**: Create the CLI tool structure. Store user credentials locally in `~/.portalis/config.json` (`serverUrl`, `userEmail`, `agentToken`). Implement `portalis config get` and `portalis version`.
- **Acceptance Criteria**: CLI creates and reads `~/.portalis/config.json` safely with restricted file permissions (`0600`).

---

### Issue #19: `[Agent] Implement portalis login command`
- **Area**: `area:agent` | **Type**: `type:feature` | **Complexity**: `S`
- **Scope**: `agent/cmd/agent/main.go`, `agent/internal/api/`
- **Description**: `portalis login <agent-token> [--server <url>]`. Validates token format (`portalis_...`), contacts Control Plane (`POST /v1/auth/agent-login`), and persists configuration.
- **Acceptance Criteria**: Successful login displays welcome message with user email.

---

### Issue #20: `[Agent] Implement portalis http <port> tunnel command`
- **Area**: `area:agent` | **Type**: `type:feature` | **Complexity**: `M`
- **Scope**: `agent/cmd/agent/main.go`, `agent/internal/tunnel/client.go`
- **Description**: `portalis http <port> [--subdomain <name>]`.
  - Validate port (1-65535).
  - Request Tunnel Grant from Control Plane (`POST /v1/tunnel/sessions`).
  - Establish persistent WebSocket connection to Gateway (`/v1/tunnel/ws`) with Bearer token.
  - Display active tunnel banner with public URL and local target.
- **Acceptance Criteria**: CLI establishes and displays active tunnel connection.

---

### Issue #21: `[Agent] Implement local HTTP request forwarder & response streaming`
- **Area**: `area:agent` | **Type**: `type:feature` | **Complexity**: `L`
- **Scope**: `agent/internal/tunnel/client.go`
- **Description**: In the WebSocket read loop:
  - On `RequestFrame`: Dispatch goroutine to forward HTTP request to `http://127.0.0.1:<localPort><path>`.
  - Copy request headers and body.
  - If local port is down, respond with 502 Bad Gateway.
  - On response, package status, headers, and body into `ResponseFrame` and transmit over WebSocket.
  - On `ping`: Immediately answer `pong`.
- **Acceptance Criteria**: Public requests to the tunnel successfully reach local server and return responses.

---

### Issue #22: `[Agent] Implement automatic reconnection with exponential backoff`
- **Area**: `area:agent` | **Type**: `type:feature` | **Complexity**: `M`
- **Scope**: `agent/internal/tunnel/reconnect.go`
- **Description**: If the WebSocket connection drops unexpectedly (network blip, gateway restart), do not crash. Automatically retry connecting with exponential backoff (1s, 2s, 4s, 8s, up to 30s) while keeping the session alive.
- **Acceptance Criteria**: Reconnecting network automatically recovers tunnel without user restart.

---

### Issue #23: `[Agent] Add verbose debug logging & latency timing metrics`
- **Area**: `area:agent` | **Type**: `type:feature` | **Complexity**: `S` | **Label**: `good first issue`
- **Scope**: `agent/internal/tunnel/client.go`
- **Description**: Add `--verbose` / `-v` flag to output real-time request logs: `[200 OK] GET /api/users - 24ms`.
- **Acceptance Criteria**: Verbose mode logs formatted HTTP status codes, method, path, and duration.

---

## Phase 5: React Dashboard SPA

### Issue #24: `[Dashboard] Scaffold React 18 + Vite + TypeScript application`
- **Area**: `area:dashboard` | **Type**: `type:feature` | **Complexity**: `S` | **Label**: `good first issue`
- **Scope**: `apps/client/`
- **Description**: Initialize React 18 with Vite, TypeScript, React Router DOM v6, Lucide Icons, and global layout with responsive navbar and sidebar.
- **Acceptance Criteria**: `bun run dev` renders responsive application shell.

---

### Issue #25: `[Dashboard] Implement Passwordless OTP Authentication Flow`
- **Area**: `area:dashboard` | **Type**: `type:feature` | **Complexity**: `M`
- **Scope**: `apps/client/src/pages/AuthPage.tsx`, `apps/client/src/store/authContext.tsx`
- **Description**: Email input form, 6-digit OTP verification input, JWT token storage in memory/localStorage, protected route wrapper, and logout button.
- **Acceptance Criteria**: User can complete full login flow and access protected routes.

---

### Issue #26: `[Dashboard] Implement Live Active Tunnels Dashboard`
- **Area**: `area:dashboard` | **Type**: `type:feature` | **Complexity**: `M`
- **Scope**: `apps/client/src/pages/ActiveTunnelsPage.tsx`
- **Description**: Live table of user's active tunnels with 5-second auto-polling. Display subdomain, public URL link, target local port, heartbeat indicator, and "Disconnect" button.
- **Acceptance Criteria**: Newly created tunnels appear in real time; disconnect terminates session.

---

### Issue #27: `[Dashboard] Implement Agent Token Management UI`
- **Area**: `area:dashboard` | **Type**: `type:feature` | **Complexity**: `S`
- **Scope**: `apps/client/src/pages/AgentTokensPage.tsx`
- **Description**: List user's active tokens. "Generate Token" modal displaying plaintext token once with click-to-copy. Revoke token action with confirmation dialog.
- **Acceptance Criteria**: Token generation and one-click revocation function smoothly.

---

### Issue #28: `[Dashboard] Implement Subdomain Reservation UI`
- **Area**: `area:dashboard` | **Type**: `type:feature` | **Complexity**: `S`
- **Scope**: `apps/client/src/pages/SubdomainsPage.tsx`
- **Description**: View list of reserved vanity subdomains. Reservation input with live client-side validation (3-63 lowercase alphanumeric). Delete/release reservation action.
- **Acceptance Criteria**: Valid reservations are created and displayed; collision errors are handled gracefully.

---

### Issue #29: `[Dashboard] Implement Operator Admin Console`
- **Area**: `area:dashboard` | **Type**: `type:feature` | **Complexity**: `L`
- **Scope**: `apps/client/src/pages/admin/`
- **Description**: Admin view (`role === 'admin'`) with:
  - KPI overview metrics (total users, active tunnels, sessions).
  - User governance (search, status approve/suspend, quota overrides).
  - Cluster tunnel oversight (view all tunnels, force disconnect).
  - Blocked subdomains registry.
  - Paginated audit log explorer.
- **Acceptance Criteria**: Admin console provides full control over users and active tunnels.

---

## Phase 6: Testing & Quality Assurance

### Issue #30: `[Testing] Add comprehensive unit tests for Gateway Router and Proxy`
- **Area**: `area:gateway` | **Type**: `type:test` | **Complexity**: `M`
- **Scope**: `gateway/internal/router/router_test.go`, `gateway/internal/proxy/http_server_test.go`
- **Description**: Write Go unit tests verifying subdomain extraction, thread-safe concurrent routing, frame creation, and 404/504 error handling.
- **Acceptance Criteria**: `go test -race ./...` in `gateway/` achieves >80% coverage.

---

### Issue #31: `[Testing] Add comprehensive unit tests for Agent CLI HTTP forwarder`
- **Area**: `area:agent` | **Type**: `type:test` | **Complexity**: `M`
- **Scope**: `agent/internal/tunnel/client_test.go`
- **Description**: Use `httptest.Server` to mock local HTTP targets and verify request/response frame transformations, status code propagation, and 502 handling when port is down.
- **Acceptance Criteria**: `go test -v ./...` in `agent/` passes.

---

### Issue #32: `[Testing] Implement End-to-End Tunnel Integration Test`
- **Area**: `area:gateway` | **Type**: `type:test` | **Complexity**: `L`
- **Scope**: `tests/e2e_test.go`
- **Description**: Automated integration test that boots Control Plane, boots Gateway, runs Agent CLI against a dummy HTTP test server, sends public HTTP request through Gateway, and asserts response matches expected payload.
- **Acceptance Criteria**: Automated test validates end-to-end tunnel path.
