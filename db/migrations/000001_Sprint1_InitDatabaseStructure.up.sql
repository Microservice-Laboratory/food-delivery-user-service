CREATE TABLE "users" (
    "id" uuid,
    "first_name" varchar(100) NOT NULL,
    "last_name" varchar(100) NOT NULL,
    "email" varchar(100) NOT NULL,
    "password" text NOT NULL,
    "phone" text NOT NULL,
    "role" varchar(30) NOT NULL,
    "created_at" timestamptz,
    "updated_at" timestamptz,
    "deleted_at" timestamptz,
    PRIMARY KEY ("id")
);

CREATE UNIQUE INDEX IF NOT EXISTS "user_idx_phone" ON "users" ("phone");

CREATE UNIQUE INDEX IF NOT EXISTS "user_idx_email" ON "users" ("email");

CREATE TABLE "user_addresses" (
    "id" uuid,
    "created_at" timestamptz,
    "updated_at" timestamptz,
    "deleted_at" timestamptz,
    "zip_code" text NOT NULL,
    "country" text,
    "state" text,
    "city" text,
    "street" text,
    "is_default" boolean,
    "user_id" uuid,
    PRIMARY KEY ("id"),
    CONSTRAINT "fk_users_user_address" FOREIGN KEY ("user_id") REFERENCES "users" ("id")
);