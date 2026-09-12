# Security Policy

The Portalis team takes the security of our platform and its users seriously. As a developer tunneling tool designed to expose local network services to the internet, maintaining robust security boundaries is our top priority.

---

## Supported Versions

We provide security fixes and patches for the following versions:

| Version | Supported          |
| ------- | ------------------ |
| 0.1.x   | :white_check_mark: |
| < 0.1.0 | :x:                |

---

## Reporting a Vulnerability

If you discover a potential security vulnerability within Portalis, please **do NOT report it in public GitHub issues**.

Instead, report it privately to the maintainers:
- **Email**: Contact `dineth@portalis.dev` (or via private GitHub Security Advisory).
- Include detailed steps to reproduce the vulnerability, proof-of-concept payloads where applicable, and the affected components.
- We will acknowledge receipt within 48 hours and provide a timeline for a security patch.

---

## Core Security Guidelines for Contributors

1. **Zero Secrets in Code**: Never commit `.env` files, API keys, private keys, database files, or test credentials.
2. **Short-Lived Grants**: Data plane tokens must remain short-lived. Never bypass Tunnel Grant verification.
3. **Authentication Boundary**: Never move authentication checks or database access into the Go Gateway.
