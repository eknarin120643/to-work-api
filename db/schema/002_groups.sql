CREATE TABLE "public"."groups" (
    "group_id" BIGSERIAL PRIMARY KEY,
    "group_name" VARCHAR(150) NOT NULL,
    "owner_id" BIGINT,
    "created_at" TIMESTAMP DEFAULT now(),
    CONSTRAINT "fk_groups_owner" FOREIGN KEY ("owner_id") REFERENCES "public"."users"("user_id")
);
