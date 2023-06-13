-- +goose Up
CREATE TABLE movies(
        id SERIAL PRIMARY KEY, 
        name VARCHAR(50) NOT NULL, 
        source_id INTEGER REFERENCES sources(id), 
        metadata JSONB NOT NULL, 
        showtimes JSONB NOT NULL, 
        available BOOL NOT NULL,
        created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP
);

-- +goose Down
DROP TABLE movies;
