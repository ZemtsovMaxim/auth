CREATE TABLE IF NOT EXISTS "owner" (
  "id" SERIAL PRIMARY KEY,
  "full_name" varchar UNIQUE NOT NULL,
  "citizenship" varchar NOT NULL
);

CREATE TABLE IF NOT EXISTS "account" (
  "id" SERIAL PRIMARY KEY,
  "balance" int NOT NULL DEFAULT 0,
  "owner_id" int NOT NULL,
  "is_locked" bool NOT NULL DEFAULT false
);

CREATE TABLE IF NOT EXISTS "transaction" (
  "id" SERIAL PRIMARY KEY,
  "account_id" int NOT NULL,
  "participating_account_id" int,
  "transaction_type" varchar NOT NULL,
  "amount" int NOT NULL,
  "date" timestamp NOT NULL DEFAULT 'now()'
);

ALTER TABLE "account" ADD FOREIGN KEY ("owner_id") REFERENCES "owner" ("id") ON DELETE CASCADE;

ALTER TABLE "transaction" ADD FOREIGN KEY ("account_id") REFERENCES "account" ("id") ON DELETE CASCADE;

ALTER TABLE "transaction" ADD FOREIGN KEY ("participating_account_id") REFERENCES "account" ("id") ON DELETE CASCADE;
