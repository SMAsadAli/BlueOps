# BlueOps

**BlueOps** is an enterprise-grade internal CLI built in Go for managing services, environments, deployments, and policy-driven workflows at scale.

It is designed to reflect real-world internal tooling used by large engineering organizations, focusing on security, extensibility, operational safety, and developer experience.

---

## ✨ Key Features

- 🔐 Secure authentication with token caching
- 🌍 Multi-environment configuration management
- 📦 Service discovery and metadata inspection
- 🚀 Deployment workflows with dry-run and validation
- 🧩 Plugin-based extensibility model
- 📜 Policy and guardrail enforcement
- 📊 Structured output (`--json`) for automation
- 🧠 Designed for CI/CD and internal platforms

---

## 🧠 Design Philosophy

BlueOps is intentionally **not** a simple demo CLI.

It is built to showcase:
- Production-ready CLI architecture
- Enterprise configuration patterns
- Secure secret handling
- Clear separation of concerns
- Extensibility without recompilation
- Safe operational defaults (dry-run, validation, guardrails)

---

## 🛠 Tech Stack

- **Language:** Go
- **CLI Framework:** Cobra
- **Configuration:** Viper (YAML, env vars, overrides)
- **Auth:** Token-based auth with secure OS keyring storage
- **Logging:** Structured logging (Zap / Zerolog)
- **Output:** Table + JSON modes
- **Packaging:** GoReleaser
- **Testing:** Go testing + golden tests + integration tests
