-- +goose Up
create table tag (
    id char(36) not null primary key,
    name varchar(255), -- for static tags
    discriminator enum ('static','dynamic') not null,
    tagcategory_id char(36) not null,
    parent_tag_id char(36),
    foreign key (tagcategory_id) references tagcategory (id),
    foreign key (parent_tag_id) references tag (id),
    value varchar(255) -- for dynamic tags
);

-- +goose Down
drop table tag;
