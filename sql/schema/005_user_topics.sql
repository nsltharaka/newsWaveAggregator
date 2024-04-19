-- +goose Up
create table user_topics (
    id UUID primary key,
    created_at timestamp not null,
    updated_at timestamp not null,
    user_id integer references users(id) on delete cascade not null,
    topic_id UUID references topics(id) on delete cascade not null,
    unique(user_id, topic_id)
);
-- +goose Down
drop table user_topics;