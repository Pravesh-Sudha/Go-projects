# 3 Go-lang Projects

Welcome to the **3 Go-lang Projects** repository! This collection contains three beginner-to-intermediate-level Go projects tailored for DevOps enthusiasts and programmers learning Go. These projects demonstrate practical applications of Go for system monitoring, web services, and RESTful API development.

---

## Projects Overview

### 1. **Monitor Disk Usage**
- **Description**: A CLI tool that displays disk usage statistics for a specified directory, including total, used, and free space.
- **Key Features**:
  - Takes a directory path as a command-line argument.
  - Outputs disk usage details in a human-readable format.
  - Includes error handling for invalid paths.
- **Usage**:
  ```bash
  go run disk_usage.go /path/to/directory
  ```

---

### 2. **HTTP Server with Health Check**
- **Description**: A simple HTTP server with a health endpoint and a modern `index.html` page that showcases the author's profile.
- **Key Features**:
  - Serves a static HTML page (`index.html`) with personal information (name, title, profile picture).
  - Includes a health check endpoint (`/health`) to monitor server status.
  - Demonstrates serving static files in Go.
- **Endpoints**:
  - `GET /`: Displays the `index.html` page.
  - `GET /health`: Returns a JSON response with server health status.
- **Usage**:
  ```bash
  go run main.go
  ```
  Visit `http://localhost:8080` in your browser.

---

### 3. **RESTful API for Tools**
- **Description**: A RESTful API for managing a collection of DevOps tools.
- **Key Features**:
  - Implements basic CRUD operations.
  - Uses JSON for request and response payloads.
  - Demonstrates struct handling and routing in Go.
- **Endpoints**:
  - `GET /tools`: Retrieves all tools.
  - `POST /tools/add`: Adds a new tool (expects JSON payload).
  - `GET /tools/details`: Retrieves details of a specific tool based on a query parameter.
- **Example**:
  ```bash
  curl -X POST -H "Content-Type: application/json" \
  -d '{"name":"Docker","description":"A platform to build, run, and share containerized applications.","category":"Containerization"}' \
  http://localhost:8080/tools/add
  ```

---

## Repository Structure
```
3-go-lang-project/
├── monitor_disk_usage/
│   └── disk_usage.go
├── http_server_with_health/
│   ├── main.go
│   └── index.html
├── restful_api/
│   ├── main.go
│   ├── data.go
│   └── handlers.go
└── README.md
```

---

## Getting Started

### Prerequisites
- Go installed (v1.16+ recommended)
- A terminal or IDE for running and debugging Go programs

### Clone the Repository
```bash
git clone https://github.com/your-username/3-go-lang-project.git
cd 3-go-lang-project
```

### Running a Project
Navigate to the specific project folder and execute the respective Go program:
```bash
cd monitor_disk_usage
go run disk_usage.go /path/to/directory
```

---

## Author
**Pravesh Sudha**  
- 💼 DevOps Engineer | Freelancer at Fiverr  
- ✍️ Tech Blogger on [Hashnode](https://hashnode.com/@praveshsudha) and [Medium](https://medium.com/@praveshsudha)  
- 🌐 [Personal Website](https://praveshsudha.com)  

---

## Contributing
Feel free to submit issues or fork the repository and submit pull requests to contribute new features or improvements.

---

## License
This project is open-source and available under the [MIT License](LICENSE).

---

Let me know if you'd like to customize this further! 🚀
