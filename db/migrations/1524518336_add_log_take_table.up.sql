CREATE TABLE oasis.log_take (
  db_id            SERIAL,
  vulcanize_log_id INTEGER NOT NULL UNIQUE,
  id               INTEGER NOT NULL,
  pair             CHARACTER VARYING(66),
  guy              CHARACTER VARYING(66),
  gem              CHARACTER VARYING(66),
  lot              DECIMAL,
  gal              CHARACTER VARYING(66),
  pie              CHARACTER VARYING(66),
  bid              DECIMAL,
  block            INTEGER                  NOT NULL,
  time             TIMESTAMP WITH TIME ZONE NOT NULL,
  tx               CHARACTER VARYING(66)    NOT NULL,
  CONSTRAINT log_index_fk FOREIGN KEY (vulcanize_log_id)
  REFERENCES logs (id)
  ON DELETE CASCADE
);

CREATE INDEX log_take_id_index
  ON oasis.log_take (id);
CREATE INDEX log_take_pair_index
  ON oasis.log_take (pair);
CREATE INDEX log_take_gem_index
  ON oasis.log_take (gem);
CREATE INDEX log_take_pie_index
  ON oasis.log_take (pie);
CREATE INDEX log_take_guy_index
  ON oasis.log_take (guy);
CREATE INDEX log_take_gal_index
  ON oasis.log_take (gal);
