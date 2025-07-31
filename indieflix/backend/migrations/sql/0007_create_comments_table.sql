-- +goose Up
-- +goose StatementBegin
CREATE TABLE comments(
        id SERIAL PRIMARY KEY, 
        user_id INTEGER REFERENCES users(id), 
        user_text VARCHAR(250) NOT NULL,				
        created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP

);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE comments;
-- +goose StatementEnd
