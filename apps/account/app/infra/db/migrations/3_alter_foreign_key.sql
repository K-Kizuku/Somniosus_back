-- +goose Up
ALTER TABLE "authentication"
ADD FOREIGN KEY ("id") REFERENCES "accounts" ("id");