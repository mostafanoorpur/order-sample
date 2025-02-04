-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS balances
(
    id               BIGSERIAL       PRIMARY KEY,
    user_id          varchar(128)    NOT NULL,
    currency         varchar(64)     NOT NULL,
    amount           DECIMAL(72, 12) NOT NULL,
    created_at       TIMESTAMP       NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at       TIMESTAMP       NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_balances_user_id  ON balances (user_id);
CREATE INDEX IF NOT EXISTS idx_balances_currency ON balances (currency);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS idx_balances_user_id;
DROP INDEX IF EXISTS idx_balances_currency;

DROP TABLE IF EXISTS balances;
-- +goose StatementEnd
