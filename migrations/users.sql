-- +migrate Up
CREATE TABLE "users" (
                         "username" varchar(255)  NOT NULL,
                         "password" varchar(255)  NOT NULL,
                         "email" varchar(255)  NOT NULL
);
INSERT INTO "users" ("username","password","email")
VALUES ('admin','admin1234','admin@gmail.com');
-- +migrate Down
DROP TABLE users;