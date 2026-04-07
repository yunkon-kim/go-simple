---
applyTo: "**/*.go"
---

# Go Coding Standards

## Style & Conventions

- **Style:** Follow **Effective Go** and **Go Code Review Comments**.
- **Formatting:** Use `gofmt` or `goimports`.
- **Linting:** Ensure code passes `golangci-lint` (run `make lint`).

## Logging & Error Handling

- **Logging:** Use `zerolog`.
  - Pattern: `log.Error().Err(err).Msg("message")`
- **Error Handling:**
  - Handle errors explicitly.
