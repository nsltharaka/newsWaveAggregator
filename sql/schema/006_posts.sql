-- +goose Up
CREATE TABLE posts (
    post_id UUID primary key,
    title text not null,
    description text,
    pub_date timestamp,
    url text unique,
    feed_id UUID not null references feeds (id) on delete CASCADE
);
-- +goose Down
DROP TABLE posts;