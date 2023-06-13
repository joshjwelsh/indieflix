-- +goose Up
-- +goose StatementBegin
CREATE TABLE scores(
        user_id INTEGER REFERENCES users(id), 
        movie_id INTEGER REFERENCES movies(id),
        rating INTEGER DEFAULT 50,
	created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE scores;
-- +goose StatementEnd
