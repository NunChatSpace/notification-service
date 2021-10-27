CREATE TABLE "expression"(
    "id"                        text,
    "created_at"                timestamptz NOT NULL,
    "updated_at"                timestamptz NOT NULL,
    "deleted_at"                timestamptz,
    "owner_id"                  text NOT NULL,
    "source_id"                 text NOT NULL,
    "destination_id"            text NOT NULL,
    "condition"                 text NOT NULL,
    "interval"                  text NOT NULL,
    "is_incondition"            text NOT NULL,
    "message_in_condition"      text NOT NULL,
    "message_out_condition"     text NOT NULL,

    PRIMARY KEY("id")
);