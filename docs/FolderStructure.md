# Repository Folder Structure

This document outlines the directory layout and component responsibilities across the **Portalis** repository.

---

## 1. Directory Layout

```
Portalis/
├── .github/                  # GitHub Actions CI, community automation, and issue templates
│   ├── workflows/            # Workflow definitions (ci.yml, discord-notify.yml)
│   ├── ISSUE_TEMPLATE/       # GitHub issue report forms (bug, feature, refactor, security)
│   └── PULL_REQUEST_TEMPLATE.md
│
├── apps/                     # High-level applications
│   ├── server/               # Bun + Hono Control Plane & SQLite persistence
│   │   ├── src/              # Application source code
│   │   │   ├── core/         # Core db connections, env validation
│   │   │   ├── modules/      # Domain modules (auth, tunnel, agent-token, subdomain, admin)
│   │   │   ├── app.ts        # Hono route composition
│   │   │   └── server.ts     # Bun HTTP server listener entrypoint
│   │   ├── migrations/       # Forward-only SQL migrations (0001_init.sql, etc.)
│   │   ├── scripts/          # Database migration runner script
│   │   ├── package.json      # Server dependencies
│   │   └── .env.example      # Environment variables template
│   │
│   └── client/               # React 18 + Vite SPA Developer Dashboard
│       ├── src/              # Dashboard components, pages, and context
│       │   ├── pages/        # Dashboard views (Landing, Auth, ActiveTunnels, Tokens, Admin)
│       │   ├── services/     # Typed API client functions
│       │   └── store/        # Authentication context
│       ├── vite.config.ts    # Vite bundler configuration
│       ├── package.json      # Client dependencies
│       └── .env.example      # Client environment variables template
│
├── gateway/                  # Go Ingress Gateway (High-concurrency data plane)
│   ├── cmd/gateway/          # Entrypoint binary package (main.go)
│   ├── internal/             # Private gateway modules
│   │   ├── router/           # In-memory subdomain routing table
│   │   ├── proxy/            # Public wildcard HTTP ingress listener
│   │   ├── auth/             # Agent WebSocket server & tunnel grant validation
│   │   ├── heartbeat/        # Ping/pong keepalive monitor
│   │   └── config/           # Gateway environment loader
│   └── go.mod                # Gateway Go dependencies
│
├── agent/                    # Go CLI Agent (`portalis`)
│   ├── cmd/agent/            # Entrypoint binary package (main.go)
│   ├── internal/             # Private agent modules
│   │   ├── api/              # Control Plane HTTP client (login, get grant)
│   │   ├── tunnel/           # Outbound WebSocket tunnel client & local HTTP forwarder
│   │   └── config/           # Local config file manager (~/.portalis/config.json)
│   └── go.mod                # Agent Go dependencies
│
├── frames/                   # Canonical Shared Frame Protocol
│   ├── protocol.go           # Strong Go frame models (RequestFrame, ResponseFrame, etc.)
│   ├── validation.go         # Envelope validation rules
│   ├── protocol_test.go      # Frame serialization unit tests
│   └── go.mod                # Shared protocol Go module
│
├── docs/                     # Project documentation
│   ├── TechStack.md          # Technology choices and libraries
│   ├── Architecture.md       # Master architectural specification
│   ├── Logs.md               # Logging conventions and redaction rules
│   ├── FolderStructure.md    # Directory layout and component guide (this document)
│   └── Backlog.md            # 32-issue contributor backlog
│
├── scripts/                  # Developer utility scripts
├── Makefile                  # Unified build, test, and run automation
├── CONTRIBUTING.md           # Contribution guidelines, branching model, commit standards
├── SECURITY.md               # Responsible vulnerability disclosure policy
├── LICENSE                   # MIT License
└── README.md                 # Project landing documentation
```

---

## 2. Component Responsibilities

| Directory | Responsibility | Tech Stack |
| :--- | :--- | :--- |
| **`gateway/`** | Data Plane: High-concurrency HTTP ingress and WebSocket frame multiplexer. | Go 1.23 (`net/http`, `gorilla/websocket`) |
| **`agent/`** | Client CLI: Forwards public requests to localhost ports. | Go 1.23 |
| **`frames/`** | Wire Contract: Canonical JSON frame definitions. | Go 1.23 |
| **`apps/server/`** | Control Plane: User identity, grants, quotas, and SQLite storage. | Bun + Hono + SQLite WAL |
| **`apps/client/`** | Frontend: User dashboard for tunnels, tokens, and admin metrics. | React 18 + Vite |
| **`docs/`** | Documentation: Specifications, guides, logs, and development backlog. | Markdown |
| **`.github/`** | CI/CD: Automated Go & TypeScript testing, Discord notifications. | GitHub Actions |
