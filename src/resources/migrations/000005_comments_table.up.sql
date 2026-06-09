CREATE TABLE comments (
    id         SERIAL PRIMARY KEY,
    content    TEXT NOT NULL,
    article_id INT NOT NULL,
    user_id    INT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),

    CONSTRAINT fk_article FOREIGN KEY (article_id) REFERENCES articles(id) ON DELETE CASCADE,
    CONSTRAINT fk_user    FOREIGN KEY (user_id)    REFERENCES users(id)    ON DELETE CASCADE
);