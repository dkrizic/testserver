-- +goose Up
CREATE TABLE `asset`
(
    `id`   char(36)     NOT NULL primary key,
    `name` varchar(200) NOT NULL,
    FULLTEXT KEY `name` (`name`)
);

CREATE TABLE `file`
(
    `id`       char(36)     NOT NULL primary key,
    `name`     varchar(200) NOT NULL,
    `asset_id` char(36)     NOT NULL,
    `size`     int          NOT NULL,
    `mimetype` varchar(100) NOT NULL,
    fulltext key `name` (`name`),
    foreign key (`asset_id`) references `asset` (`id`)
);

-- +goose Down
DROP TABLE `file`;
DROP TABLE `asset`;
