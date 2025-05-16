# Simple Golang Clinic Web Application

## Features

- Single login endpoint for both Receptionist and Doctor.
- Receptionist can register, view, update, and delete patients.
- Doctor can view and update patient details.
- JWT authentication.
- Uses Gin, Gorm, Postgres, Viper.
- Simple and easy to understand.

## Setup

1. **Clone the repo and navigate to the folder.**

2. **Install Go dependencies**
   ```bash
   go mod tidy
   ```

3. **Set up PostgreSQL database**
   - Create a database `clinicdb` and user/password as in `.env`, or update `.env` accordingly.

4. **Set up initial users**
   - Insert user(s) into the users table manually, e.g. using `psql`:

     ```sql
     INSERT INTO users (username, password, role, created_at, updated_at)
     VALUES ('reception1', '<bcrypt_hash>', 'receptionist', now(), now());

     INSERT INTO users (username, password, role, created_at, updated_at)
     VALUES ('doctor1', '<bcrypt_hash>', 'doctor', now(), now());
     ```

     To get `<bcrypt_hash>`, use the `utils.HashPassword` function in Go, or use an online bcrypt generator.

5. **Run the application**
   ```bash
   go run main.go
   ```

## Endpoints

- `POST /api/login`  
  **Request:** `{ "username": "...", "password": "..." }`  
  **Response:** `{ "token": "...", "role": "...", "expires": "..." }`
- `POST /api/patients` (Receptionist only)
- `GET /api/patients` (Both)
- `GET /api/patients/:id` (Both)
- `PUT /api/patients/:id` (Both)
- `DELETE /api/patients/:id` (Receptionist only)

All endpoints except `/api/login` require `Authorization: Bearer <token>` header.

## Configuration

Edit `.env` for DB, JWT secret.

