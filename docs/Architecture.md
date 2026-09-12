# Platform Architecture

This document describes the architectural design, component boundaries, and security model of **Portalis**.

---

## 1. High-Level System Topology

![Portalis Architecture](images/portalis-architecture.png)

```
[Public Web Browser / Webhook Sender]
             │
             │ HTTPS (*.domain.com)
             ▼
   [Go Ingress Gateway] (:8080 / :9000)
             │
             ├── (Introspect Tunnel Grants) ──► [Bun Control Plane] (:4310)
             │                                          │
             │ Persistent WebSocket                     ▼
             │ (/v1/tunnel/ws)                   [SQLite (WAL)]
             ▼
     [Go Agent CLI]
             │
             │ Local HTTP Request
             ▼
  [Local Service (:3000 / :8080)]
```

---

## 2. Core Subsystems

### A. Go Ingress Gateway (`gateway/`) — Data Plane
- **Role**: High-concurrency network edge.
- **Responsibilities**:
  - Accept public wildcard HTTP traffic.
  - Terminate persistent WebSocket connections from Agent CLIs.
  - Multiplex inbound HTTP requests into JSON request frames.
  - Match response frames back to waiting HTTP client connections using unique `requestId`s.
  - Send keepalive ping/pong frames and disconnect dead tunnels.
- **Constraint**: Must *never* connect to the database directly or handle password hashing.

### B. Bun / Hono Control Plane (`apps/server/`) — Control Plane
- **Role**: Governance, authentication, and platform management.
- **Responsibilities**:
  - Passwordless 2-step email OTP authentication.
  - User identity, rate limiting, and tunnel session tracking.
  - Agent token generation (bcrypt hashed) and short-lived Tunnel Grant (JWT) issuance.
  - Vanity subdomain reservation, collision detection, and word blocklist enforcement.
  - REST APIs for the React Dashboard.

### C. Go Agent CLI (`agent/`) — Developer Endpoint
- **Role**: Developer machine binary.
- **Responsibilities**:
  - Authenticate against Control Plane and obtain a short-lived Tunnel Grant.
  - Connect via outbound WebSocket to the Gateway.
  - Unpack incoming `RequestFrame`s, execute local HTTP requests against `localhost:<port>`, and stream back `ResponseFrame`s.

### D. Canonical Protocol Layer (`frames/`) — Shared Contract
- **Role**: Single source of truth for transport frames.
- **Frame Types**:
  - `request`: Carries HTTP method, path, headers, and body.
  - `response`: Carries status code, response headers, and body.
  - `ping` / `pong`: Keepalive heartbeats.
  - `error`: Diagnostic error messages.

---

## 3. Security & Token Hierarchy

Portalis separates privileges across four distinct token tiers:

```
1. Admin Session Token (JWT, 15-min TTL)
   └─ Used by developers to manage account, subdomains, and tokens via Web Dashboard.

2. Agent Token (Permanent, bcrypt hashed)
   └─ Configured in developer CLI (~/.portalis/config.json) to authenticate the agent.

3. Tunnel Grant (JWT, 5-min TTL, HS256)
   └─ Ephemeral grant passed by Agent to Gateway to establish a WebSocket tunnel.

4. Correlation Request ID (UUID)
   └─ Ephemeral identifier isolating individual concurrent HTTP requests over a tunnel.
```

---

## 4. Request / Response Multiplexing Flow

1. External user requests `https://myapp.domain.com/users`.
2. **Go Gateway** matches `myapp` in its memory routing table and generates a unique `requestId`.
3. Gateway wraps the HTTP request into a `RequestFrame` and sends it over the active WebSocket to the Agent.
4. **Agent CLI** unpacks the frame, executes `GET http://localhost:3000/users`, and gathers the response.
5. Agent packages the response into a `ResponseFrame` and sends it back across the WebSocket.
6. Gateway receives the frame, matches `requestId`, and writes the HTTP status, headers, and body back to the public caller.
