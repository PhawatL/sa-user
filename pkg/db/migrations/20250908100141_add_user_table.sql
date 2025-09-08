-- +goose Up
-- +goose StatementBegin
CREATE TYPE "roles" AS ENUM (
  'patient',
  'doctor',
  'admin'
);

CREATE TABLE IF NOT EXISTS "users" (
    "id" uuid PRIMARY KEY,
    "hospital_id" text NOT NULL UNIQUE,
    "email" text NOT NULL UNIQUE,
    "password" text NOT NULL,
    "first_name" text NOT NULL,
    "last_name" text NOT NULL,
    "gender" text NOT NULL,
    "phone_number" text NOT NULL,
    "created_at" timestamp with time zone,
    "updated_at" timestamp with time zone,
    "deleted_at" timestamp with time zone
);

CREATE TABLE IF NOT EXISTS "patients" (
    "user_id" uuid PRIMARY KEY,
    "address" text,
    "allergies" text,
    "emergency_contact" text,
    "blood_type" varchar(5),
    "created_at" timestamp with time zone,
    "updated_at" timestamp with time zone,
    "deleted_at" timestamp with time zone
);

CREATE TABLE IF NOT EXISTS "doctors" (
    "user_id" uuid PRIMARY KEY,
    "specialty" text,
    "bio" text,
    "years_experience" int,
    "created_at" timestamp with time zone,
    "updated_at" timestamp with time zone,
    "deleted_at" timestamp with time zone
);

CREATE TABLE IF NOT EXISTS "user_roles" (
  "user_id" uuid,
  "role" roles,
  PRIMARY KEY ("user_id", "role")
);

ALTER TABLE "user_roles" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "user_roles";
DROP TABLE IF EXISTS "doctors";
DROP TABLE IF EXISTS "patients";
DROP TABLE IF EXISTS "users";
DROP TYPE IF EXISTS "roles";
-- +goose StatementEnd
