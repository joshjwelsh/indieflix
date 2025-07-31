-- +goose Up
-- +goose StatementBegin
CREATE TABLE entity_list_entries(
        id SERIAL PRIMARY KEY,
        entity_list_id INTEGER REFERENCES entity_lists(id), 
        movie_id INTEGER REFERENCES movies(id),
        description VARCHAR(100),
        created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE entity_list_entries;
-- +goose StatementEnd
