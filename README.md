# FUTO Marching Dashboard

A dashboard application for managing schedules and tasks for the FUTO Marching Band.

## Features

- User authentication (login/logout)
- User management (creation, editing, deletion)
- Role management (admin and general users)
- Calendar functionality
- Task management
- Practice menu management
- Time tracking

## Technology Stack

### Backend
- Go 1.24.2
- Echo Web Framework
- MongoDB

### Frontend
- Next.js
- Tailwind CSS
- TypeScript

## Getting Started

### Prerequisites
- Docker and Docker Compose
- Node.js 18+
- Go 1.24+
- MongoDB

### Running with Docker Compose

```bash
# Clone the repository
git clone https://github.com/kynmh69/futo-marching-dashboad.git
cd futo-marching-dashboad

# Start the application
docker-compose up -d
```

The application will be available at:
- Frontend: http://localhost:3000
- Backend API: http://localhost:8080

### Development Setup

#### Backend
```bash
cd backend
cp .env.example .env  # Configure your environment variables
go mod download
go run cmd/server/main.go
```

#### Frontend
```bash
cd frontend
npm install
npm run dev
```

## Testing

The repository includes unit tests for both backend and frontend components.

### Backend Tests
```bash
cd backend
go test -v ./...
```

### Frontend Tests
```bash
cd frontend
npm test
```

## GitHub Actions

This project uses GitHub Actions for CI/CD. Tests are automatically run on pull requests and pushes to the main branch.