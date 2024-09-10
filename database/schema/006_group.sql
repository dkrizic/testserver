-- +goose Up
CREATE TABLE `group` (
    id char(36) PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

-- +goose Down
DROP TABLE `group`;
