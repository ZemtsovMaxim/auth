CREATE TABLE IF NOT EXISTS users
(
    id        SERIAL PRIMARY KEY,
    email     TEXT NOT NULL UNIQUE,
    pass_hash bytea NOT NULL
);

CREATE INDEX IF NOT EXISTS idx_email ON users (email);

CREATE TABLE IF NOT EXISTS secrets
(
    id        SERIAL PRIMARY KEY,
    secret    TEXT NOT NULL UNIQUE
);

INSERT INTO secrets (id, secret)
VALUES (1, 'test-secret')
ON CONFLICT DO NOTHING;
