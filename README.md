# Course Management System API

A backend REST API for managing online courses, built with **Go and Gin**.  
The system supports **user authentication, role-based access control, course organization, and student enrollments**.

The platform organizes educational content hierarchically:

```
Course → Sections → Topics
```

Admins can create and manage course content, while students can enroll and access courses.

---

## Features

### User Management
- User registration and login
- JWT token authentication
- Role-based authorization (Admin / Student)

### Course Organization
- Hierarchical structure: **Courses → Sections → Topics**
- CRUD operations for all entities
- Document attachment support

### Learning Experience
- Course enrollment system
- Student dashboard (view enrolled courses)
- Multiple content types (text, video, PDF)

---

## Technology Stack

| Component | Technology |
|----------|-----------|
| Language | Go |
| Framework | Gin |
| ORM | GORM |
| Database | SQLite |
| Authentication | JWT |

---

## Project Architecture

The project follows a **Clean Architecture structure** separating application layers.

```
internal/
  entity/        # Data models
  repository/    # Database layer
  usecase/       # Business logic
  handler/       # HTTP API controllers

main.go          # Application entry point
```

---

## Running the Project

### 1. Clone the repository

```bash
git clone https://github.com/RajbirW/course-system-api-public.git
cd course-system-api-public
```

### 2. Install dependencies

```bash
go mod tidy
```

### 3. Run the server

```bash
go run main.go
```

The API will start on:

```
http://localhost:8080
```

---

## API Endpoints

### Authentication

#### Register User

```http
POST /register
Content-Type: application/json
```

```json
{
  "username": "student1",
  "password": "password123",
  "role": "student"
}
```

#### Login

```http
POST /login
Content-Type: application/json
```

```json
{
  "username": "student1",
  "password": "password123"
}
```

---

### Courses

#### Create Course (Admin only)

```http
POST /courses
Authorization: Bearer <token>
Content-Type: application/json
```

```json
{
  "title": "Advanced Go",
  "description": "Master Go programming",
  "instructor": "Jane Doe"
}
```

#### Get All Courses

```http
GET /courses
Authorization: Bearer <token>
```

---

### Sections

#### Create Section (Admin only)

```http
POST /courses/{id}/sections
Authorization: Bearer <token>
Content-Type: application/json
```

```json
{
  "title": "Concurrency in Go",
  "course_id": 1
}
```

---

### Topics

#### Create Topic (Admin only)

```http
POST /sections/{id}/topics
Authorization: Bearer <token>
Content-Type: application/json
```

```json
{
  "title": "Goroutines",
  "content": "Lightweight threads in Go",
  "type": "text",
  "section_id": 1
}
```

---

### Enrollments

#### Enroll in Course

```http
POST /courses/{id}/enroll
Authorization: Bearer <student_token>
```

```json
{
  "course_id": 1
}
```

#### Get Enrolled Courses

```http
GET /enrollments
Authorization: Bearer <student_token>
```

---

## Database Schema

```
USER
├── ID
├── Username
├── Password
├── Token
└── Role

COURSE
├── ID
├── Title
├── Description
├── Instructor
└── Sections

SECTION
├── ID
├── Title
└── CourseID

TOPIC
├── ID
├── Title
├── Content
├── Type
└── SectionID

ENROLLMENT
├── ID
├── UserID
└── CourseID

DOCUMENT
├── ID
├── Filename
├── Path
├── UserID
└── ParentID
```

---

## Error Handling

### 400 Bad Request
```json
{
  "error": "Invalid input data"
}
```

### 401 Unauthorized
```json
{
  "error": "Invalid or expired token"
}
```

### 403 Forbidden
```json
{
  "error": "Admin access required"
}
```

### 404 Not Found
```json
{
  "error": "Course not found"
}
```

### 409 Conflict
```json
{
  "error": "User already enrolled in this course"
}
```