-- +goose Up
create table topic_contains_feed (
    topic_id UUID references topics(id) not null,
    feed_id UUID references feeds(id) not null,
    user_id integer references users(id) not null,
    primary key(topic_id, feed_id, user_id)
);
-- +goose Down
drop table topic_contains_feed;