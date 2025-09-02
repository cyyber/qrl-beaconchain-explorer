-- +goose Up
-- +goose StatementBegin

SELECT('up SQL query - add msg_sender, to_address, and log_index columns to execution_deposits');
ALTER TABLE execution_deposits ADD msg_sender bytea NULL;
ALTER TABLE execution_deposits ADD to_address bytea NULL;
ALTER TABLE execution_deposits ADD log_index int4 NULL;

SELECT('up SQL query - remove duplicate rows from execution_deposits');
delete from
	execution_deposits
where
	merkletree_index in (
	select
		merkletree_index
	from
		(
		select
			merkletree_index,
			row_number() over (partition by merkletree_index) as rn,
			COUNT(*) over (partition by merkletree_index) as cnt
		from
			execution_deposits
) t
	where
		t.cnt > 1
);

SELECT('up SQL query - changing the primary key of execution_deposits to be merkletree_index alone');
ALTER TABLE execution_deposits DROP CONSTRAINT IF EXISTS execution_deposits_pkey;
ALTER TABLE execution_deposits ADD PRIMARY KEY (merkletree_index);

SELECT('up SQL query - add block_number index to execution_deposits');
CREATE INDEX execution_deposits_block_number_idx ON execution_deposits (block_number);


-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

SELECT('down SQL query - remove msg_sender, to_address, and log_index columns from execution_deposits');
ALTER TABLE execution_deposits DROP COLUMN msg_sender;
ALTER TABLE execution_deposits DROP COLUMN to_address;
ALTER TABLE execution_deposits DROP COLUMN log_index;

SELECT('down SQL query - changing the primary key of execution_deposits back to be tx_hash & merkletree_index');
ALTER TABLE execution_deposits DROP CONSTRAINT IF EXISTS execution_deposits_pkey;
ALTER TABLE execution_deposits ADD PRIMARY KEY (tx_hash, merkletree_index);

SELECT('down SQL query - remove block_number index from execution_deposits');
DROP INDEX IF EXISTS execution_deposits_block_number_idx;

-- +goose StatementEnd
