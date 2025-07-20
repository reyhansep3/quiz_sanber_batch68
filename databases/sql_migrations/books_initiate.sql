-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE books (
    id SERIAL PRIMARY KEY,
    title VARCHAR(60) NOT NULL,
    description VARCHAR(256),
    image_url VARCHAR(256),
    release_year INTEGER,
    price INTEGER,
    total_page INTEGER,
    thickness VARCHAR(50),
    category_id INTEGER,
    created_at TIMESTAMP,
    created_by VARCHAR(100),
    modified_at TIMESTAMP,
    modified_by VARCHAR(100)
);

-- +migrate StatementEnd