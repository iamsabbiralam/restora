-- +goose Up
-- SQL in this section is executed when the migration is applied.

CREATE TABLE IF NOT EXISTS user_role
(
    id                          VARCHAR(100)     PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id                     VARCHAR(100)     NOT NULL DEFAULT '',
    role_id                     VARCHAR(100)     NOT NULL DEFAULT '',
    created_at                  TIMESTAMP        DEFAULT current_timestamp,
    created_by                  VARCHAR(100)     NOT NULL DEFAULT '',
    updated_at                  TIMESTAMP        DEFAULT current_timestamp,
    updated_by                  VARCHAR(100)     NOT NULL DEFAULT '',
    deleted_at                  TIMESTAMP        DEFAULT NULL,
    deleted_by                  VARCHAR(100)     NOT NULL DEFAULT ''
);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE IF EXISTS user_role;