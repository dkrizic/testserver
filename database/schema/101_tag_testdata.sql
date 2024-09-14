-- +goose Up
insert into tagcategory (id,name) values("0191eb8b-c756-7c39-8a0a-508b7cc13620", 'Business Unit');
insert into tag (id,discriminator,name,tagcategory_id) values("0191eb8e-1628-7ec5-b47e-bcc3daf82987", 'static', 'Mobile', "0191eb8b-c756-7c39-8a0a-508b7cc13620");

insert into tagcategory (id,name,parent) values("0191eb8b-7c39-8a0a-c756-508b7cc13620", "Produktkategorie", "0191eb8b-c756-7c39-8a0a-508b7cc13620");
insert into tag (id,discriminator,name,tagcategory_id,parent_tag_id) values("0191ebaf-2238-7b99-b5e0-2c615eab5897", 'static', 'Smartphone', "0191eb8b-7c39-8a0a-c756-508b7cc13620","0191eb8e-1628-7ec5-b47e-bcc3daf82987");
insert into tag (id,discriminator,name,tagcategory_id,parent_tag_id) values("0191ebaf-962d-79bd-934f-906bb8a44030", 'static', 'Tablet',     "0191eb8b-7c39-8a0a-c756-508b7cc13620", "0191eb8e-1628-7ec5-b47e-bcc3daf82987");

insert into tagcategory (id,name,parent,discriminator,format) values("0191ebe7-4f8d-78a2-9f87-4695c1a15f0d", "Freigabedatum", "0191eb8b-7c39-8a0a-c756-508b7cc13620","dynamic", "[0-9]{4}-[0-9]{2}-[0-9]{2}");
insert into tag (id,discriminator,value,tagcategory_id) values("0191f041-a7a2-7451-85db-01e23e5c10f2", 'dynamic', "2024-01-01","0191ebe7-4f8d-78a2-9f87-4695c1a15f0d");

-- +goose Down
