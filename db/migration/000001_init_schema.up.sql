CREATE TABLE "breed" (
                         "breed_id" BIGSERIAL UNIQUE PRIMARY KEY,
                         "breed_name" varchar UNIQUE NOT NULL,
                         "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "production" (
                              "id" BIGSERIAL PRIMARY KEY,
                              "breed_id" bigint UNIQUE NOT NULL,
                              "eggs" bigint NOT NULL,
                              "dirty" bigint NOT NULL,
                              "wrong_shape" bigint NOT NULL,
                              "weak_shell" bigint NOT NULL,
                              "damaged" bigint NOT NULL,
                              "hatching_eggs" bigint NOT NULL,
                              "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "hatchery" (
                            "id" BIGSERIAL PRIMARY KEY,
                            "production_id" bigint UNIQUE NOT NULL,
                            "infertile" bigint NOT NULL,
                            "early" bigint NOT NULL,
                            "middle" bigint NOT NULL,
                            "late" bigint NOT NULL,
                            "dead_chicks" bigint NOT NULL,
                            "alive_chicks" bigint NOT NULL,
                            "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "premises" (
                            "id" BIGSERIAL PRIMARY KEY,
                            "breed_id" bigint UNIQUE NOT NULL,
                            "farm" varchar NOT NULL,
                            "house" varchar NOT NULL,
                            "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX ON "breed" ("breed_name");

CREATE INDEX ON "production" ("breed_id");

CREATE INDEX ON "hatchery" ("production_id");

CREATE INDEX ON "premises" ("breed_id");

ALTER TABLE "production" ADD FOREIGN KEY ("breed_id") REFERENCES "breed" ("breed_id") ON DELETE CASCADE;

ALTER TABLE "hatchery" ADD FOREIGN KEY ("production_id") REFERENCES "production" ("id") ON DELETE CASCADE;

ALTER TABLE "premises" ADD FOREIGN KEY ("breed_id") REFERENCES "breed" ("breed_id") ON DELETE CASCADE;
