-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query - add qrns lookup tables';
CREATE TABLE IF NOT EXISTS
    qrns (
        name_hash bytea NOT NULL,
        qrns_name TEXT NOT NULL,
        address bytea,
        is_primary_name BOOLEAN NOT NULL DEFAULT FALSE,
        valid_to TIMESTAMP WITHOUT TIME ZONE,
        PRIMARY KEY (name_hash)
    );
CREATE INDEX IF NOT EXISTS idx_qrns_name ON qrns (qrns_name);
CREATE INDEX IF NOT EXISTS idx_qrns_valid_name ON qrns (qrns_name, valid_to);
CREATE INDEX IF NOT EXISTS idx_qrns_valid_address_primary ON qrns (address, valid_to, is_primary_name);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query - remove qrns lookup tables';
DROP INDEX IF EXISTS idx_qrns_name;
DROP INDEX IF EXISTS idx_qrns_valid_name;
DROP INDEX IF EXISTS idx_qrns_valid_address_primary;
DROP TABLE IF EXISTS qrns CASCADE;
-- +goose StatementEnd
