CREATE TABLE articles (
    id         SERIAL PRIMARY KEY,
    title      VARCHAR(255) NOT NULL,
    content    TEXT NOT NULL,
    view_count INT DEFAULT 0,
    active     BOOLEAN DEFAULT true,
    created_at TIMESTAMP DEFAULT NOW(),
    user_id    INT NOT NULL,

    CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);
