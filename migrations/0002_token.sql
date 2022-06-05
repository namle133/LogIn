-- +migrate Up
CREATE TABLE "public"."tokens" (
                         "username" varchar(255)  ,
                         "token_string" varchar(500) NOT NULL,
                         PRIMARY KEY ("username","token_string"),
                         FOREIGN KEY ("username") REFERENCES "public"."users" ("username")
);
-- +migrate Down
DROP TABLE "public"."tokens";