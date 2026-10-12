
CREATE TABLE users (
    id UUID PRIMARY KEY,
    login TEXT NOT NULL UNIQUE,
    password_hash TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);


-- логин сделал уникальным чтобы пользователи не были с одинаковым логинм