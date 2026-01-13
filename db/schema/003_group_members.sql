CREATE TABLE "public"."group_members" (
    "id" BIGSERIAL PRIMARY KEY,
    "group_id" BIGINT,
    "user_id" BIGINT,
    "joined_at" TIMESTAMP DEFAULT now(),
    UNIQUE ("group_id", "user_id"),
    CONSTRAINT "fk_group_members_group" FOREIGN KEY ("group_id") REFERENCES "public"."groups"("group_id") ON DELETE CASCADE,
    CONSTRAINT "fk_group_members_user" FOREIGN KEY ("user_id") REFERENCES "public"."users"("user_id") ON DELETE CASCADE
);
