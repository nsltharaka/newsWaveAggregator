-- +goose Up
create table users (
    id serial primary key,
    created_at timestamp not null,
    updated_at timestamp not null,
    username text not null,
    email text not null,
    password text not null,
    api_key  varchar(64) unique not null default (encode(sha256(random()::text::bytea), 'hex'))
);

-- +goose Down
drop table users;