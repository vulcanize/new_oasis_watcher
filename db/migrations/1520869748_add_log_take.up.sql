BEGIN;
CREATE SCHEMA oasis;

CREATE TABLE oasis.log_takes (
  id            SERIAL,
  eth_log_id    INTEGER,
  oasis_log_id  VARCHAR(65),
  pair          VARCHAR(65),
  maker         VARCHAR(42),
  have_token    VARCHAR(42),
  want_token    VARCHAR(42),
  taker         VARCHAR(42),
  take_amount   NUMERIC,
  give_amount   NUMERIC,
  TIMESTAMP     NUMERIC,
  CONSTRAINT log_index_fk FOREIGN KEY (eth_log_id)
  REFERENCES logs (id)
  ON DELETE CASCADE
);
COMMIT;
