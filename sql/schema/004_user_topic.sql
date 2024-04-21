-- +goose Up
create table user_topic (
    id UUID primary key,
    user_id integer references users(id) on delete cascade not null,
    topic_id UUID references topics(id) on delete cascade not null,
    unique(user_id, topic_id)
);

-- +goose Down
drop table user_topic;