-- +goose Up
-- +goose StatementBegin
CREATE TABLE blocks (
    hash bytea PRIMARY KEY CHECK (LENGTH(hash) = 32),
    "size" bigint NOT NULL,
    stripped_size bigint NOT NULL,
    weight integer NOT NULL,
    "height" integer NOT NULL CHECK (HEIGHT >= -2),
    "version" integer NOT NULL,
    hash_merkle_root bytea CHECK (LENGTH(hash_merkle_root) = 32) NOT NULL,
    "time" integer NOT NULL,
    median_time integer NOT NULL,
    nonce bigint NOT NULL,
    bits bytea CHECK (LENGTH(bits) = 4) NOT NULL,
    difficulty double precision NOT NULL,
    chainwork bytea CHECK (LENGTH(chainwork) = 32) NOT NULL,
    orphan boolean NOT NULL DEFAULT FALSE
);

CREATE UNIQUE INDEX unique_height_index ON blocks (height) WHERE height >= 0;
CREATE INDEX block_height_index ON blocks (height); 
CREATE EXTENSION pg_trgm;


-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX unique_height_index;
DROP INDEX block_height_index;
DROP TABLE blocks;
DROP EXTENSION pg_trgm;
-- +goose StatementEnd
