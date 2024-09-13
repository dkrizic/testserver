-- +goose Up
insert into tagcategory (id,name) values("0191eb8b-c756-7c39-8a0a-508b7cc13620", 'Business Unit');
insert into tag (id,discriminator,name,tagcategory_id) values("0191eb8e-1628-7ec5-b47e-bcc3daf82987", 'static', 'Mobile', "0191eb8b-c756-7c39-8a0a-508b7cc13620");

insert into tagcategory (id,name,parent) values("0191eb8b-7c39-8a0a-c756-508b7cc13620", "Produktkategorie", "0191eb8b-c756-7c39-8a0a-508b7cc13620");
insert into tag (id,discriminator,name,tagcategory_id) values("0191ebaf-2238-7b99-b5e0-2c615eab5897", 'static', 'Smartphone', "0191eb8b-7c39-8a0a-c756-508b7cc13620");
insert into tag (id,discriminator,name,tagcategory_id) values("0191ebaf-962d-79bd-934f-906bb8a44030", 'static', 'Tablet', "0191eb8b-7c39-8a0a-c756-508b7cc13620");



-- +goose Down
delete from tag;
delete from tagcategory;
