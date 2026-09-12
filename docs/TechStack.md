# Technology Stack

This document details the core technologies, runtimes, and libraries powering **Portalis**.

---

## 1. Stack Overview

| Subsystem | Technology | Runtime / Version | Key Libraries | Purpose |
| :--- | :--- | :--- | :--- | :--- |
| **Data Plane (Gateway)** | Go | Go 1.23+ | `gorilla/websocket`, `golang-jwt/jwt/v5` | Ingress HTTP routing, WebSocket termination, frame multiplexing |
| **Developer CLI (Agent)** | Go | Go 1.23+ | `gorilla/websocket`, `net/http` | Cross-platform binary, local port forwarding, outbound tunnel connection |
| **Shared Protocol** | Go | Go 1.23+ | Standard Library `encoding/json` | Canonical frame models and payload validation |
| **Control Plane** | TypeScript | Bun 1.1+ | `hono`, `jose`, `bcryptjs`, `bun:sqlite` | Identity, passwordless OTP auth, token hashing, grant issuance, quotas |
| **Dashboard** | React 18 / TypeScript | Vite 6+ | `react-router-dom`, `lucide-react` | Developer self-service dashboard & admin governance console |
| **Persistence** | SQLite 3 | WAL Mode | Native `bun:sqlite` driver | Zero-setup embedded relational database with high-concurrency WAL |

---

## 2. Why These Technologies?

### Go 1.23 (Gateway & Agent CLI)
- **High Concurrency**: Goroutines consume ~2KB initial stack memory, enabling thousands of concurrent WebSocket tunnels with minimal RAM usage.
- **Cross-Platform Compilation**: Compiles into static, zero-dependency binaries for Linux, macOS (Apple Silicon / Intel), and Windows with a single `go build` command.
- **Battle-Tested Networking**: Standard library `net/http` provides robust, high-performance HTTP multiplexing.

### Bun & Hono (Control Plane)
- **High-Performance TypeScript**: Native TypeScript execution without transpile or build steps, sub-millisecond startup times.
- **In-Process SQLite (`bun:sqlite`)**: Direct C-level SQLite bindings that execute queries significantly faster than traditional Node drivers.
- **Hono Framework**: Ultra-lightweight, edge-ready router with end-to-end type safety and middleware support.

### React 18 & Vite (Dashboard)
- **Instant HMR**: Sub-second hot-module replacement during local development.
- **Optimized Bundling**: Fast production rollup builds producing lightweight asset bundles.

### SQLite in WAL Mode (Persistence)
- **Zero Ops Overhead**: No separate database cluster (PostgreSQL/MySQL) to provision or maintain for self-hosting.
- **Write-Ahead Logging (WAL)**: Readers never block writers, and writers never block readers.
