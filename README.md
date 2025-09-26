# RSS Web Scraper

A robust RSS feed aggregator and web scraper built in Go that automatically fetches and stores RSS feed content with a RESTful API for managing feeds and users.

## Features

- **RSS Feed Parsing**: Automatically fetches and parses RSS feeds from various sources
- **Concurrent Scraping**: Multi-threaded feed fetching with configurable concurrency
- **User Management**: Create and manage users with API key authentication
- **Feed Management**: Add, retrieve, and manage RSS feeds
- **Feed Following**: Users can follow/unfollow specific feeds
- **Post Aggregation**: Automatically stores and retrieves posts from followed feeds
- **RESTful API**: Complete HTTP API for all operations
- **PostgreSQL Integration**: Persistent storage with proper database schema
- **CORS Support**: Cross-origin resource sharing enabled

## Tech Stack

- **Language**: Go 1.24.4
- **Database**: PostgreSQL
- **HTTP Router**: Chi v5
- **Database Migrations**: Goose
- **Code Generation**: SQLC
- **Environment Management**: godotenv

## Project Structure

```
webscrapper/
├── internal/
│   ├── auth/           # Authentication utilities
│   └── database/       # Generated database code (SQLC)
├── sql/
│   ├── queries/        # SQL queries for SQLC
│   └── schema/         # Database migration files
├── vendor/             # Go modules cache
├── main.go            # Application entry point
├── rss.go             # RSS feed parsing logic
├── scrapper.go        # Background scraping worker
├── models.go          # Data models and converters
├── handler*.go        # HTTP request handlers
├── middlewear.go      # Authentication middleware
├── json.go            # JSON response utilities
├── .env               # Environment variables
└── go.mod             # Go module definition
```

## Installation

### Prerequisites

- Go 1.24.4 or higher
- PostgreSQL database
- Git

### Setup

1. **Clone the repository**
   ```bash
   git clone <repository-url>
   cd webscrapper
   ```

2. **Install dependencies**
   ```bash
   go mod download
   ```

3. **Set up PostgreSQL database**
   - Create a PostgreSQL database named `rssagg`
   - Update the connection string in `.env` file

4. **Configure environment variables**
   ```bash
   cp .env.example .env
   ```
   Edit `.env` with your configuration:
   ```
   PORT=8080
   DB_URL="postgres://username:password@localhost:5432/rssagg?sslmode=disable"
   ```

5. **Run database migrations**
   ```bash
   # Install goose if not already installed
   go install github.com/pressly/goose/v3/cmd/goose@latest
   
   # Run migrations
   goose -dir sql/schema postgres "your-db-connection-string" up
   ```

6. **Build and run**
   ```bash
   go build -o webscrapper
   ./webscrapper
   ```

## API Endpoints

### Health Check
- `GET /v1/healthz` - Health check endpoint
- `GET /v1/error` - Error testing endpoint

### User Management
- `POST /v1/users` - Create a new user
- `GET /v1/users` - Get current user info (requires authentication)

### Feed Management
- `POST /v1/feed` - Create a new RSS feed (requires authentication)
- `GET /v1/feed` - Get all available feeds

### Feed Following
- `POST /v1/feed_follow` - Follow a feed (requires authentication)
- `GET /v1/feed_follow` - Get user's followed feeds (requires authentication)
- `DELETE /v1/feed_follow/{feedfollowid}` - Unfollow a feed (requires authentication)

### Posts
- `GET /v1/upost` - Get posts from followed feeds (requires authentication)

## Authentication

The API uses API key authentication. Include your API key in the `Authorization` header:

```
Authorization: ApiKey your-api-key-here
```

## Usage Examples

### Create a User
```bash
curl -X POST http://localhost:8080/v1/users \
  -H "Content-Type: application/json" \
  -d '{"name": "John Doe"}'
```

### Add an RSS Feed
```bash
curl -X POST http://localhost:8080/v1/feed \
  -H "Content-Type: application/json" \
  -H "Authorization: ApiKey your-api-key" \
  -d '{"name": "Tech Blog", "url": "https://example.com/rss.xml"}'
```

### Follow a Feed
```bash
curl -X POST http://localhost:8080/v1/feed_follow \
  -H "Content-Type: application/json" \
  -H "Authorization: ApiKey your-api-key" \
  -d '{"feed_id": "feed-uuid-here"}'
```

### Get Your Posts
```bash
curl -X GET http://localhost:8080/v1/upost \
  -H "Authorization: ApiKey your-api-key"
```

## Background Scraping

The application automatically scrapes RSS feeds in the background:

- **Concurrency**: 10 concurrent workers by default
- **Interval**: Scrapes every minute
- **Duplicate Handling**: Automatically skips duplicate posts
- **Error Handling**: Continues operation even if individual feeds fail

## Database Schema

The application uses the following main tables:

- **users**: User accounts with API keys
- **feeds**: RSS feed sources
- **feed_follows**: User-feed relationships
- **posts**: Individual RSS feed items

## Development

### Code Generation

This project uses SQLC for type-safe database operations:

```bash
# Generate database code
sqlc generate
```

### Database Migrations

Add new migrations in `sql/schema/` following the naming convention:
```
XXX_description.sql
```

Run migrations:
```bash
goose -dir sql/schema postgres "connection-string" up
```

## Configuration

### Environment Variables

- `PORT`: Server port (default: 8080)
- `DB_URL`: PostgreSQL connection string

### Scraper Configuration

Modify scraping behavior in `main.go`:
```go
go scrappig(dbcon, 10, time.Minute) // 10 workers, 1-minute interval
```

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests if applicable
5. Submit a pull request

## License

This project is open source and available under the [MIT License](LICENSE).

## Troubleshooting

### Common Issues

1. **Database Connection Failed**
   - Verify PostgreSQL is running
   - Check connection string in `.env`
   - Ensure database exists

2. **Port Already in Use**
   - Change the PORT in `.env`
   - Kill existing processes on the port

3. **RSS Feed Not Updating**
   - Check feed URL is valid
   - Verify feed follows XML RSS format
   - Check application logs for errors

### Logs

The application provides detailed logging for:
- Server startup
- Feed scraping operations
- Database operations
- HTTP requests

Monitor logs to troubleshoot issues and track application performance.