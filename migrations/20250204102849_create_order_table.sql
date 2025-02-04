-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS orders
(
    id                      BIGSERIAL        PRIMARY KEY,
    user_id                 varchar(128)     NOT NULL,
    ask_currency            varchar(64)      NOT NULL,
    ask_currency_amount     DECIMAL(72, 12)  NOT NULL,
    status                  varchar(128)     NOT NULL,
    created_at              TIMESTAMP        NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at              TIMESTAMP        NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at              TIMESTAMP        NULL
);

CREATE INDEX IF NOT EXISTS idx_orders_user_id ON orders (user_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS idx_orders_user_id;

DROP TABLE IF EXISTS orders;
-- +goose StatementEnd
