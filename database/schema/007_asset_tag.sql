-- +goose Up
create table asset_tag (
    asset_id char(36) not null,
    tag_id char(36) not null,
    primary key (asset_id, tag_id),
    foreign key (asset_id) references asset (id),
    foreign key (tag_id) references tag (id)
);

-- +goose Down
drop table asset_tag;

