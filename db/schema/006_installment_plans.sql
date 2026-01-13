CREATE TABLE "public"."installment_plans" (
    "plan_id" BIGSERIAL PRIMARY KEY,
    "expense_id" BIGINT UNIQUE,
    "debtor_id" BIGINT,
    "total_months" INT NOT NULL,
    "monthly_amount" NUMERIC(12,2) NOT NULL,
    "start_date" DATE NOT NULL,
    "state" INT DEFAULT 0,
    CHECK ("state" IN (0, 1)),
    CONSTRAINT "fk_installment_plans_expense" FOREIGN KEY ("expense_id") REFERENCES "public"."expenses"("expense_id") ON DELETE CASCADE,
    CONSTRAINT "fk_installment_plans_debtor" FOREIGN KEY ("debtor_id") REFERENCES "public"."users"("user_id")
);

CREATE INDEX "idx_installment_plan_debtor" ON "public"."installment_plans"("debtor_id");
