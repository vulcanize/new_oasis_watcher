BEGIN;
CREATE SCHEMA oasis;

CREATE TABLE oasis.trade (
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

CREATE INDEX trade_id_index
  ON oasis.trade (id);
CREATE INDEX trade_pair_index
  ON oasis.trade (pair);
CREATE INDEX trade_gem_index
  ON oasis.trade (gem);
CREATE INDEX trade_pie_index
  ON oasis.trade (pie);
CREATE INDEX trade_guy_index
  ON oasis.trade (guy);
CREATE INDEX trade_gal_index
  ON oasis.trade (gal);
COMMIT;