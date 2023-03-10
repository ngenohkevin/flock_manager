CREATE TABLE "breed" (
                         "breed_id" bigserial PRIMARY KEY,
                         "breed_name" varchar UNIQUE NOT NULL,
                         "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "production" (
                              "production_id" bigserial PRIMARY KEY,
                              "eggs" bigint NOT NULL,
                              "dirty" bigint NOT NULL,
                              "wrong_shape" bigint NOT NULL,
                              "weak_shell" bigint NOT NULL,
                              "damaged" bigint NOT NULL,
                              "hatching_eggs" bigint NOT NULL,
                              "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "hatchery" (
                            "hatchery_id" bigserial PRIMARY KEY,
                            "infertile" bigint NOT NULL,
                            "early" bigint NOT NULL,
                            "middle" bigint NOT NULL,
                            "late" bigint NOT NULL,
                            "dead_chicks" bigint NOT NULL,
                            "alive_chicks" bigint NOT NULL,
                            "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "premises" (
                            "premises_id" bigserial PRIMARY KEY,
                            "farm" varchar,
                            "house" varchar,
                            "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX ON "production" ("production_id");

CREATE INDEX ON "hatchery" ("hatchery_id");

ALTER TABLE "production" ADD FOREIGN KEY ("production_id") REFERENCES "breed" ("breed_id");

ALTER TABLE "hatchery" ADD FOREIGN KEY ("hatchery_id") REFERENCES "production" ("production_id");

ALTER TABLE "premises" ADD FOREIGN KEY ("premises_id") REFERENCES "breed" ("breed_id");
