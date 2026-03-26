-- +goose Up
-- +goose StatementBegin
CREATE TABLE posts (
    id UUID NOT NULL PRIMARY KEY,
    url TEXT NOT NULL,
    title TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    description TEXT NOT NULL,
    published_at TIMESTAMP NOT NULL, 
    feed_id UUID NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE posts;
-- +goose StatementEnd
