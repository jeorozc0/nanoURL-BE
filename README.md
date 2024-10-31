# Go URL Shortener Service

A high-performance URL shortener service built with Go, featuring PostgreSQL persistence, configurable CORS, and production-ready deployment configuration for fly.io.

## Table of Contents
- [Features](#features)
- [Tech Stack](#tech-stack)
- [Prerequisites](#prerequisites)
- [Local Development](#local-development)
- [API Documentation](#api-documentation)
- [Project Structure](#project-structure)
- [Deployment](#deployment)
- [Environment Variables](#environment-variables)
- [Database](#database)
- [Monitoring](#monitoring)
- [Security](#security)
- [Contributing](#contributing)
- [License](#license)

## Features

- ⚡ Fast URL shortening using base62 encoding
- 🔒 Persistent storage with PostgreSQL
- 📝 RESTful API endpoints
- 🚀 Production-ready deployment configuration
- 🔑 Configurable CORS for multiple origins
- 📊 Request logging middleware
- ⌛ Graceful shutdown handling
- 🛡️ Security best practices
- ⚙️ Environment-based configuration
- 📈 Performance optimized

## Tech Stack

- Go 1.22+
- PostgreSQL
- fly.io for deployment
- CORS middleware (rs/cors)
- UUID generation (google/uuid)
- Connection pooling
- Graceful shutdown handling

## Prerequisites

- Go 1.22 or higher
- PostgreSQL
- [flyctl CLI](https://fly.io/docs/hands-on/install-flyctl/)
- Git

## Local Development

1. **Clone the repository**
   ```bash
   git clone https://github.com/yourusername/url-shortener.git
   cd url-shortener
   ```

2. **Install dependencies**
   ```bash
   go mod download
   ```

3. **Set up environment variables**
   ```bash
   export DATABASE_URL="postgresql://username:password@localhost:5432/urlshortener?sslmode=disable"
   export ALLOWED_ORIGINS="http://localhost:5173,http://localhost:3000"
   export CORS_DEBUG="true"
   ```

4. **Create and initialize database**
   ```bash
   createdb urlshortener
   ```

5. **Run the application**
   ```bash
   go run main.go
   ```
   Server will start on `http://localhost:8080`

6. **Run tests**
   ```bash
   go test ./...
   ```

## API Documentation

### Create Short URL
**POST** `/url`

Request:
```json
{
  "url": "https://example.com/very-long-url"
}
```

Response:
```json
{
  "short_url": "http://yourdomain.fly.dev/Ab3x"
}
```

### Get Original URL
**GET** `/{id}`

Response:
- Success: Original URL as JSON string
- Not Found: 404 status code
- Error: 500 status code

## Project Structure
```
├── db/
│   └── db.go           # Database connection and initialization
├── handlers/
│   ├── getHandler.go   # GET endpoint handler
│   ├── postHandler.go  # POST endpoint handler
│   └── rootHandler.go  # Root endpoint handler
├── middleware/
│   └── logging.go      # Logging middleware
├── models/
│   └── urlModel.go     # URL data model and database operations
├── services/
│   └── base62.go       # Base62 encoding service
├── Dockerfile          # Container configuration
├── fly.toml            # fly.io configuration
├── go.mod             
├── go.sum
├── main.go            # Application entry point
└── README.md
```

## Deployment

### Initial Setup
1. **Install flyctl and login**
   ```bash
   curl -L https://fly.io/install.sh | sh
   fly auth login
   ```

2. **Initialize fly.io project**
   ```bash
   fly launch
   ```

3. **Create PostgreSQL database**
   ```bash
   fly postgres create --name your-app-db
   fly postgres attach your-app-db
   ```

### Environment Configuration
1. **Set required secrets**
   ```bash
   fly secrets set ALLOWED_ORIGINS="https://your-production-frontend.com"
   fly secrets set CORS_DEBUG="false"
   ```

2. **Deploy the application**
   ```bash
   fly deploy
   ```

### Scaling (Optional)
```bash
# Set the number of instances
fly scale count 2

# Configure machine size
fly scale vm shared-cpu-1x
```

## Environment Variables

### Required Variables
- `DATABASE_URL`: PostgreSQL connection string (automatically set by fly.io)

### Optional Variables
- `PORT`: Application port (default: 8080)
- `ALLOWED_ORIGINS`: Comma-separated list of allowed CORS origins 
  - Default: "http://localhost:5173,http://localhost:3000"
  - Example: "http://localhost:5173,https://your-production-frontend.com"
- `CORS_DEBUG`: Enable CORS debug logging (default: "false")

### Local Configuration Example
```bash
export DATABASE_URL="postgresql://username:password@localhost:5432/urlshortener?sslmode=disable"
export ALLOWED_ORIGINS="http://localhost:5173,http://localhost:3000"
export CORS_DEBUG="true"
```

## Database

### Schema
```sql
CREATE TABLE urls (
    id VARCHAR(4) PRIMARY KEY,
    original_url TEXT NOT NULL,
    new_url TEXT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);
```

### Connection Configuration
- Connection pooling enabled
- Configurable timeouts
- Automatic reconnection handling
- Prepared statement support

## Monitoring

### View Application Logs
```bash
fly logs
```

### Check Application Status
```bash
fly status
```

### Database Management
```bash
# Connect to PostgreSQL console
fly postgres connect -a your-app-db

# Monitor database metrics
fly postgres metrics -a your-app-db
```

### Performance Monitoring
- Request logging with duration tracking
- Database query monitoring
- Error rate tracking
- Response time metrics

## Security

### CORS Configuration
- Configurable allowed origins
- Secure default headers
- Credentials handling
- Preflight request support

### Best Practices
- Input validation
- Prepared statements for SQL
- Rate limiting support
- Secure headers
- TLS enforcement
- Graceful error handling

## Error Handling

The service implements comprehensive error handling:
- Invalid URLs return 400 Bad Request
- Not found URLs return 404 Not Found
- Server errors return 500 Internal Server Error
- All errors are logged with timestamps and request IDs
- Structured error responses

## Performance Considerations

- Efficient base62 encoding
- Database connection pooling
- Proper mutex handling
- Response caching capability
- Optimized database queries
- Memory usage optimization

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

### Development Guidelines
- Follow Go best practices and idioms
- Maintain test coverage
- Document public functions
- Use meaningful commit messages
- Update documentation as needed

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Support

For support:
- Open an issue in the GitHub repository
- Check existing documentation
- Contact the maintenance team

## Acknowledgments

- [fly.io](https://fly.io) for deployment platform
- [PostgreSQL](https://www.postgresql.org/) for database
- All contributors and maintainers

---

Developed with ❤️ by Julio
