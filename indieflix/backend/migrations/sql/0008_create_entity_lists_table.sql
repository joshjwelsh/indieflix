-- +goose Up
-- +goose StatementBegin
CREATE TABLE entity_lists(
        id SERIAL PRIMARY KEY, 
        entity_type VARCHAR(50) NOT NULL, 
        user_id INTEGER REFERENCES users(id),
        created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE entity_lists;
-- +goose StatementEnd
