-- +goose Up
CREATE TABLE user_group (
    id int not null PRIMARY KEY auto_increment,
    user_id INT NOT NULL,
    group_id INT NOT NULL
);