-- +goose NO TRANSACTION
-- +goose Up

SELECT 'create index missing indices';
-- +goose StatementBegin
CREATE INDEX CONCURRENTLY IF NOT EXISTS idx_blocks_deposits_block_slot_block_root ON public.blocks_deposits USING btree (block_slot, block_root);
-- +goose StatementEnd
-- +goose StatementBegin
CREATE INDEX CONCURRENTLY IF NOT EXISTS idx_validators_activationeligibilityepoch ON public.validators USING btree (activationeligibilityepoch);
-- +goose StatementEnd
-- +goose StatementBegin
CREATE INDEX CONCURRENTLY IF NOT EXISTS idx_sync_committees_validatorindex_period ON public.sync_committees USING btree (validatorindex, period DESC);
-- +goose StatementEnd
-- +goose StatementBegin
DROP INDEX CONCURRENTLY idx_sync_committees_period;
-- +goose StatementEnd
-- +goose StatementBegin
CREATE INDEX CONCURRENTLY IF NOT EXISTS idx_sync_committees_period ON public.sync_committees USING btree (period);
-- +goose StatementEnd
-- +goose StatementBegin
CREATE INDEX CONCURRENTLY IF NOT EXISTS idx_service_status_name_last_update ON public.service_status USING btree (name, last_update);
-- +goose StatementEnd
-- +goose StatementBegin
CREATE INDEX CONCURRENTLY IF NOT EXISTS idx_blocks_status_proposer_exec_block_number ON public.blocks USING btree (status, proposer, exec_block_number);
-- +goose StatementEnd

-- +goose Down
SELECT 'drop index missing indices';
-- +goose StatementBegin
DROP INDEX CONCURRENTLY idx_blocks_deposits_block_slot_block_root;
-- +goose StatementEnd
-- +goose StatementBegin
DROP INDEX CONCURRENTLY idx_validators_activationeligibilityepoch;
-- +goose StatementEnd
-- +goose StatementBegin
DROP INDEX CONCURRENTLY idx_sync_committees_validatorindex_period;
-- +goose StatementEnd
-- +goose StatementBegin
DROP INDEX CONCURRENTLY idx_sync_committees_period;
-- +goose StatementEnd
-- +goose StatementBegin
CREATE INDEX CONCURRENTLY IF NOT EXISTS idx_sync_committees_period ON public.sync_committees USING btree (validatorindex, period DESC);
-- +goose StatementEnd
-- +goose StatementBegin
DROP INDEX CONCURRENTLY idx_service_status_name_last_update;
-- +goose StatementEnd
-- +goose StatementBegin
DROP INDEX CONCURRENTLY idx_blocks_status_proposer_exec_block_number;
-- +goose StatementEnd
