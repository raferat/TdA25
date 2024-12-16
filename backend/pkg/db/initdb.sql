CREATE TABLE IF NOT EXISTS games (
  uuid TEXT NOT NULL PRIMARY KEY,
  created_at datetime NOT NULL,
  updated_at datetime NOT NULL,
  name TEXT NOT NULL,
  difficulty TEXT NOT NULL,
  game_state TEXT NOT NULL,
  board BLOB NOT NULL
);
