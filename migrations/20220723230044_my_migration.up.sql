-- modify "leads" table
ALTER TABLE "leads" ADD COLUMN "create_time" timestamptz NOT NULL, ADD COLUMN "update_time" timestamptz NOT NULL;
