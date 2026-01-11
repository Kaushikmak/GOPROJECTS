GOSHRT
A High-Performance URL Shortener in Go
=====================================

Goshrt is a high-performance, open-source URL shortener service designed to demonstrate low-latency request routing using Go and Redis. The system combines an efficient API backend with a minimal frontend interface, optimized for high throughput and low resource overhead.

The project emphasizes performance, correctness, and simplicity, making it suitable both as a production-grade service and as a reference project for backend engineering in Go.

--------------------------------------------------------------------

OVERVIEW
--------

Goshrt provides:

- Fast URL shortening and redirection
- Sub-millisecond key lookups using Redis
- IP-based rate limiting to prevent abuse
- Optional custom aliases for shortened URLs
- Time-to-live (TTL) support for link expiry
- Fully containerized deployment using Docker

--------------------------------------------------------------------

FEATURES
--------

High Performance
Built using the Fiber web framework (powered by fasthttp) to minimize memory allocations and reduce request latency in hot paths.

In-Memory Speed
Redis is used as the primary data store, enabling extremely fast read and write operations.

Rate Limiting
IP-based throttling ensures fair usage and protects the service from excessive requests.

Custom Aliases
Users can define custom short codes instead of randomly generated identifiers.

TTL (Expiry) Support
Links can expire automatically after a configurable duration such as 24 hours, 7 days, or never.

Dockerized Architecture
The complete system, including Redis, can be deployed using Docker Compose with minimal setup.

--------------------------------------------------------------------

TECH STACK
----------

Language
Go (Golang)

Backend Framework
Fiber v2

Database
Redis

Containerization
Docker and Docker Compose

Frontend
HTML5, CSS3, Vanilla JavaScript

--------------------------------------------------------------------

PREREQUISITES
-------------

- Go version 1.19 or higher
- Docker and Docker Compose
- Redis (only if not using Docker)

--------------------------------------------------------------------

INSTALLATION AND LOCAL SETUP
----------------------------

METHOD A: DOCKER COMPOSE (RECOMMENDED)

1. git clone https://github.com/kaushikmak/goshrt.git
2. cd goshrt
3. Create .env file

DB_ADDR="localhost:6379"
DB_PASS=""
APP_PORT=":3000"
DOMAIN="localhost:3000"
API_QUOTA=10

4. docker-compose up --build

Frontend: http://localhost:3000
API: http://localhost:3000/api/v1

--------------------------------------------------------------------

METHOD B: BARE METAL

1. Start Redis
2. go mod tidy
3. go run main.go

--------------------------------------------------------------------

API USAGE
---------

POST /api/v1

{
  "url": "https://google.com",
  "customshortner": "goog",
  "expiry": "24h"
}

--------------------------------------------------------------------

LICENSE
-------

MIT License