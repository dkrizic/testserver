-- +goose Up
CREATE TABLE xgroup (
    id int not null PRIMARY KEY auto_increment,
    name VARCHAR(255) NOT NULL
);