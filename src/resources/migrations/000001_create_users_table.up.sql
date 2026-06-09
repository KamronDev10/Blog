CREATE TABLE users (
    id            SERIAL PRIMARY KEY,
    username      VARCHAR(50) UNIQUE NOT NULL,
    email         VARCHAR(100) UNIQUE NOT NULL,
    password      VARCHAR(255) NOT NULL,
    avatar        VARCHAR(255),
    bio           TEXT,
    registered_at TIMESTAMP DEFAULT NOW()
);
