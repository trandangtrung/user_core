CREATE TABLE "Platform" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "config" varchar NOT NULL,
  "created_by" integer,
  "updated_by" integer,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);
