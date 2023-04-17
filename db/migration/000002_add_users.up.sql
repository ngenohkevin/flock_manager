CREATE TABLE "users" (
                         "username" varchar UNIQUE PRIMARY KEY NOT NULL,
                         "hashed_password" varchar NOT NULL,
                         "full_name" varchar NOT NULL,
                         "email" varchar UNIQUE NOT NULL,
                         "password_changed_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00Z',
                         "created_at" timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE "breed" ADD FOREIGN KEY ("breed_name") REFERENCES "users" ("username");