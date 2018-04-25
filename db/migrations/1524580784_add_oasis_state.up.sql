CREATE VIEW oasis.state AS (
  WITH trade_state AS (
    SELECT
      db_id,
      vulcanize_log_id,
      id,
      pair,
      gem,
      lot,
      NULL AS gal,
      pie,
      bid,
      guy,
      block,
      time,
      tx
    FROM oasis.log_make
    UNION ALL
    SELECT
      db_id,
      vulcanize_log_id,
      id,
      pair,
      gem,
      lot,
      gal,
      pie,
      bid,
      guy,
      block,
      time,
      tx
    FROM oasis.log_take
    ORDER BY id, time
  )
  SELECT DISTINCT ON (id)
    ts.db_id,
    ts.vulcanize_log_id,
    ts.id,
    ts.pair,
    ts.gem,
    ts.lot,
    gal,
    ts.pie,
    ts.bid,
    ts.guy,
    ts.block,
    ts.time,
    ts.tx,
    CASE WHEN tk.id NOTNULL
      THEN TRUE
    ELSE FALSE END AS deleted
  FROM trade_state ts
    LEFT JOIN oasis.kill tk
      ON ts.id = tk.id
  ORDER BY id, block DESC, TIME DESC
);