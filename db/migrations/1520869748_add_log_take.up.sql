BEGIN;
CREATE SCHEMA oasis;

CREATE TABLE oasis.log_takes (
  id           SERIAL,
  eth_log_id   INTEGER,
  oasis_log_id VARCHAR(65),
  pair         VARCHAR(65),
  maker        VARCHAR(42),
  have_token   VARCHAR(42),
  want_token   VARCHAR(42),
  taker        VARCHAR(42),
  take_amount  NUMERIC,
  give_amount  NUMERIC,
  block        NUMERIC,
  TIMESTAMP    NUMERIC,
  CONSTRAINT log_index_fk FOREIGN KEY (eth_log_id)
  REFERENCES logs (id)
  ON DELETE CASCADE
);

CREATE VIEW oasis.log_takes_with_status AS
  SELECT
    eth_log_id,
    oasis_log_id,
    pair,
    maker,
    have_token,
    want_token,
    taker,
    take_amount,
    give_amount,
    block,
    timestamp,
    tx_hash,
    is_final
  FROM oasis.log_takes
    JOIN logs l ON log_takes.eth_log_id = l.id
    JOIN blocks b ON log_takes.block = b.number;

COMMIT;
