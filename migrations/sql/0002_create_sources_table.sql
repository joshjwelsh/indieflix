-- +goose Up
CREATE TABLE sources(
        id SERIAL PRIMARY KEY,
        name VARCHAR(50) NOT NULL,
        website VARCHAR(50) NOT NULL,
        created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP 
);

-- +goose Down
DROP TABLE sources;
