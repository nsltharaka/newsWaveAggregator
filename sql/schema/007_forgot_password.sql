-- +goose Up
CREATE TABLE forgot_password (
    case_number UUID PRIMARY KEY,
    opened BOOLEAN NOT NULL,
    user_id INTEGER REFERENCES users (id) NOT NULL
);

-- +goose Down

DROP TABLE forgot_password;