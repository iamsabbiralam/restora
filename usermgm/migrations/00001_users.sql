-- +goose Up
-- SQL in this section is executed when the migration is applied.
-- This has to be run at one time. no need to include the follwoing command into the other file
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS users
(

    id                          VARCHAR(100)    PRIMARY KEY DEFAULT uuid_generate_v4(),
    username                    VARCHAR(100)    UNIQUE NOT NULL DEFAULT '',
    email                       VARCHAR(150)    UNIQUE NOT NULL DEFAULT '',
    password                    VARCHAR(255)    NOT NULL DEFAULT '',
    status                      SMALLINT        DEFAULT 0,
    is_mfa                      BOOLEAN         DEFAULT FALSE,
    mfa_type                    VARCHAR(255)    DEFAULT 'EMAIL',
    created_at                  TIMESTAMP       DEFAULT current_timestamp,
    created_by                  VARCHAR(100)    NOT NULL DEFAULT '',
    updated_at                  TIMESTAMP       DEFAULT current_timestamp,
    updated_by                  VARCHAR(100)    NOT NULL DEFAULT '',
    deleted_at                  TIMESTAMP       DEFAULT NULL,
    deleted_by                  VARCHAR(100)    NOT NULL DEFAULT ''
);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE IF EXISTS users;
