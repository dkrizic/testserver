-- +goose Up

-- alice -> write
insert into access(id,asset_id, identity_id, permission, discriminator)
values ("0191f23d-72ef-7d13-90d2-65bd36e4bf42",
        "0191f003-681b-7042-b740-fb9a914ecd42",
        "0191db3f-afbc-7895-b622-ef18506a8434",
        "write",
        'user');

-- Display -> read
insert into access(id,asset_id, identity_id, permission, discriminator)
values ("0191f23d-97a6-7afc-8839-4fd187350187",
        "0191f003-681b-7042-b740-fb9a914ecd42",
        "0191db41-347d-7883-8586-6592ddaea0a2",
        "read",
        'group');

-- +goose Down
delete from access;
