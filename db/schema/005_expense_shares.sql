CREATE TABLE "public"."expense_shares" (
    "share_id" BIGSERIAL PRIMARY KEY,
    "expense_id" BIGINT,
    "user_id" BIGINT,
    "amount" NUMERIC(12,2) NOT NULL,
    CONSTRAINT "fk_expense_shares_expense" FOREIGN KEY ("expense_id") REFERENCES "public"."expenses"("expense_id") ON DELETE CASCADE,
    CONSTRAINT "fk_expense_shares_user" FOREIGN KEY ("user_id") REFERENCES "public"."users"("user_id")
);

CREATE INDEX "idx_expense_shares_expense" ON "public"."expense_shares"("expense_id");
