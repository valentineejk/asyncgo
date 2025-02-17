-- Users Table
CREATE TABLE users (
    id VARCHAR(21) PRIMARY KEY,
    email VARCHAR(320) NOT NULL UNIQUE,
    hashed_password VARCHAR(96) NOT NULL, -- bcrypt base64 encoded
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Refresh Tokens Table
CREATE TABLE refresh_tokens (
    user_id VARCHAR(21) REFERENCES users(id) ON DELETE CASCADE,
    hashed_token VARCHAR(500) NOT NULL UNIQUE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    expires_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP + INTERVAL '1 day',
    PRIMARY KEY (user_id, hashed_token)
);

-- Reports Table
CREATE TABLE reports (
    id VARCHAR(21) REFERENCES users(id) ON DELETE CASCADE,
    user_id VARCHAR(21) REFERENCES users(id) ON DELETE CASCADE,
    report_type VARCHAR NOT NULL,
    output_file_path VARCHAR,
    download_url VARCHAR,
    download_url_expires_at TIMESTAMPTZ,
    error_message VARCHAR,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    started_at TIMESTAMPTZ,
    failed_at TIMESTAMPTZ,
    completed_at TIMESTAMPTZ,
    PRIMARY KEY (user_id, id)
);