-- +goose Up
CREATE TABLE `tagvalue`
(
    `id`       int NOT NULL AUTO_INCREMENT,
    `asset_id` int NOT NULL,
    `tag_id`   int NOT NULL,
    `value`    varchar(200) DEFAULT NULL,
    PRIMARY KEY (`id`),
    KEY        `asset_id` (`asset_id`),
    KEY        `tag_id` (`tag_id`),
    FULLTEXT KEY `value` (`value`),
    CONSTRAINT `tagvalue_ibfk_1` FOREIGN KEY (`asset_id`) REFERENCES `asset` (`id`),
    CONSTRAINT `tagvalue_ibfk_2` FOREIGN KEY (`tag_id`) REFERENCES `tag` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- +goose Down
DROP TABLE `tagvalue`;
