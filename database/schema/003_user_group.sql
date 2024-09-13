-- +goose Up
CREATE TABLE user_group (
    id char(36) not null PRIMARY KEY,
    user_id char(36) not NULL,
    group_id char(36) NOT NULL,
    key (user_id, group_id),

    CONSTRAINT user_group_user_id_fk FOREIGN KEY (user_id) REFERENCES user (id),
    CONSTRAINT user_group_group_id_fk FOREIGN KEY (group_id) REFERENCES `group` (id)
);

-- +goose Down
DROP TABLE user_group;
