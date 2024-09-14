-- +goose Up
insert into asset_tag (asset_id, tag_id) values ("0191f003-681b-7042-b740-fb9a914ecd42", "0191ebaf-2238-7b99-b5e0-2c615eab5897"); -- Galaxy S24 +smartphone
insert into asset_tag (asset_id, tag_id) values ("0191f003-681b-7042-b740-fb9a914ecd42", "0191f041-a7a2-7451-85db-01e23e5c10f2");

-- +goose Down
delete from asset_tag;
