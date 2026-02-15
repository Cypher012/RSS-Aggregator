-- +goose up
CREATE EXTENSION IF NOT EXISTS "pgcrypto";

Create Table feeds (
	id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	url varchar(255) NOT NULL,
	name varchar(255) NOT NULL,
	user_id UUID REFERENCES users(id) ON DELETE CASCADE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- +goose down
DROP EXTENSION IF EXISTS "pgcrypto";
DROP TABLE feeds;
