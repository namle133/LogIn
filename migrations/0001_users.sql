-- +migrate Up
CREATE TABLE "public"."users" (
                         "username" varchar(255)  PRIMARY KEY,
                         "password" bytea  NOT NULL,
                         "email" varchar(255)  NOT NULL
);
-- +migrate Down
DROP TABLE users;
CREATE SCHEMA public;