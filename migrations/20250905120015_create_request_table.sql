-- +goose Up
-- +goose StatementBegin
CREATE TABLE requests
(
    id         SERIAL PRIMARY KEY,
    payload    BYTEA,
    status     request_status NOT NULL,
    type       request_type   NOT NULL,
    created_at TIMESTAMP,
    user_id INTEGER NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
