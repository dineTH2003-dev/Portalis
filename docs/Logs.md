# Logging Guidelines & Standards

This document establishes logging conventions, log levels, structured formats, and security redaction policies across **Portalis**.

---

## 1. Standard Log Levels

Every log statement must specify an appropriate log level:

| Level | Usage | Production Behavior |
| :--- | :--- | :--- |
| **`DEBUG`** | Verbose diagnostic traces (e.g., frame serialization bytes, socket transitions). | Disabled by default. Enabled via `LOG_LEVEL=debug`. |
| **`INFO`** | Normal operational milestones (e.g., server startup, tunnel established, agent connected). | Enabled by default. Emitted as structured JSON. |
| **`WARN`** | Unexpected but recoverable anomalies (e.g., slow heartbeat response, rate-limit reached). | Always enabled. Should be monitored. |
| **`ERROR`** | Operations that failed to complete (e.g., database connection loss, invalid JWT signature). | Always enabled. Triggers alerting if monitored. |

---

## 2. Structured JSON Log Format

All server and gateway components emit structured JSON to `stdout`:

```json
{
  "timestamp": "2026-09-12T08:00:00.000Z",
  "level": "INFO",
  "component": "gateway.router",
  "message": "Tunnel connected successfully",
  "subdomain": "myapp",
  "tunnelId": "tun_abc123",
  "remoteAddr": "192.168.1.50:52341"
}
```

### Common Fields:
- `timestamp`: ISO 8601 UTC timestamp.
- `level`: `DEBUG`, `INFO`, `WARN`, `ERROR`.
- `component`: Dot-delimited subsystem name (e.g., `server.auth`, `gateway.proxy`).
- `message`: Human-readable summary of the event.
- `requestId`: Correlation ID when logging an in-flight HTTP request.

---

## 3. Mandatory Log Redaction Policy

> [!CAUTION]
> **Zero Credential Leakage**: Never log credentials or sensitive data under any circumstances.

### Fields That Must Always Be Masked or Omitted:
- Passwords or raw email OTP codes.
- `Authorization` header values (e.g., `Bearer <token>` must be masked to `Bearer ***`).
- Agent tokens (`portalis_***` or permanent tokens).
- Database connection credentials or encryption keys.

---

## 4. How to Inspect Logs Locally

### Control Plane (`apps/server`)
```bash
cd apps/server
bun run dev
```

### Go Ingress Gateway (`gateway`)
```bash
cd gateway
go run ./cmd/gateway
```

### Go Agent CLI (`agent`)
```bash
# Run agent with debug logging enabled
./bin/portalis http 3000 --subdomain myapp --debug
```
