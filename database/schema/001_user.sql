-- +goose Up
CREATE TABLE user (
    id char(36) not null PRIMARY KEY,
    email VARCHAR(255) NOT NULL
);

-- +goose Down
DROP TABLE user;
