CREATE TABLE IF NOT EXISTS games (
    uuid TEXT NOT NULL PRIMARY KEY,
    created_at datetime NOT NULL,
    updated_at datetime NOT NULL,
    name TEXT NOT NULL,
    difficulty TEXT NOT NULL,
    game_state TEXT NOT NULL,
    board BLOB NOT NULL
);

CREATE TABLE IF NOT EXISTS users (
    uuid TEXT NOT NULL PRIMARY KEY,
    created_at datetime NOT NULL,
    username TEXT NOT NULL UNIQUE,
    email TEXT NOT NULL UNIQUE,
    password TEXT NOT NULL,
    elo INTEGER NOT NULL,
    wins INTEGER NOT NULL DEFAULT 0,
    draws INTEGER NOT NULL DEFAULT 0,
    losses INTEGER NOT NULL DEFAULT 0,
    banned INTEGER NOT NULL DEFAULT FALSE
);

CREATE UNIQUE INDEX IF NOT EXISTS idx_mails ON users (email);
CREATE UNIQUE INDEX IF NOT EXISTS idx_usernames ON users (username);
