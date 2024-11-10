CREATE TYPE difficulty AS ENUM ('beginner', 'easy', 'medium', 'hard', 'extreme' );
CREATE TYPE state AS ENUM ('opening', 'midgame', 'endgame', 'unknown');
CREATE TYPE board_tile AS ENUM ('X', 'O', '');

CREATE TABLE games (
  game_id uuid DEFAULT gen_random_uuid() PRIMARY KEY,
  name text NOT NULL,

  game_difficulty difficulty NOT NULL,
  game_state state NOT NULL,

  board board_tile[15][15] CHECK (cardinality(board) = 225),

  created timestamp NOT NULL DEFAULT now(),
  updated timestamp NOT NULL DEFAULT now()
);
