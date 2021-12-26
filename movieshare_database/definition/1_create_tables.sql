-- Movie
DROP   TABLE    IF     EXISTS "public"."movie" CASCADE;
CREATE TABLE    IF NOT EXISTS "public"."movie" (
    "id"                    BIGSERIAL                NOT NULL,
    "created_at"            TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    "updated_at"            TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    "user_id"               TEXT                     NOT NULL CHECK (CHARACTER_LENGTH("user_id") > 0),
    "user_name"             TEXT                     NOT NULL CHECK (CHARACTER_LENGTH("user_name") > 0),
    "title"                 TEXT                     NOT NULL CHECK (CHARACTER_LENGTH("title") > 0),
    "overview"              TEXT                     NOT NULL CHECK (CHARACTER_LENGTH("overview") > 0),
    "genre"                 TEXT                     NOT NULL CHECK (CHARACTER_LENGTH("genre") > 0),
    "youtube_title_id"      TEXT                     NOT NULL CHECK (CHARACTER_LENGTH("youtube_title_id") > 0),
    "grinning_score"        INTEGER                           CHECK ("grinning_score" > 0),
    PRIMARY KEY ("id")
);
