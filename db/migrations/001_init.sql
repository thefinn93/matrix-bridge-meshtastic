-- +goose Up
-- +goose StatementBegin
PRAGMA foreign_keys = ON;

CREATE TABLE meshtastic_nodes (
    id INTEGER PRIMARY KEY NOT NULL,
    node_num INTEGER NOT NULL,
    meshtastic_id TEXT NOT NULL,
    long_name TEXT,
    short_name TEXT,
    mac TEXT,
    hw_model TEXT,
    public_key BLOB,
    matrix_id TEXT
);

CREATE TABLE channels (
    id INTEGER PRIMARY KEY NOT NULL,
    meshtastic_index INTEGER NOT NULL,
    meshtastic_psk TEXT UNIQUE NOT NULL,
    name TEXT,
    matrix_room TEXT
);

CREATE TABLE matrix_filters (
    user_id TEXT UNIQUE NOT NULL,
    filter_id TEXT NOT NULL
);

CREATE TABLE matrix_next_batch (
    user_id TEXT UNIQUE NOT NULL,
    token TEXT NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE channels;
DROP TABLE meshtastic_nodes;
-- +goose StatementEnd
