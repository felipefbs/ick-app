CREATE TABLE "users" (
  "id" uuid PRIMARY KEY,
  "username" varchar,
  "name" varchar,
  "gender" varchar,
  "birthdate" date,
  "password" varchar,
  "created_at" timestamp DEFAULT (CURRENT_TIMESTAMP)
);

CREATE TABLE "icks" (
  "id" uuid PRIMARY KEY,
  "ick" text,
  "registered_by" int default 0 REFERENCES "users" ("id"),
  "created_at" timestamp DEFAULT (CURRENT_TIMESTAMP)
);

CREATE TABLE "user_icks" (
  "user_id" int REFERENCES "users" ("id"),
  "icks_id" int REFERENCES "icks" ("id"),
  PRIMARY KEY ("user_id", "icks_id")
);