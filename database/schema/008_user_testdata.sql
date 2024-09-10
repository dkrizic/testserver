-- +goose Up
insert into user (id, email) values (1, 'alice@bla.com');
insert into user (id, email) values (2, 'bob@foo.com');

insert into xgroup (id, name) values (1, 'Display');
insert into xgroup (id, name) values (2, 'Storage');

insert into user_group (user_id, group_id) values (1, 1);
insert into user_group (user_id, group_id) values (2, 2);
insert into user_group (user_id, group_id) values (1, 2);
