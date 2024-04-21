-- +goose Up
create table feeds (
    id UUID primary key,
    created_at timestamp not null,
    updated_at timestamp not null,
    url text not null,
    topic_id uuid references topics(id) on delete cascade not null,
    user_id integer references users(id) on delete cascade not null
);
-- +goose Down
drop table feeds;