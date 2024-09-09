-- +goose Up
CREATE TABLE user (
    id int not null PRIMARY KEY auto_increment,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL
);
