-- +goose Up
-- +goose StatementBegin
create table orders
(
    order_uid varchar
        constraint orders_pk
            primary key,
    data      jsonb
);
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
drop table orders;
-- +goose StatementEnd
