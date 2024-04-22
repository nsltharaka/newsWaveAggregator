-- +goose Up
create table user_follows_topic (
    user_id integer references users(id) not null,
    topic_id UUID references topics(id) not null,
    primary key(user_id, topic_id)
);
-- +goose Down
drop table user_follows_topic;