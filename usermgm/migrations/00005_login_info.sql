-- +goose Up
-- SQL in this section is executed when the migration is applied.

CREATE TABLE IF NOT EXISTS login
(

    id                  SERIAL PRIMARY  KEY,
    user_id             VARCHAR(100)    NOT NULL DEFAULT '',
    login_info          TEXT            NULL DEFAULT '{}',
    created_at          TIMESTAMP       DEFAULT current_timestamp
);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE IF EXISTS login;
