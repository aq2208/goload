# 🚀 GoLoad – Distributed Download Manager in Golang

GoLoad is a high-performance, extensible, and self-hostable file download server inspired by [PyLoad](https://pyload.net/) and Internet Download Manager (IDM).  
Built using Go, Kafka, and modern cloud-native tools, GoLoad allows users to submit file URLs and download them asynchronously — with real-time tracking and secure access.

---

## ✨ Features

- 🔁 **Asynchronous Background Downloads**  
  Submit links and let the system download in the background via Kafka workers.

- 📡 **gRPC + HTTP Streaming**  
  Built with `grpc-gateway` for efficient RESTful API and support for large-file streaming.

- 🔐 **Authentication & Authorization**  
  Secure access via Google OAuth or custom user/password. Users only access their own files.

- ⚙️ **Resilient Task Execution**  
  Dual mechanism for download jobs: passive (Kafka consumers) + active (cron fallback).

- ☁️ **Flexible Storage Backends**  
  Choose between local file storage or S3-compatible object storage (e.g. MinIO).

- 📊 **Download Tracking**  
  Monitor real-time progress and access downloaded files via the HTTP interface.

---

## 🧱 System Architecture

| Component     | Description                                                                 |
|---------------|-----------------------------------------------------------------------------|
| **Go HTTP Server** | API server using `net/http` and `grpc-gateway` for REST + streaming |
| **Kafka**         | Message queue for handling background download jobs                    |
| **MySQL**         | Stores user metadata and file records                                   |
| **Redis**         | Speeds up common queries and stores session/cache data                 |
| **Cron Job**      | Periodic backup worker for download jobs (fallback in case Kafka fails)|
| **S3 / MinIO**    | Blob storage for downloaded files                                       |

---

## 📦 Tech Stack

- **Language:** Go (Golang)
- **API Protocols:** gRPC + grpc-gateway (JSON-RPC over HTTP)
- **Message Queue:** Apache Kafka
- **Database:** MySQL
- **Cache Layer:** Redis
- **Object Storage:** AWS S3 / MinIO (self-hosted S3-compatible)
- **Authentication:** Google OAuth2, Username/Password

---

## 🚧 Work in Progress

- [ ] File categorization & tagging
- [ ] Download retries with exponential backoff
- [ ] UI dashboard (React-based frontend)
- [ ] Role-based access control (RBAC)

---

## 📂 Project Structure

```bash
.
├── cmd/                  # Entry points
├── internal/             # Core logic
│   ├── downloader/       # Download manager workers
│   ├── storage/          # File storage handlers (S3, local)
│   ├── auth/             # OAuth2 and session management
│   └── db/               # MySQL + Redis integration
├── proto/                # Protobuf definitions
├── api/                  # gRPC-Gateway handlers
├── config/               # App config files
├── scripts/              # Kafka setup, DB migrations
└── README.md
```

---

## 🛠️ Getting Started

- **Prerequisites**: Go 1.20+, Docker (for Kafka, Redis, MySQL, MinIO)

```bash
# Clone the repo
git clone https://github.com/yourname/goload.git
cd goload

# Run using Docker Compose (WIP)
docker-compose up

# OR run locally
go run cmd/server/main.go
```


