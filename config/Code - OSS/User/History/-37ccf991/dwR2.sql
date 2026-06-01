CREATE TABLE chats (
    id SERIAL PRIMARY KEY,
    chat_type TEXT NOT NULL DEFAULT "cloud private",
    title TEXT NULL,
    created_by BIGINT 
)