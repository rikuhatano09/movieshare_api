CREATE OR REPLACE FUNCTION set_update_time()
RETURNS trigger AS $$
    BEGIN
      NEW.updated_at = NOW();
      RETURN NEW;
    END;
$$ LANGUAGE plpgsql;

-- Movie
DROP TRIGGER IF EXISTS "movie_updated_at_trigger" ON "movie";
CREATE TRIGGER "movie_updated_at_trigger"
BEFORE UPDATE ON "movie"
FOR EACH ROW
EXECUTE PROCEDURE set_update_time();
