CREATE TABLE books
(
    id         UUID PRIMARY KEY,
    title      TEXT NOT NULL,
    author     TEXT NOT NULL,
    year       INTEGER,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW()
);
