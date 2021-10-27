CREATE TABLE "destination" (
    "id"                        text,
    "created_at"                timestamptz NOT NULL,
    "updated_at"                timestamptz NOT NULL,
    "deleted_at"                timestamptz,
    "owner_id"                  text NOT NULL,
    "name"                      text NOT NULL,
    "type"                      text NOT NULL,
    "topic"                     text NOT NULL,
    "data"                      text NOT NULL,
    
    PRIMARY KEY("id")
);