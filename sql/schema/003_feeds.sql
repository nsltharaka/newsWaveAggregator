-- +goose Up
create table feeds (
    id UUID primary key,
    created_at timestamp not null,
    updated_at timestamp,
    url text unique not null
);
-- +goose Down
drop table feeds;