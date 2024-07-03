-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS store(
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    description TEXT NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS store;
-- +goose StatementEnd