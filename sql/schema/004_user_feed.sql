-- +goose Up
create table user_feeds (
    id UUID primary key,
    created_at timestamp not null,
    updated_at timestamp not null,
    user_id integer references users(id) on delete cascade not null,
    feed_id UUID references feeds(id) on delete cascade not null,
    unique(user_id, feed_id)
);
-- +goose Down
drop table user_feeds;