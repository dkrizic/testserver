-- +goose Up
insert into asset_tag (asset_id, tag_id) values ("0191f003-681b-7042-b740-fb9a914ecd42", "0191ebaf-2238-7b99-b5e0-2c615eab5897"); -- Galaxy S24 +smartphone

-- +goose Down
delete from asset_tag;
