
CREATE TABLE games (
    id UUID PRIMARY KEY,
    board JSONB NOT NULL,
    current_player UUID NOT NULL,
    status SMALLINT NOT NULL,
    winner UUID,
    mode SMALLINT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
