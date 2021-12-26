-- Movie
DROP   TABLE    IF     EXISTS "public"."movie" CASCADE;
CREATE TABLE    IF NOT EXISTS "public"."movie" (
    "id"                    BIGSERIAL                NOT NULL,
    "created_at"            TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    "updated_at"            TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    "user_id"               BIGINT                   NOT NULL,
    "title"                 TEXT                     NOT NULL CHECK (CHARACTER_LENGTH("title") > 0),
    "overview"              TEXT                     NOT NULL CHECK (CHARACTER_LENGTH("overview") > 0),
    "genre"                 TEXT                     NOT NULL CHECK (CHARACTER_LENGTH("genre") > 0),
    "youtube_link_url"      TEXT                     NOT NULL CHECK (CHARACTER_LENGTH("youtube_link_url") > 0),
    "youtube_thumbnail_url" TEXT                     NOT NULL CHECK (CHARACTER_LENGTH("youtube_thumbnail_url") > 0),
    "grinning_score"        INTEGER                           CHECK ("grinning_score" > 0),
    PRIMARY KEY ("id")
);
