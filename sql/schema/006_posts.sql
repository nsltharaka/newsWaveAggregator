-- +goose Up
CREATE TABLE posts (
    post_id UUID primary key,
    title text not null,
    description text,
    author text,
    pub_date timestamp not null,
    post_image text,
    url text unique not null,
    feed_id UUID not null references feeds (id) on delete CASCADE
);
-- +goose Down
DROP TABLE posts;