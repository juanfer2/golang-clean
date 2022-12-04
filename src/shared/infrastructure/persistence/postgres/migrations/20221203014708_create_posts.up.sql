BEGIN;
CREATE TABLE posts (
  id SERIAL
  --
  ,name   text
  ,description  varchar(255)
  --
  ,created_at timestamptz  NOT NULL  DEFAULT current_timestamp
  ,updated_at timestamptz  NOT NULL  DEFAULT current_timestamp
  -- soft delete
  ,deleted_at   timestamptz
  --
  ,PRIMARY KEY (id)
);

-- CREATE INDEX IF NOT EXISTS "idx_posts_deleted_at" ON "posts" ("deleted_at");
-- CREATE INDEX IF NOT EXISTS "idx_posts_name" ON "citizens" ("name");
COMMIT;
