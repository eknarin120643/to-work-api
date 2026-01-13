CREATE TABLE "public"."users" (
    "user_id" BIGSERIAL PRIMARY KEY,
    "name" VARCHAR(100) NOT NULL,
    "email" VARCHAR(150) UNIQUE NOT NULL,
    "password_hash" TEXT NOT NULL,
    "promptpay_id" VARCHAR(50),
    "promptpay_qr" TEXT,
    "created_at" TIMESTAMP DEFAULT now()
);
