CREATE DATABASE goload;

USE goload;

CREATE TABLE user (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    email VARCHAR(255) UNIQUE,
    username VARCHAR(100) UNIQUE,
    password_hash VARCHAR(255),         -- nullable if using OAuth
    auth_provider ENUM('google', 'local') NOT NULL DEFAULT 'local',
    google_id VARCHAR(100),             -- optional: store Google account ID
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE download_task (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    user_id BIGINT NOT NULL,
    download_type ENUM('http', 'ftp', 'torrent') NOT NULL DEFAULT 'http',
    url TEXT NOT NULL,
    metadata JSON,                           -- flexible for file info, headers, etc.
    status ENUM('queued', 'in_progress', 'completed', 'failed') NOT NULL DEFAULT 'queued',
    file_path TEXT,                          -- local/S3 path to downloaded file
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE user_token (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    user_id BIGINT NOT NULL,
    refresh_token VARCHAR(512) NOT NULL,
    user_agent VARCHAR(255),
    ip_address VARCHAR(50),
    expires_at DATETIME NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    UNIQUE(user_id, refresh_token)
);