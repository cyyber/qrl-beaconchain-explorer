-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query remove cl_proposer columns';
ALTER TABLE validator_stats DROP COLUMN IF EXISTS cl_proposer_rewards_gplanck;
ALTER TABLE validator_stats DROP COLUMN IF EXISTS cl_proposer_rewards_gplanck_total;
ALTER TABLE validator_performance DROP COLUMN IF EXISTS cl_proposer_performance_total;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query add cl_proposer columns';
ALTER TABLE validator_stats ADD COLUMN IF NOT EXISTS cl_proposer_rewards_gplanck BIGINT;
ALTER TABLE validator_stats ADD COLUMN IF NOT EXISTS cl_proposer_rewards_gplanck_total BIGINT;
ALTER TABLE validator_performance ADD COLUMN IF NOT EXISTS cl_proposer_performance_total BIGINT;
-- +goose StatementEnd
