-- name: MeshtasticNodeUpdate :exec
INSERT INTO meshtastic_nodes (node_num, meshtastic_id, long_name, short_name, hw_model, public_key) VALUES (?, ?, ?, ?, ?, ?) ON CONFLICT (node_num) DO UPDATE SET
    meshtastic_id = excluded.meshtastic_id,
    long_name = excluded.long_name,
    short_name = excluded.short_name,
    hw_model = excluded.hw_model,
    public_key = excluded.public_key;

-- name: MeshtasticNodeGet :one
SELECT * FROM meshtastic_nodes WHERE node_num = ? LIMIT 1;
