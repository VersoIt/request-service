-- +goose Up
-- +goose StatementBegin
CREATE TYPE request_status AS ENUM('pending', 'processing', 'success', 'failed');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TYPE request_status;
-- +goose StatementEnd
