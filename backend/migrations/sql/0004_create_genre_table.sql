-- +goose Up
CREATE TABLE genres(
	id SERIAL PRIMARY KEY, 
	name VARCHAR(50) NOT NULL,
	created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP
);
-- +goose Down
DROP TABLE genres;
