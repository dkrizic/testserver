-- +goose Up
CREATE TABLE `asset`
(
    `id`   char(36)         NOT NULL primary key,
    `name` varchar(200) NOT NULL,
    FULLTEXT KEY `name` (`name`)
);

-- +goose Down
DROP TABLE `asset`;
