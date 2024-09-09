-- +goose Up
CREATE TABLE user (
    id int not null PRIMARY KEY auto_increment,
    email VARCHAR(255) NOT NULL
);
