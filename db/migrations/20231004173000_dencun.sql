// TODO(rgeraldes24)
-- +goose Up
-- +goose StatementBegin
DROP TABLE IF EXISTS blocks_transactions;
DROP TABLE IF EXISTS validator_balances_recent;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS
    blocks_transactions (
        block_slot INT NOT NULL,
        block_index INT NOT NULL,
        block_root bytea NOT NULL DEFAULT '',
        raw bytea NOT NULL,
        txhash bytea NOT NULL,
        nonce INT NOT NULL,
        gas_price bytea NOT NULL,
        gas_limit BIGINT NOT NULL,
        sender bytea NOT NULL,
        recipient bytea NOT NULL,
        amount bytea NOT NULL,
        payload bytea NOT NULL,
        max_priority_fee_per_gas BIGINT,
        max_fee_per_gas BIGINT,
        PRIMARY KEY (block_slot, block_index)
    );

CREATE TABLE IF NOT EXISTS
    validator_balances_recent (
        epoch INT NOT NULL,
        validatorindex INT NOT NULL,
        balance BIGINT NOT NULL,
        PRIMARY KEY (epoch, validatorindex)
    );
-- +goose StatementEnd