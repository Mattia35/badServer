# BadServer
Academic and educational project designed to study and experiment with SQL Injection vulnerabilities in a controlled environment.

⚠️ **Security Notice**
This repository contains **deliberately insecure code**. It is intended **exclusively for academic, educational, and research purposes** (e.g., university projects, cybersecurity laboratories, OWASP demonstrations). **Do not deploy this software in production environments or expose it to public networks.**

## Abstract

**BadServer** is a client–server application implemented in **Go** with a **MySQL** backend, intentionally designed to violate common secure coding practices. The project provides a realistic environment for analyzing **SQL Injection vulnerabilities** in RESTful APIs, including the impact of unsafe query construction, missing input validation, and insecure database configuration options.

## Possible Attacks

This project is intentionally designed to support the study and practical experimentation of SQL Injection attack techniques. The backend implementation deliberately enables multiple classes of attacks by allowing user-controlled input to directly influence SQL query execution.

In particular, the application supports the analysis of in-band SQL Injection attacks, including:
* Tautology-based injections, where conditional statements are manipulated to always evaluate to true, enabling authentication bypass and unrestricted data access.
* End-of-line comment injections, which exploit SQL comment syntax to truncate legitimate query logic and ignore security-relevant conditions.
* Piggybacked queries, made possible by unsafe query concatenation and the explicit use of multi-statement execution, allowing attackers to append and execute additional malicious SQL commands.

In addition to in-band techniques, the system is also suitable for studying inferential SQL Injection attacks (Illicit or unauthorized queries), where sensitive information is extracted indirectly through the application’s behavior rather than direct query output.

## System Architecture

The application follows a classical client–server model:

```text
Client (Browser / HTTP tools)
        │
        ▼
┌────────────────────────┐
│   HTTP Server (Go)     │
│  net/http + Router    │
└──────────┬────────────┘
           │
           ▼
┌────────────────────────┐
│    MySQL Database      │
│  Dynamic SQL Queries   │
└────────────────────────┘
```

* **Client**: static frontend or external HTTP clients (e.g., curl, Postman)
* **Server**: REST API implemented in Go
* **Database**: MySQL with automatically initialized schema

## Technology Stack

* **Programming Language**: Go
* **Database**: MySQL
* **HTTP Framework**: `net/http`
* **Router**: `github.com/julienschmidt/httprouter`
* **Database Driver**: `github.com/go-sql-driver/mysql`
* **Logging**: `github.com/sirupsen/logrus`

## Project Structure

```text
badServer/
├── backend/                  # API layer and Database layer
├── frontend/                 # Static client files
├── main.go                   # Application entry point
└── README.md
```

## How to Run the Project

This section describes how to set up and run the application locally for educational and experimental purposes.

### Prerequisites

Ensure the following components are installed on your system:

* **Go** (version 1.20 or higher)
* **MySQL** (version 8.x recommended)
* A Unix-like operating system or Windows with WSL

### Database Setup

The application is designed to run entirely on a single host for academic and experimental purposes. In this configuration, the client, the backend server, and the database server all execute on the same machine, while preserving a clear logical separation between components.

The backend HTTP server is explicitly configured in main.go to listen on the local loopback interface at:
```code
localhost:8080
```
The database server is a MySQL-compatible system (MariaDB) running locally on the same host and accessed via a TCP connection.

### Database Server Requirements

A local installation of MariaDB (MySQL-compatible) is required. During development and testing, the database server runs on the loopback interface and listens on the standard MySQL/MariaDB port.

This setup ensures that:

* no external network communication is involved
* all database interactions can be inspected locally
* the environment remains isolated and reproducible

### Creating the Database and Dedicated User

Database initialization is performed manually using the MariaDB client, which acts as a command-line interface to the local database server.

Access the database client using an administrative account:

```bash
mysql -u root -p
```
Once connected, execute the following SQL commands:
```sql
CREATE DATABASE badserver;

CREATE USER 'baduser'@'localhost' IDENTIFIED BY 'badpass';
GRANT ALL PRIVILEGES ON badserver.* TO 'baduser'@'localhost';
FLUSH PRIVILEGES;
```

The application connects to the database using the following Data Source Name (DSN):

```go
baduser:badpass@tcp(127.0.0.1:3306)/badserver?parseTime=true&multiStatements=true
```

> **Note**: The `multiStatements=true` option is intentionally enabled to facilitate advanced SQL Injection scenarios.

### Application Startup

Clone the repository and start the server:

```bash
git clone https://github.com/Mattia35/badServer.git
cd badServer
go run main.go
```

On startup, the application will:

1. Establish a connection to the MySQL database
2. Automatically create the required tables
3. Start the HTTP server on `localhost:8080`

### Accessing the Application

Once the server is running, the backend services can be accessed through the following base URL:

```
http://localhost:8080
```

