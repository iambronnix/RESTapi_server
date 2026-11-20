# Go REST API Server

A robust, production-ready REST API server built from scratch in Go. This project emphasizes clean architecture, secure routing, and concurrency-safe request handling without relying on heavy external web frameworks.

## 🚀 Features

*   **Custom Router (`Methods` map):** Implements the `http.Handler` interface to dynamically route requests based on HTTP methods, with proper `405 Method Not Allowed` and `OPTIONS` handling.
*   **Layered Middleware Architecture:**
    *   **Authentication & Authorization:** Simulates a token-based auth flow with Role-Based Access Control (RBAC), ensuring only authorized "admin" roles can access specific endpoints.
    *   **URL Inspection:** Custom middleware acts as a secondary multiplexer, inspecting URL paths to route requests to the correct domain handlers (GET, POST, PUT, DELETE).
    *   **Prefix Restriction:** A security middleware that scans and blocks requests attempting to access restricted path prefixes.
*   **Concurrency Safe:** Utilizes `sync/atomic` for safe request counting across multiple goroutines, preventing data races under high load.
*   **Structured Logging:** Uses Go's modern `log/slog` package (Go 1.21+) for high-performance, structured request logging.
*   **Graceful Request Cleanup:** Implements `io.Copy(io.Discard, r.Body)` and explicit body closing to prevent memory leaks and ensure connection reuse.
*   **TLS Support:** Configured to serve over HTTPS using `ServeTLS`, ensuring data in transit is encrypted.

## 🏗️ Architecture

The project is structured to separate routing, middleware, and business logic:

```text
├── api/
│   ├── DELETE/
│   │   ├── order/
│   │   └── user/
│   ├── GET/
│   │   ├── order/
│   │   ├── payment/
│   │   └── user/
│   ├── POST/
│   │   ├── order/
│   │   ├── payment/
│   │   └── user/
│   ├── PUT/
│   │   ├── order/
│   │   ├── payment/
│   │   └── user/
│   └── middleware/
│       └── middleware.go
├── server.go
└── README.md
