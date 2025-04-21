CREATE TABLE "UserRole" (
  "id" bigserial PRIMARY KEY,
  "user_id" integer UNIQUE NOT NULL,
  "role_id" integer UNIQUE NOT NULL,
  "created_by" integer,
  "updated_by" integer,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
  FOREIGN KEY ("user_id") REFERENCES "Users"("id") ON DELETE CASCADE,
  FOREIGN KEY ("role_id") REFERENCES "Role"("id") ON DELETE CASCADE
);

