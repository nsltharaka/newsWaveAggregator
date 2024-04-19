-- +goose Up
create table topics (
    id UUID primary key,
    name text not null
);
-- +goose Down
drop table topics;