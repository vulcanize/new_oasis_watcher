CREATE TABLE oasis.kill (
  db_id            SERIAL,
  vulcanize_log_id INTEGER NOT NULL UNIQUE,
  id               INTEGER NOT NULL UNIQUE,
  pair             CHARACTER VARYING(66),
  gem              CHARACTER VARYING(66),
  lot              DECIMAL,
  pie              CHARACTER VARYING(66),
  bid              DECIMAL,
  guy              CHARACTER VARYING(66),
  block            INTEGER                  NOT NULL,
  time             TIMESTAMP WITH TIME ZONE NOT NULL,
  tx               CHARACTER VARYING(66)    NOT NULL,
  CONSTRAINT log_index_fk FOREIGN KEY (vulcanize_log_id)
  REFERENCES logs (id)
  ON DELETE CASCADE
);

CREATE INDEX kill_id_index
  ON oasis.kill (id);
