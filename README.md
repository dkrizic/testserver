## Create database

```
create table asset (
    id int not null primary key auto_increment, 
    name varchar(200) not null);
create table tag (
    id int not null primary key auto_increment, 
    name varchar(200) not null);
create table tagvalue (
    id int not null primary key auto_increment, 
    asset_id int not null, 
    tag_id int not null, 
    value varchar(200),
    foreign key (asset_id) references asset(id),
    foreign key (tag_id) references tag(id));
```
