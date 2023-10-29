-- +goose Up
CREATE TYPE "account_status" AS ENUM (
    'progressing',
    'general',
    'official',
    'subscriber'
);
CREATE TABLE IF NOT EXISTS "accounts" (
    "id" VARCHAR(255) PRIMARY KEY,
    "status" account_status,
    "display_id" VARCHAR(255),
    "name" VARCHAR(255),
    "description" text,
    "image_url" text,
    "birth_day" VARCHAR(63),
    "tel_num" VARCHAR(63),
    "website_url" text,
    "created_at" timestamp,
    "deleted_at" timestamp,
    "updated_at" timestamp
);