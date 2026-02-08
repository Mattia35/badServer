BadServer

An Intentionally Vulnerable Client–Server Application for SQL Injection Analysis

Academic and educational project designed to study and experiment with SQL Injection vulnerabilities in a controlled environment.

⚠️ Security Notice
This repository contains deliberately insecure code. It is intended exclusively for academic, educational, and research purposes (e.g., university projects, cybersecurity laboratories, OWASP demonstrations). Do not deploy this software in production environments or expose it to public networks.

⸻

Abstract

BadServer is a client–server application implemented in Go with a MySQL backend, intentionally designed to violate common secure coding practices. The project provides a realistic environment for analyzing SQL Injection vulnerabilities in RESTful APIs, including the impact of unsafe query construction, missing input validation, and insecure database configuration options.

The goal is to support hands-on learning in courses related to:
	•	Cybersecurity
	•	Secure Software Engineering
	•	Web Application Security
	•	OWASP Top 10 vulnerabilities

⸻

System Architecture

The application follows a classical client–server model:

Client (Browser / HTTP tools)
        │
        ▼
┌────────────────────────┐
│   HTTP Server (Go)     │
│  net/http + Router    │
└──────────┬────────────┘n           │
           ▼
┌────────────────────────┐n│    MySQL Database      │n│  Dynamic SQL Queries  │n└────────────────────────┘n```

- **Client**: static frontend or external HTTP clients (e.g., curl, Postman)
- **Server**: REST API implemented in Go
- **Database**: MySQL with automatically initialized schema

---

## Technology Stack

- **Programming Language**: Go
- **Database**: MySQL
- **HTTP Framework**: `net/http`
- **Router**: `github.com/julienschmidt/httprouter`
- **Database Driver**: `github.com/go-sql-driver/mysql`
- **Logging**: `github.com/sirupsen/logrus`

---

## Project Structure

badServer/
├── backend/
│   ├── api/                  # HTTP handlers (intentionally vulnerable)
│   └── api/requestContext/   # Request-scoped context and logging
│
├── frontend/                 # Static client files
├── main.go                   # Application entry point
└── README.md

---

## Database Configuration

The database connection is intentionally configured with insecure options:

```go
baduser:badpass@tcp(127.0.0.1:3306)/badserver?parseTime=true&multiStatements=true

The parameter multiStatements=true is deliberately enabled to facilitate advanced SQL Injection scenarios.

Database Initialization

The schema is created automatically at server startup and includes the following tables:
	•	profile
	•	token
	•	department
	•	project
	•	employee

The schema uses real foreign key constraints, including cyclic dependencies, to enable complex attack scenarios.

⸻

API Endpoints

Authentication

PUT /login

Handles user authentication. The implementation intentionally lacks robust security mechanisms.

⸻

Employee Data

GET /profiles/:profile/employees

Returns employee information associated with a profile.

Security note: user-controlled parameters are directly interpolated into SQL queries.

⸻

Project Search

GET /profiles/:profile/projects

Designed to support experimentation with:
	•	Boolean-based SQL Injection
	•	UNION-based SQL Injection
	•	Error-based SQL Injection

⸻

Department Data

GET /profiles/:profile/departments

Retrieves department information.

⸻

Department Modification

PUT /profiles/:profile/departments/:department

Updates department data.

Critical endpoint: suitable for demonstrating destructive SQL Injection attacks (UPDATE, INSERT, DROP).

⸻

Intended Vulnerabilities

This project intentionally violates multiple secure coding principles:
	•	Use of dynamic SQL built via string concatenation
	•	Absence of prepared statements
	•	Lack of input validation and sanitization
	•	Hard-coded database credentials
	•	Overly permissive CORS configuration
	•	Weak authentication and session handling

These weaknesses are introduced by design for educational analysis.

⸻

Educational Use Cases

The repository can be used for:
	•	University laboratory assignments
	•	Secure coding and penetration testing courses
	•	OWASP Top 10 demonstrations
	•	Controlled Red Team / Blue Team exercises

⸻

Ethical and Legal Disclaimer

This software is provided solely for educational and research purposes.

Any attempt to use the techniques demonstrated here against systems without explicit authorization may be illegal and unethical. The authors assume no liability for misuse of this code.

⸻

Future Work

Possible extensions include:
	•	Secure refactoring using prepared statements
	•	Side-by-side comparison between vulnerable and hardened implementations
	•	Automated test cases for vulnerability detection
	•	Integration into Capture The Flag (CTF) environments

⸻

License

This project is intended for academic use. Add an explicit open-source license (e.g., MIT, Apache 2.0) if you plan to redistribute or extend it.
