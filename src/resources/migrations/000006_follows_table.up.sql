CREATE TABLE follows (
    follower_id  INT NOT NULL,
    following_id INT NOT NULL,

    PRIMARY KEY (follower_id, following_id),

    CONSTRAINT fk_follower  FOREIGN KEY (follower_id)  REFERENCES users(id) ON DELETE CASCADE,
    CONSTRAINT fk_following FOREIGN KEY (following_id) REFERENCES users(id) ON DELETE CASCADE
);