-- +goose Up
insert into user (id, email) values ("0191db3f-afbc-7895-b622-ef18506a8434", 'alice@bla.com');
insert into user (id, email) values ("0191db40-b756-7785-89fc-7d384b721c2e", 'bob@foo.com');

insert into `group` (id, name) values ("0191db41-347d-7883-8586-6592ddaea0a2", 'Display');
insert into `group` (id, name) values ("0191db41-b496-7b68-a2b7-758f6c84f162", 'Storage');

-- alice -> Display
insert into user_group (user_id, group_id) values ("0191db3f-afbc-7895-b622-ef18506a8434", "0191db41-347d-7883-8586-6592ddaea0a2");

-- bob -> Storage
insert into user_group (user_id, group_id) values ("0191db40-b756-7785-89fc-7d384b721c2e", "0191db41-b496-7b68-a2b7-758f6c84f162");

-- alice -> Storage
insert into user_group (user_id, group_id) values ("0191db3f-afbc-7895-b622-ef18506a8434", "0191db41-b496-7b68-a2b7-758f6c84f162");

-- +goose Down
delete from user_group;
delete from `group`;
delete from user;
