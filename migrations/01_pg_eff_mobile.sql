-- +goose Up
-- +goose StatementBegin

-- Persons
CREATE TABLE IF NOT EXISTS persons (
    id          INTEGER PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    name        VARCHAR(128) NOT NULL,
    patronymic  VARCHAR(128),
    surname     VARCHAR(128) NOT NULL,
    age         INTEGER NOT NULL,
    gender      VARCHAR(8) NOT NULL,
    nationality VARCHAR(128),
    created_at  TIMESTAMP NOT NULL DEFAULT now(),
    updated_at  TIMESTAMP NOT NULL DEFAULT now(),
    CONSTRAINT chck_gender_type CHECK (
        gender IN ('male', 'female')
    )
);
-- +goose StatementEnd

-- +goose Down

-- +goose StatementBegin
DROP TABLE IF EXISTS persons;
-- +goose StatementEnd