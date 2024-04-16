-- +goose Up
create table users (
    id serial primary key,
    created_at timestamp not null,
    updated_at timestamp not null,
    username text not null,
    email text not null,
    password text not null
);

-- +goose Down
drop table users;