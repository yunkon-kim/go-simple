# go-simple Copilot Instructions

## Purpose
- Keep AI-generated changes small, correct, and easy to review.
- Preserve the current architecture and API behavior unless a task explicitly requests a change.

## Project Context
- Language: Go (standard library HTTP server).
- App type: Web server that shows visitor IP and server public IP.
- Main entry: `cmd/server/main.go`.
- Internal logic: `internal/ip/ip.go`.
- UI template: `web/templates/index.html`.

## Source of Truth
- Follow language-specific rules in `.github/instructions/go.instructions.md` for Go files.
- If rules conflict, prefer:
  1. User request
  2. This file
  3. File-scoped instructions

## Architecture and Boundaries
- Keep request routing and HTTP concerns in `cmd/server/`.
- Keep IP-related parsing and lookup logic in `internal/ip/`.
- Do not move internal package code to public packages unless explicitly requested.
- Keep HTML, CSS, and browser-side JS in `web/templates/`.

## API and Behavior Rules
- Preserve existing endpoints unless the task asks otherwise:
  - `GET /`
  - `GET /api/my-ip`
  - `GET /api/server-ip`
- JSON responses must remain simple and stable:
  - success payload uses `{"ip":"..."}`
  - error payload uses `{"error":"..."}` when applicable
- Set `Content-Type: application/json` for JSON APIs.
- Return explicit HTTP status codes for error cases.

## Coding Conventions
- Prefer standard library solutions before adding external dependencies.
- Keep handlers thin; move reusable logic into `internal/` packages.
- Keep functions small and intention-revealing.
- Handle errors explicitly and return early on failure.
- Avoid broad refactors in feature or bug-fix tasks.

## Frontend Conventions
- Keep template updates minimal and readable.
- Maintain responsive layout behavior for desktop and mobile.
- Avoid introducing large UI frameworks for small changes.

## Documentation and Change Hygiene
- Update `README.md` when routes, startup commands, or externally visible behavior changes.
- Keep comments concise and only for non-obvious intent.
- Do not modify unrelated files.

## Validation Checklist
- Run `go test ./...` when Go logic changes.
- Manually verify key routes when server or template behavior changes.
- Ensure generated code is `gofmt`-compatible.
