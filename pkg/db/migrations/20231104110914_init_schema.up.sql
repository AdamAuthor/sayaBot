CREATE TABLE message_summaries (
    id SERIAL PRIMARY KEY,
    chat_id BIGINT NOT NULL,
    message_text TEXT NOT NULL
);