# go-simple

A simple Go web server that displays your IP address.

## Project Structure

```
go-simple/
├── cmd/server/          # Application entry point
├── internal/ip/         # IP detection utilities (private)
├── web/templates/       # HTML templates
└── .github/instructions # Copilot context files (deactivated)
```

Based on [Standard Go Project Layout](https://github.com/golang-standards/project-layout).

## Quick Start

```bash
go run cmd/server/main.go
```

Open http://localhost:8080 in your browser.

## API Endpoints

| Method | Path             | Description               |
| ------ | ---------------- | ------------------------- |
| GET    | `/`              | Main page (HTML)          |
| GET    | `/api/my-ip`     | Visitor's IP (JSON)       |
| GET    | `/api/server-ip` | Server's public IP (JSON) |

## References

- IP inquiry: [golang-playground/inquire-public-ip](https://github.com/yunkon-kim/golang-playground/blob/master/inquire-public-ip/inquire-public-ip.go)
