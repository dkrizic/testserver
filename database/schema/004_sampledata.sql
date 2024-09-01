-- +goose Up
insert into tag (id,name) values (1, "filename");
insert into tag (id,name) values (2, "country");
insert into tag (id,name) values (3, "product_number");
insert into tag (id,name) values (4, "country");

insert into asset (id,name) values (1, "Galaxy S24");
insert into asset (id,name) values (2, "EarBuds");
insert into asset (id,name) values (3, "TV 7090");
insert into asset (id,name) values (4, "Watch Ultra");

-- +goose Down
delete from tag;
