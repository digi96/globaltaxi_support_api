CREATE TABLE "payload" (
  "payload_id" SERIAL PRIMARY KEY,
  "body" varchar NOT NULL,
  "status" varchar NOT NULL DEFAULT (0),
  "created_at" timestamp NOT NULL DEFAULT (now())
);