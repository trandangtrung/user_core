CREATE TABLE "Role" (
  "id" bigserial PRIMARY KEY,
  "user_platform_id" integer UNIQUE NOT NULL,
  "name" varchar NOT NULL,
  "description" varchar NOT NULL,
  "created_by" integer,
  "updated_by" integer,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);


