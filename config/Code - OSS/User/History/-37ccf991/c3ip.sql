CREATE TABLE IF NOT EXISTS chats (
    id SERIAL PRIMARY KEY,
    chat_type VARCHAR(30) NOT NULL DEFAULT 'cloud private',
    title VARCHAR(255) NULL,
    created_by INT NOT NULL
);

CREATE TABLE IF NOT EXISTS chat_users (
    user_id INT,
    chat_id INT REFERENCES chats(id) ON DELETE CASCADE,
    user_role VARCHAR(20) NOT NULL DEFAULT 'member', -- 'admin' or 'member'
    joined_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (chat_id, user_id)
);