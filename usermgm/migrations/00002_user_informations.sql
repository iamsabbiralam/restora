-- +goose Up
-- SQL in this section is executed when the migration is applied.
-- This has to be run at one time. no need to include the follwoing command into the other file
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS user_information
(

    id                          VARCHAR(100)    PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id                     VARCHAR(100)    UNIQUE NOT NULL DEFAULT '',
    first_name                  VARCHAR(150)    NOT NULL DEFAULT '',
    last_name                   VARCHAR(150)    NOT NULL DEFAULT '',
    mobile                      VARCHAR(20)     UNIQUE NOT NULL DEFAULT '',
    gender                      SMALLINT        DEFAULT 0,
    dob                         TIMESTAMP       DEFAULT current_timestamp,
    address                     VARCHAR(150)    NOT NULL DEFAULT '',
    city                        VARCHAR(100)    NOT NULL DEFAULT '',
    country                     VARCHAR(100)    NOT NULL DEFAULT '',
    created_at                  TIMESTAMP       DEFAULT current_timestamp,
    created_by                  VARCHAR(100)    NOT NULL DEFAULT '',
    updated_at                  TIMESTAMP       DEFAULT current_timestamp,
    updated_by                  VARCHAR(100)    NOT NULL DEFAULT '',
    deleted_at                  TIMESTAMP       DEFAULT NULL,
    deleted_by                  VARCHAR(100)    NOT NULL DEFAULT ''
);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE IF EXISTS user_information;
