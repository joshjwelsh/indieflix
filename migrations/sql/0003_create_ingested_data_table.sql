-- +goose Up
CREATE TABLE ingested_data(
        source VARCHAR(50) NOT NULL,
        identifier VARCHAR(50) NOT NULL,
        data JSONB NOT NULL,
        created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP
);

-- +goose Down
DROP TABLE ingested_data;
