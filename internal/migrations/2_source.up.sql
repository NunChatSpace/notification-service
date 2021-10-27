CREATE TABLE "source" (
    "id"                        text,
    "created_at"                timestamptz NOT NULL,
    "updated_at"                timestamptz NOT NULL,
    "deleted_at"                timestamptz,
    "owner_id"                  text NOT NULL,
    "server_address"            text NOT NULL,
    "database_name"             text NOT NULL,
    "table_name"                text NOT NULL,
    "data_key"                  text NOT NULL,
    PRIMARY KEY("id")
);
CREATE INDEX "idx_source_deleted_at" ON "source" ("deleted_at");
