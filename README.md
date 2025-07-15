# Flashlight Backend

A Go-based REST API backend for managing student data, built as a take-home project for Flashlight Learning. This application provides endpoints for creating and retrieving student information with PostgreSQL as the database.

## 🚀 Live Demo

The application is deployed and running on [Railway](https://railway.app/) (https://flashlight-frontend-production.up.railway.app).

## 🛠️ Tech Stack

- **Language**: Go 1.23.0+
- **Web Framework**: [Gin](https://gin-gonic.com/) - High-performance HTTP web framework
- **Database**: PostgreSQL with [GORM](https://gorm.io/) ORM
- **Environment Management**: [godotenv](https://github.com/joho/godotenv)
- **CORS**: [gin-contrib/cors](https://github.com/gin-contrib/cors)
- **Deployment**: Railway with Nixpacks

## 📋 API Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/api/students/all` | Retrieve all students |
| POST | `/api/students/create` | Create a new student |

### Student Model
```go
type Student struct {
    StudentID uint   `json:"studentId"`
    Name      string `json:"name"`
    Grade     uint   `json:"grade"`
}
```

### Create Student Request Body
```json
{
    "name": "John Doe",
    "grade": 10
}
```

## 🏗️ Project Structure

```
flashlight-backend/
├── cmd/
│   └── main.go              # Application entry point
├── internal/
│   ├── handlers/
│   │   └── student.go       # HTTP request handlers
│   ├── models/
│   │   └── student.go       # Data models
│   ├── routes/
│   │   └── students.go      # Route definitions
│   └── server/
│       └── server.go        # Server configuration
├── go.mod                   # Go module dependencies
├── go.sum                   # Go module checksums
├── Makefile                 # Build and run commands
├── nixpacks.toml           # Railway deployment config
└── README.md               # Project documentation
```

## 🚀 Local Setup

### Prerequisites

- Go 1.23.0 or later
- PostgreSQL database
- Git

### Installation Steps

1. **Clone the repository**
   ```bash
   git clone https://github.com/1ShoukR/flashlight-backend
   cd flashlight-backend
   ```

2. **Install dependencies**
   ```bash
   go mod download
   ```

3. **Set up environment variables**
   ```bash
   cp .env.example .env
   ```
   
   Edit the `.env` file with your configuration:
   ```env
   PORT=8080
   DSN=postgres://username:password@localhost:5432/flashlight_db?sslmode=disable
   ```

4. **Set up PostgreSQL Database**
   
   Create a PostgreSQL database:
   ```sql
   CREATE DATABASE flashlight_students;
   ```

5. **Run the application**
   ```bash
   make run
   ```
   
   The server will start on the port specified in your `.env` file (default: 8080).

## 🔧 Makefile Commands

The project includes a Makefile with the following commands:

| Command | Description |
|---------|-------------|
| `make all` | Alias for `make build` |
| `make build` | Build the application binary as `main.exe` |
| `make run` | Run the application directly without building |

### Usage Examples

```bash
# Build the application
make build

# Run the application
make run

# Build everything (same as build)
make all
```

## 🌍 Environment Variables

| Variable | Description | Required | Default |
|----------|-------------|----------|---------|
| `PORT` | Port number for the server | Yes | 8080 |
| `DSN` | PostgreSQL connection string | Yes | - |

### DSN Format
```
postgres://username:password@hostname:port/database_name?sslmode=disable
```

## 🗄️ Database

The application uses PostgreSQL with GORM for ORM functionality. The database schema is automatically migrated on startup.

### Student Table Schema
```sql
CREATE TABLE students (
    student_id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    grade INTEGER NOT NULL
);
```

## 🌐 CORS Configuration

The application is configured to accept requests from:
- `http://localhost:3000`
- `http://localhost:5173`
- `http://127.0.0.1:3000`
- `http://127.0.0.1:5173`
- `https://flashlight-frontend-production.up.railway.app`

## 🚀 Deployment

The application is deployed on Railway using Nixpacks. The deployment configuration is specified in `nixpacks.toml`:

```toml
[phases.build]
cmd = "go build -o main ./cmd/"

[phases.start]
cmd = "./main"
```

## 🧪 Testing the API

### Get All Students
```bash
curl -X GET http://localhost:8080/api/students/all
```

### Create a Student
```bash
curl -X POST http://localhost:8080/api/students/create \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Jane Doe",
    "grade": 12
  }'
```

## 📦 Dependencies

Main dependencies:
- `github.com/gin-gonic/gin` - Web framework
- `gorm.io/gorm` - ORM library
- `gorm.io/driver/postgres` - PostgreSQL driver
- `github.com/joho/godotenv` - Environment variable loader
- `github.com/gin-contrib/cors` - CORS middleware

## 🚀 Future Features
- **Delete Students** - Add route to handle deleting students 
- **Create Classrooms** - Create classrooms and assign students via foreign key 
- **Different Auth Levels** - Add auth levels to differentiate between staff and students and users

- **Assignments** - Add ability to upload and process PDFs, word docs, etc for student assignments
- **Caching** -- Add caching ability for reduced API costs (this would probably be best with Redis)
- **Auth Protected Routes** -- Protect important routes with auth protection


