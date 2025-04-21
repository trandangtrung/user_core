CREATE TABLE "UserPlatform" (
  "id" bigserial PRIMARY KEY,
  "user_id" integer UNIQUE NOT NULL,
  "platform_id" integer UNIQUE NOT NULL,
  "created_by" integer,
  "updated_by" integer,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
  FOREIGN KEY ("user_id") REFERENCES "Users"("id") ON DELETE CASCADE,
  FOREIGN KEY ("platform_id") REFERENCES "Platform"("id") ON DELETE CASCADE
);


ALTER TABLE "Role" ADD CONSTRAINT "fk_UserRole" FOREIGN KEY ("user_platform_id") REFERENCES "UserPlatform"("id") ON DELETE CASCADE;