CREATE TABLE IF NOT EXISTS users_tokens (
    id bigserial PRIMARY KEY,
    hash bytea NOT NULL,
    user_id INT NOT NULL,
    type text NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id),
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    updated_at timestamp(0) with time zone
);
