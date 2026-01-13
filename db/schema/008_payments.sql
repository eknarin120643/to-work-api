CREATE TABLE "public"."payments" (
    "payment_id" BIGSERIAL PRIMARY KEY,
    "from_user_id" BIGINT,
    "to_user_id" BIGINT,
    "expense_id" BIGINT,
    "installment_payment_id" BIGINT,
    "amount" NUMERIC(12,2) NOT NULL,
    "state" INT DEFAULT 0,
    "created_at" TIMESTAMP DEFAULT now(),
    CHECK ("state" IN (0, 1)),
    CHECK (
        "expense_id" IS NOT NULL
        OR "installment_payment_id" IS NOT NULL
    ),
    CONSTRAINT "fk_payments_from_user" FOREIGN KEY ("from_user_id") REFERENCES "public"."users"("user_id"),
    CONSTRAINT "fk_payments_to_user" FOREIGN KEY ("to_user_id") REFERENCES "public"."users"("user_id"),
    CONSTRAINT "fk_payments_expense" FOREIGN KEY ("expense_id") REFERENCES "public"."expenses"("expense_id"),
    CONSTRAINT "fk_payments_installment" FOREIGN KEY ("installment_payment_id") REFERENCES "public"."installment_payments"("payment_id")
);
