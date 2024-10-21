-- name: MeshtasticNodeUpdate :exec
INSERT INTO meshtastic_nodes (node_num, meshtastic_id, long_name, short_name, mac, hw_model, public_key) VALUES (?, ?, ?, ?, ?, ?, ?) ON CONFLICT (node_num) DO UPDATE SET
    meshtastic_id = excluded.meshtastic_id,
    long_name = excluded.long_name,
    short_name = excluded.short_name,
    mac = excluded.mac,
    hw_model = excluded.hw_model,
    public_key = excluded.public_key;
