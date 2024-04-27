-- +goose Up
create table topics (
    id UUID primary key,
    name text unique not null,
    img_url text, 
    updated_at timestamp not null,
    created_by integer references users(id) not null
);  
-- +goose Down
drop table topics;