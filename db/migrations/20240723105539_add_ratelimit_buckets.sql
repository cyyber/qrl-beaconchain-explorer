-- +goose Up
-- +goose StatementBegin
SELECT('up SQL query - update api_ratelimits');
ALTER TABLE api_ratelimits ADD COLUMN IF NOT EXISTS bucket TEXT NOT NULL DEFAULT 'default';
ALTER TABLE api_ratelimits DROP CONSTRAINT api_ratelimits_pkey, ADD PRIMARY KEY (user_id, bucket);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT('down SQL query - update api_ratelimits');
ALTER TABLE api_ratelimits DROP COLUMN IF EXISTS bucket;
ALTER TABLE api_ratelimits DROP CONSTRAINT api_ratelimits_pkey, ADD PRIMARY KEY (user_id);
-- +goose StatementEnd
