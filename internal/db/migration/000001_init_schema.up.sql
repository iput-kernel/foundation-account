CREATE EXTENSION "uuid-ossp";

CREATE TABLE "follow" (
  "following_user_id" uuid,
  "followed_user_id" uuid,
  "created_at" timestamptz DEFAULT (now())
);

CREATE TABLE "user" (
  "id" uuid PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "name" varchar(20) NOT NULL,
  "email" varchar(100) UNIQUE NOT NULL,
  "password_hash" varchar(60) NOT NULL,
  "role" string NOT NULL,
  "class_id" uuid,
  "is_verified" bool NOT NULL DEFAULT false,
  "updated_at" timestamptz DEFAULT (now()),
  "created_at" timestamptz DEFAULT (now())
);

CREATE TABLE "profile" (
  "id" uuid PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "user_id" uuid NOT NULL,
  "real_name" varchar(30),
  "birthday" date,
  "sex" varchar(6),
  "phone" varchar(14),
  "describe" varchar(200),
  "created_at" timestamptz DEFAULT (now())
);

CREATE TABLE "verify_email" (
  "id" uuid PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "user_id" uuid NOT NULL,
  "email" varchar(100) NOT NULL,
  "secret_code" varchar NOT NULL,
  "is_used" bool NOT NULL DEFAULT false,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "expired_at" timestamptz NOT NULL DEFAULT (now() + interval '15 minutes')
);

ALTER TABLE "follow" ADD FOREIGN KEY ("following_user_id") REFERENCES "user" ("id");

ALTER TABLE "follow" ADD FOREIGN KEY ("followed_user_id") REFERENCES "user" ("id");

ALTER TABLE "profile" ADD FOREIGN KEY ("user_id") REFERENCES "user" ("id");

ALTER TABLE "verify_email" ADD FOREIGN KEY ("id") REFERENCES "user" ("id");
