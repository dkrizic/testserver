-- +goose Up
insert into user (id, email) values ("0191db3f-afbc-7895-b622-ef18506a8434", 'alice@bla.com');
insert into user (id, email) values ("0191db40-b756-7785-89fc-7d384b721c2e", 'bob@foo.com');

insert into `group` (id, name) values ("0191db41-347d-7883-8586-6592ddaea0a2", 'Display');
insert into `group` (id, name) values ("0191db41-b496-7b68-a2b7-758f6c84f162", 'Storage');

insert into user_group (id,user_id, group_id) values ("0191db4c-8486-78b0-9e75-c5137dd8de3d", "0191db3f-afbc-7895-b622-ef18506a8434", "0191db41-347d-7883-8586-6592ddaea0a2");
insert into user_group (id,user_id, group_id) values ("0191db4c-c565-721e-9038-253b21efb662", "0191db40-b756-7785-89fc-7d384b721c2e", "0191db41-b496-7b68-a2b7-758f6c84f162");
insert into user_group (id,user_id, group_id) values ("0191db4c-fdd5-79cf-9ceb-ebdea5f3d8c8", "0191db3f-afbc-7895-b622-ef18506a8434", "0191db41-347d-7883-8586-6592ddaea0a2");

-- +goose Down
delete from user_group where user_id = "0191db3f-afbc-7895-b622-ef18506a8434" and group_id = "0191db41-b496-7b68-a2b7-758f6c84f162";
delete from user_group where user_id = "0191db40-b756-7785-89fc-7d384b721c2e" and group_id = "0191db41-b496-7b68-a2b7-758f6c84f162";
delete from user_group where user_id = "0191db3f-afbc-7895-b622-ef18506a8434" and group_id = "0191db41-347d-7883-8586-6592ddaea0a2";

delete from `group` where id = "0191db41-b496-7b68-a2b7-758f6c84f162";
delete from `group` where id = "0191db41-347d-7883-8586-6592ddaea0a2";

delete from user where id = "0191db40-b756-7785-89fc-7d384b721c2e";
delete from user where id = "0191db3f-afbc-7895-b622-ef18506a8434";
