-- +goose Up
insert into asset (id, name) values ("0191f003-681b-7042-b740-fb9a914ecd42", 'Galaxy S24');

-- +goose Down
delete from asset where id = "0191f003-681b-7042-b740-fb9a914ecd42";