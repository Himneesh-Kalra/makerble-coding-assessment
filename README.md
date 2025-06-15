# Hospital Management System

A simple Golang-based RESTful API for managing patients in a Hospital setting with role-based access for Receptionists and Doctors.

## Tech Stack

* **Golang** with Gin
* **PostgreSQL**
* **GORM** ORM
* **Docker** & Docker Compose
* **JWT** authentication

---

## Features

* **Login & Registration** with JWT auth
* **Receptionists** can:
* Register new patients
* Perform full CRUD operations on patients
* **Doctors** can:
* View and update patient records
* Role-based middleware to protect endpoints

---

## Prerequisites

Before you begin, ensure you have the following installed:

* [Go 1.20+](https://golang.org/doc/install)
* [Docker](https://www.docker.com/products/docker-desktop)
* [Postman](https://www.postman.com/downloads/) or similar API client

---

## Environment Variables

Create a `.env` file in the root directory and configure the following variables:

```
DB_HOST=localhost
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=hospital_db
DB_PORT=5432
JWT_SECRET=verysecret
```

Or define these variables in your `docker-compose.yml`.

---

## Setup Instructions

### 1. Clone the Repo

```bash
git clone https://github.com/your-username/clinic-api.git
cd clinic-api
```

### 2. Update `.env` (optional)

Use environment variables or modify `docker-compose.yml` accordingly.

### 3. Start PostgreSQL via Docker Compose

```bash
docker-compose up -d
```

### 4. Run Migrations (Auto handled if coded)

OR use GORM auto migration in `main.go`.

### 5. Run the Server

```bash
go run main.go
```

API will start at `http://localhost:8080`

---

## Sample API Usage (via Postman)

### Register

**POST** `/api/register`

```json
{
  "name": "Dr. John",
  "email": "john@example.com",
  "password": "yourpassword",
  "role": "doctor" // or receptionist
}
```

### Login

**POST** `/api/login`

```json
{
  "email": "john@example.com",
  "password": "yourpassword"
}
```

### Auth Header (Required for Protected Routes)

```
Authorization: Bearer <your-jwt-token>
```

---

## Patient Endpoints

### Create Patient

**POST** `/api/patients`

```json
{
  "first_name": "Alice",
  "last_name": "Johnson",
  "age": 28,
  "gender": "female",
  "diagnosis": "Fever"
}
```

### Update Patient

**PUT** \`/ap
