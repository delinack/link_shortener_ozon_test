-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS links (
    id          uuid                     NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
    token       text                     NOT NULL,
    url         text                     NOT NULL,
    created_at  TIMESTAMP WITH TIME ZONE NOT NULL             DEFAULT now()
    );

CREATE INDEX links_token_idx ON links(token);
CREATE INDEX links_url_idx ON links(url);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS links_url_idx;
DROP INDEX IF EXISTS links_token_idx;
DROP TABLE IF EXISTS links;
-- +goose StatementEnd
