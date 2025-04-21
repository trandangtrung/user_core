CREATE TABLE "Tokens" (
  "id" bigserial PRIMARY KEY,
  "user_id" integer UNIQUE NOT NULL,
  "refresh_token" varchar NOT NULL,
  "scope" varchar NOT NULL,
  "created_by" integer,
  "updated_by" integer,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
  FOREIGN KEY ("user_id") REFERENCES "Users"("id") ON DELETE CASCADE
);
