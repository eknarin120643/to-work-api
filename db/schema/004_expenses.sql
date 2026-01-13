CREATE TABLE "public"."expenses" (
    "expense_id" BIGSERIAL PRIMARY KEY,
    "group_id" BIGINT,
    "title" VARCHAR(200) NOT NULL,
    "total_amount" NUMERIC(12,2) NOT NULL,
    "payer_id" BIGINT,
    "expense_type" INT NOT NULL,
    "created_at" TIMESTAMP DEFAULT now(),
    CHECK ("expense_type" IN (0, 1)),
    CONSTRAINT "fk_expenses_group" FOREIGN KEY ("group_id") REFERENCES "public"."groups"("group_id") ON DELETE CASCADE,
    CONSTRAINT "fk_expenses_payer" FOREIGN KEY ("payer_id") REFERENCES "public"."users"("user_id")
);

CREATE INDEX "idx_expenses_group" ON "public"."expenses"("group_id");
