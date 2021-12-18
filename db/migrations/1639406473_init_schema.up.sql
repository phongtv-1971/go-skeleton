CREATE TABLE users (
   id bigserial PRIMARY KEY,
   name varchar NOT NULL,
   email varchar NOT NULL,
   created_at timestamp NOT NULL DEFAULT (now())
);
