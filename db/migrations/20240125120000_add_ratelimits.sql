-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query - add table api_ratelimits';
CREATE TABLE IF NOT EXISTS
    api_ratelimits (
        user_id INT NOT NULL,
        second INT NOT NULL DEFAULT 0,
        hour INT NOT NULL DEFAULT 0,
        month INT NOT NULL DEFAULT 0,
        valid_until TIMESTAMP WITHOUT TIME ZONE NOT NULL,
        changed_at TIMESTAMP WITHOUT TIME ZONE NOT NULL, 
        PRIMARY KEY (user_id)
    );

CREATE INDEX IF NOT EXISTS idx_api_ratelimits_changed_at_valid_until ON api_ratelimits (changed_at, valid_until);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query - drop table api_ratelimits';
DROP TABLE IF EXISTS api_ratelimits;
SELECT 'down SQL query - drop index idx_api_ratelimits_changed_at';
DROP INDEX IF EXISTS idx_api_ratelimits_changed_at;
-- +goose StatementEnd
