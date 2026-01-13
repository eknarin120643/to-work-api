CREATE TABLE "public"."installment_payments" (
    "payment_id" BIGSERIAL PRIMARY KEY,
    "plan_id" BIGINT,
    "month_no" INT NOT NULL,
    "due_date" DATE NOT NULL,
    "amount" NUMERIC(12,2) NOT NULL,
    "state" INT DEFAULT 0,
    "paid_date" TIMESTAMP,
    UNIQUE ("plan_id", "month_no"),
    CHECK ("state" IN (0, 1)),
    CONSTRAINT "fk_installment_payments_plan" FOREIGN KEY ("plan_id") REFERENCES "public"."installment_plans"("plan_id") ON DELETE CASCADE
);

CREATE INDEX "idx_installment_payment_state" ON "public"."installment_payments"("state");
