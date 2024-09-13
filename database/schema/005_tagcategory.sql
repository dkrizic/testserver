-- +goose Up
create table tagcategory (
    id char(36) not null primary key,
    name varchar(255) not null,
    parent char(36),
    discriminator enum ('static','dynamic') not null,
    foreign key (parent) references tagcategory (id),
    description text,
    format varchar(255), -- format for dynamic tags
    open boolean not null default false-- for static tags
);

-- +goose Down
drop table tagcategory;
