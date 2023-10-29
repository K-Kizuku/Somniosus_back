-- +goose Up
CREATE TABLE IF NOT EXISTS "authentications" (
    "id" VARCHAR(255) PRIMARY KEY,
    "hashed_password" VARCHAR(255),
    "expired_at" timestamp
);