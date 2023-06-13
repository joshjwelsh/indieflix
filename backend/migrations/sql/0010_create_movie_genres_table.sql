-- +goose Up
-- +goose StatementBegin
CREATE TABLE movie_genres(
        genre_id INTEGER REFERENCES genres(id), 
        movie_id INTEGER REFERENCES movies(id),
        created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP
);  
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE movie_genres;
-- +goose StatementEnd
