
CREATE TABLE game_players (
    game_id UUID NOT NULL REFERENCES games(id) ON DELETE CASCADE,
    player_id UUID NOT NULL,
    cell SMALLINT NOT NULL CHECK (cell IN (1, 2)),
    is_computer BOOLEAN NOT NULL DEFAULT FALSE,

    PRIMARY KEY (game_id, player_id),
    UNIQUE (game_id, cell)
);

CREATE INDEX idx_games_status ON games(status);
