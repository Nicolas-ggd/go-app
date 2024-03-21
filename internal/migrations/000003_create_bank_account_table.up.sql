CREATE TABLE IF NOT EXISTS user_bank_account (
    id bigserial PRIMARY KEY,
    name text NOT NULL,
    balance integer NOT NULL,
    fees integer NOT NULL,
    user_id INT NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id),
    updated_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    deleted_at timestamp(0) with time zone
);