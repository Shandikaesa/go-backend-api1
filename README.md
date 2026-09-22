# Go Backend API

REST API sederhana dengan Go, PostgreSQL, dan Docker.

## Fitur

- GET `/users` - Lihat semua user
- POST `/users/add` - Tambah user baru
- PostgreSQL sebagai database
- Docker & Docker Compose
- Nginx reverse proxy

## Teknologi

- **Go** (net/http, database/sql)
- **PostgreSQL** 16
- **Docker** & **Docker Compose**
- **Nginx** (reverse proxy)

## Cara Jalanin

### 1. Clone repository
```bash
git clone https://github.com/Shandikaesa/go-backend-api.git
cd go-backend-api
### 2. Bikin file .env
DB_USER=sanzdev
DB_PASSWORD=rahasia123
DB_NAME=belajar_go
DB_HOST=db
DB_PORT=5432
### 3. jalanin docker 
docker compose up -d
### 4. Test API
curl http://localhost:8081/users
### 5. .
├── db.go              # Koneksi database
├── server.go          # HTTP server & endpoint
├── Dockerfile         # Build image Go
├── docker-compose.yml # Orkestrasi API + DB
└── scripts/           # Script deploy, backup, health


### AUTHOR

** SANZDEV ** - [Github](https://github.com/shandikaesa)
