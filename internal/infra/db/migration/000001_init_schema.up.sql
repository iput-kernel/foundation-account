CREATE TYPE "role" AS ENUM (
  'student',
  'teacher',
  'left'
);

CREATE TABLE "follows" (
  "following_user_id" uuid NOT NULL,
  "followed_user_id" uuid NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "users" (
  "id" uuid PRIMARY KEY,
  "name" varchar UNIQUE NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "password_hash" varchar NOT NULL,
  "role" role NOT NULL,
  "credit" bigint NOT NULL,
  "level" int NOT NULL,
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "statements" (
  "id" uuid PRIMARY KEY,
  "user_id" uuid NOT NULL,
  "amount" bigint NOT NULL,
  "reason" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "transfers" (
  "id" uuid PRIMARY KEY,
  "from_user_id" uuid NOT NULL,
  "to_user_id" uuid NOT NULL,
  "amount" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "verify_emails" (
  "id" uuid PRIMARY KEY,
  "name" varchar NOT NULL,
  "email" varchar NOT NULL,
  "password_hash" varchar NOT NULL,
  "secret_code" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "expired_at" timestamptz NOT NULL DEFAULT (now() + interval '15 minutes')
);

CREATE TABLE "sessions" (
  "id" uuid PRIMARY KEY,
  "user_id" uuid NOT NULL,
  "refresh_token" varchar NOT NULL,
  "user_agent" varchar NOT NULL,
  "client_ip" varchar NOT NULL,
  "is_blocked" boolean NOT NULL DEFAULT false,
  "expires_at" timestamptz NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX ON "users" ("role");

CREATE INDEX ON "statements" ("user_id");

CREATE INDEX ON "transfers" ("from_user_id");

CREATE INDEX ON "transfers" ("to_user_id");

CREATE INDEX ON "transfers" ("from_user_id", "to_user_id");

COMMENT ON COLUMN "statements"."amount" IS '値の正負両方を許容する';

COMMENT ON COLUMN "statements"."reason" IS '必ず増減理由を明記する必要がある';

COMMENT ON COLUMN "transfers"."amount" IS '値は正でなければならない';

ALTER TABLE "follows" ADD FOREIGN KEY ("following_user_id") REFERENCES "users" ("id");

ALTER TABLE "follows" ADD FOREIGN KEY ("followed_user_id") REFERENCES "users" ("id");

ALTER TABLE "statements" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "transfers" ADD FOREIGN KEY ("from_user_id") REFERENCES "users" ("id");

ALTER TABLE "transfers" ADD FOREIGN KEY ("to_user_id") REFERENCES "users" ("id");

ALTER TABLE "sessions" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");
