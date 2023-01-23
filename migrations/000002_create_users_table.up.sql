CREATE TABLE IF NOT EXISTS users (
    id bigserial PRIMARY KEY,
    created_at timestamp(0) with time zone not null default NOW(), 
    email text not null,
    password text not null,
);