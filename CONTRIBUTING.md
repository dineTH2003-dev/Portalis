# Contributing to Portalis

Guidelines for setting up your local environment, running the services, and submitting pull requests.

---

## Contribution Workflow

> [!IMPORTANT]
> **All active development targets the `dev` branch, not `main`.**  
> `main` is our stable release branch. Pull requests opened against `main` will be asked to re-target `dev`.

### 1. Fork and Clone
Fork the repository on GitHub, then clone your fork locally:
```bash
git clone https://github.com/<your-username>/Portalis.git
cd Portalis
```

### 2. Add Upstream Remote
Keep your local clone synchronized with the original repository:
```bash
git remote add upstream https://github.com/dineTH2003-dev/Portalis.git
```

### 3. Sync Before Coding
Ensure your local `dev` branch is completely up to date with `upstream/dev`:
```bash
git checkout dev
git pull upstream dev
```

### 4. Create a Branch
Branch off `dev` using descriptive prefixes:
```bash
git checkout -b feat/your-feature-name
# or
git checkout -b fix/your-bugfix-name
```
*Naming convention: `feat/<issue-number>-<name>`, `fix/<issue-number>-<name>`, `docs/<name>`, `refactor/<name>`.*

### 5. Commit Your Changes
Follow [Conventional Commits](https://www.conventionalcommits.org/) format:
```bash
git commit -m "feat(gateway): implement in-memory subdomain routing table"
```
Allowed prefixes: `feat:`, `fix:`, `refactor:`, `docs:`, `test:`, `chore:`.

### 6. Sync Again and Resolve Conflicts
Fetch upstream updates before opening a PR:
```bash
git fetch upstream
git rebase upstream/dev
```
If conflicts occur, resolve them locally, stage the files, and run:
```bash
git rebase --continue
```

### 7. Verify Build and Tests (Zero-Failing-Builds Rule)
Before opening a PR, ensure all local verification checks pass:
```bash
# Run all unit tests across Go and Bun
make test

# Verify Go code formatting and TypeScript types
make lint

# Verify production builds
make build
```

### 8. Submit a Pull Request
Push to your fork and open a Pull Request targeting the original repository's **`dev`** branch:
```bash
git push -u origin feat/your-feature-name
```
Complete the [Pull Request Template](.github/PULL_REQUEST_TEMPLATE.md) and link the issue (e.g. `Closes #5`).

---

## Prerequisites

- **Go**: `>= 1.23` — [go.dev](https://go.dev/)
- **Bun**: `>= 1.1` — [bun.sh](https://bun.sh/)
- **Node.js**: `>= 18.0` — [nodejs.org](https://nodejs.org/)
- **Git**: Latest version — [git-scm.com](https://git-scm.com/)

---

## Quickstart (Local Development Flow)

### 1. Install Dependencies
```bash
# Install Control Plane dependencies
cd apps/server && bun install

# Install Dashboard dependencies
cd ../client && bun install
```

### 2. Configure Environment
Copy the template files:
```bash
# Control Plane
cp apps/server/.env.example apps/server/.env

# React Dashboard
cp apps/client/.env.example apps/client/.env
```

### 3. Run Database Migrations
Initialize the embedded SQLite database in WAL mode:
```bash
cd apps/server && bun run migrate
```

### 4. Start Development Servers
From the project root using `make`:
```bash
# Terminal 1: Start Control Plane (:4310)
make run-server

# Terminal 2: Start Go Ingress Gateway (:8080 / :9000)
make run-gateway

# Terminal 3: Start React Dashboard (:5310)
make run-client
```

---

## Default Development Endpoints & Ports

When running locally, Portalis exposes the following services:

| Subsystem | Port / Endpoint | Description |
| :--- | :--- | :--- |
| **Control Plane API** | `http://localhost:4310` | REST API, auth, quotas, and SQLite storage |
| **Ingress Gateway (HTTP)** | `http://localhost:8080` | Public wildcard ingress (`*.localhost:8080`) |
| **Ingress Gateway (WS)** | `ws://localhost:9000/v1/tunnel/ws` | Persistent agent tunnel termination |
| **React Dashboard** | `http://localhost:5310` | Developer portal and admin console |

---

## Development Guidelines & Policies

1. **Zero-Secret Policy**: Never commit `.env`, private keys, API tokens, or `.db` databases into git.
2. **Separation of Concerns**: Keep Control Plane business logic in `apps/server` and high-speed network proxying in `gateway`.
3. **Atomic PRs**: Keep pull requests focused on a single issue from [`docs/Backlog.md`](docs/Backlog.md).
4. **Architecture Grounding**: Review [`docs/Architecture.md`](docs/Architecture.md) to understand platform boundaries.
