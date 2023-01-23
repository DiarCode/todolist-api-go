CREATE TABLE IF NOT EXISTS todos (
    id bigserial PRIMARY KEY,
    created_at timestamp(0) with time zone not null default NOW(), 
    title text not null,
    description text,
    completed boolean not null,
    user_id REFERENCES users(id)
);