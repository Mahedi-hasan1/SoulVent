# SoulVent

SoulVent is an anonymous social platform built with Golang. This main service handles posts, feeds, user profiles, and engagement features using PostgreSQL and Elasticsearch. It is designed for high concurrency, scalability, and privacy.

## Project Structure

- `cmd/` - Main application entry point
- `internal/app/` - Application logic and server setup
- `internal/db/` - Database and Elasticsearch integration
- `internal/model/` - Data models (User, Post, Comment)
- `internal/handler/` - HTTP handlers for posts, users, comments
- `internal/service/` - Business logic (feed generation, moderation)
- `internal/middleware/` - JWT authentication, rate limiting
- `internal/util/` - Utilities (S3 upload, logging)
- `config/` - Configuration files
- `api/` - API specification (OpenAPI)
- `deployments/` - Deployment files (Docker Compose)
- `scripts/` - Database initialization and schema
- `docs/` - Documentation
- `pkg/` - Reusable packages (if needed)

## Features
- Anonymous posting and engagement
- JWT authentication
- Rate limiting
- Efficient feed generation
- Image uploads via S3
- Robust moderation tools
- PostgreSQL and Elasticsearch integration
- High concurrency and scalability

## Getting Started
1. Configure environment in `config/config.yaml`
2. Build and run with Docker Compose (`deployments/docker-compose.yaml`)
3. Initialize database with scripts in `scripts/`
4. See API spec in `api/openapi.yaml`

---

What’s in your mind right now? Say it. No one knows you. No one will ever know. Just let it go.Where every thought is a whisper from the soul, and every post is a breath released into the void .
