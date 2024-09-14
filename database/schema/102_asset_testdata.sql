-- +goose Up
insert into asset (id, name) values ("0191f003-681b-7042-b740-fb9a914ecd42", 'Galaxy S24');

insert into file(id, name, asset_id, size, mimetype) values ("0191f029-80c8-75b2-9725-ff28c0abdaac", "galaxy-s24.jpg", "0191f003-681b-7042-b740-fb9a914ecd42", 1024, "image/jpeg");
insert into file(id, name, asset_id, size, mimetype) values ("0191f029-daa5-7b30-8a74-9ad5ca282172", "galaxy-s24.png", "0191f003-681b-7042-b740-fb9a914ecd42", 2048, "image/png");

-- +goose Down
delete from file;
delete from asset;
