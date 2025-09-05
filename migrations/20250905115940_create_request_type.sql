-- +goose Up
-- +goose StatementBegin
CREATE TYPE request_type AS ENUM ('certificate', 'passport_application', 'tex_declaration', 'license_request', 'social_benefit', 'registration_update');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TYPE request_type;
-- +goose StatementEnd
