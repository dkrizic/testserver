-- +goose Up
create table access (
    id char(36) not null,
    asset_id char(36) not null,
    identity_id char(36) not null,
    permission char(30) not null,
    discriminator char(30) not null,
    primary key (id),
    unique key (asset_id, identity_id),
    CONSTRAINT access_asset_id_fk FOREIGN KEY (asset_id) REFERENCES asset (id)
);

-- +goose Down
drop table access;
