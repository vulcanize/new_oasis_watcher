--
-- PostgreSQL database dump
--

-- Dumped from database version 10.1
-- Dumped by pg_dump version 10.1

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SET check_function_bodies = false;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: oasis; Type: SCHEMA; Schema: -; Owner: -
--

CREATE SCHEMA oasis;


--
-- Name: plpgsql; Type: EXTENSION; Schema: -; Owner: -
--

CREATE EXTENSION IF NOT EXISTS plpgsql WITH SCHEMA pg_catalog;


--
-- Name: EXTENSION plpgsql; Type: COMMENT; Schema: -; Owner: -
--

COMMENT ON EXTENSION plpgsql IS 'PL/pgSQL procedural language';


SET search_path = oasis, pg_catalog;

SET default_tablespace = '';

SET default_with_oids = false;

--
-- Name: kill; Type: TABLE; Schema: oasis; Owner: -
--

CREATE TABLE kill (
    db_id integer NOT NULL,
    vulcanize_log_id integer NOT NULL,
    id integer NOT NULL,
    pair character varying(66),
    gem character varying(66),
    lot numeric,
    pie character varying(66),
    bid numeric,
    guy character varying(66),
    block integer NOT NULL,
    "time" timestamp with time zone NOT NULL,
    tx character varying(66) NOT NULL
);


--
-- Name: kill_db_id_seq; Type: SEQUENCE; Schema: oasis; Owner: -
--

CREATE SEQUENCE kill_db_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: kill_db_id_seq; Type: SEQUENCE OWNED BY; Schema: oasis; Owner: -
--

ALTER SEQUENCE kill_db_id_seq OWNED BY kill.db_id;


--
-- Name: log_make; Type: TABLE; Schema: oasis; Owner: -
--

CREATE TABLE log_make (
    db_id integer NOT NULL,
    vulcanize_log_id integer NOT NULL,
    id integer NOT NULL,
    pair character varying(66),
    gem character varying(66),
    lot numeric,
    pie character varying(66),
    bid numeric,
    guy character varying(66),
    block integer NOT NULL,
    "time" timestamp with time zone NOT NULL,
    tx character varying(66) NOT NULL
);


--
-- Name: log_make_db_id_seq; Type: SEQUENCE; Schema: oasis; Owner: -
--

CREATE SEQUENCE log_make_db_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: log_make_db_id_seq; Type: SEQUENCE OWNED BY; Schema: oasis; Owner: -
--

ALTER SEQUENCE log_make_db_id_seq OWNED BY log_make.db_id;


--
-- Name: log_take; Type: TABLE; Schema: oasis; Owner: -
--

CREATE TABLE log_take (
    db_id integer NOT NULL,
    vulcanize_log_id integer NOT NULL,
    id integer NOT NULL,
    pair character varying(66),
    guy character varying(66),
    gem character varying(66),
    lot numeric,
    gal character varying(66),
    pie character varying(66),
    bid numeric,
    block integer NOT NULL,
    "time" timestamp with time zone NOT NULL,
    tx character varying(66) NOT NULL
);


--
-- Name: log_take_db_id_seq; Type: SEQUENCE; Schema: oasis; Owner: -
--

CREATE SEQUENCE log_take_db_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: log_take_db_id_seq; Type: SEQUENCE OWNED BY; Schema: oasis; Owner: -
--

ALTER SEQUENCE log_take_db_id_seq OWNED BY log_take.db_id;


--
-- Name: state; Type: VIEW; Schema: oasis; Owner: -
--

CREATE VIEW state AS
 WITH trade_state AS (
         SELECT log_make.db_id,
            log_make.vulcanize_log_id,
            log_make.id,
            log_make.pair,
            log_make.gem,
            log_make.lot,
            NULL::character varying AS gal,
            log_make.pie,
            log_make.bid,
            log_make.guy,
            log_make.block,
            log_make."time",
            log_make.tx
           FROM log_make
        UNION ALL
         SELECT log_take.db_id,
            log_take.vulcanize_log_id,
            log_take.id,
            log_take.pair,
            log_take.gem,
            log_take.lot,
            log_take.gal,
            log_take.pie,
            log_take.bid,
            log_take.guy,
            log_take.block,
            log_take."time",
            log_take.tx
           FROM log_take
  ORDER BY 3, 12
        )
 SELECT DISTINCT ON (ts.id) ts.db_id,
    ts.vulcanize_log_id,
    ts.id,
    ts.pair,
    ts.gem,
    ts.lot,
    ts.gal,
    ts.pie,
    ts.bid,
    ts.guy,
    ts.block,
    ts."time",
    ts.tx,
        CASE
            WHEN (tk.id IS NOT NULL) THEN true
            ELSE false
        END AS deleted
   FROM (trade_state ts
     LEFT JOIN kill tk ON ((ts.id = tk.id)))
  ORDER BY ts.id, ts.block DESC, ts."time" DESC;


SET search_path = public, pg_catalog;

--
-- Name: logs; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE logs (
    id integer NOT NULL,
    block_number bigint,
    address character varying(66),
    tx_hash character varying(66),
    index bigint,
    topic0 character varying(66),
    topic1 character varying(66),
    topic2 character varying(66),
    topic3 character varying(66),
    data text,
    receipt_id integer
);


--
-- Name: block_stats; Type: VIEW; Schema: public; Owner: -
--

CREATE VIEW block_stats AS
 SELECT max(logs.block_number) AS max_block,
    min(logs.block_number) AS min_block
   FROM logs;


--
-- Name: blocks; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE blocks (
    number bigint,
    gaslimit bigint,
    gasused bigint,
    "time" bigint,
    id integer NOT NULL,
    difficulty bigint,
    hash character varying(66),
    nonce character varying(20),
    parenthash character varying(66),
    size character varying,
    uncle_hash character varying(66),
    eth_node_id integer NOT NULL,
    is_final boolean,
    miner character varying(42),
    extra_data character varying,
    reward double precision,
    uncles_reward double precision
);


--
-- Name: blocks_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE blocks_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: blocks_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE blocks_id_seq OWNED BY blocks.id;


--
-- Name: eth_nodes; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE eth_nodes (
    id integer NOT NULL,
    genesis_block character varying(66),
    network_id numeric,
    eth_node_id character varying(128),
    client_name character varying
);


--
-- Name: log_filters; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE log_filters (
    id integer NOT NULL,
    name character varying NOT NULL,
    from_block bigint,
    to_block bigint,
    address character varying(66),
    topic0 character varying(66),
    topic1 character varying(66),
    topic2 character varying(66),
    topic3 character varying(66),
    CONSTRAINT log_filters_from_block_check CHECK ((from_block >= 0)),
    CONSTRAINT log_filters_name_check CHECK (((name)::text <> ''::text)),
    CONSTRAINT log_filters_to_block_check CHECK ((to_block >= 0))
);


--
-- Name: log_filters_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE log_filters_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: log_filters_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE log_filters_id_seq OWNED BY log_filters.id;


--
-- Name: logs_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE logs_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: logs_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE logs_id_seq OWNED BY logs.id;


--
-- Name: nodes_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE nodes_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: nodes_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE nodes_id_seq OWNED BY eth_nodes.id;


--
-- Name: receipts; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE receipts (
    id integer NOT NULL,
    transaction_id integer NOT NULL,
    contract_address character varying(42),
    cumulative_gas_used numeric,
    gas_used numeric,
    state_root character varying(66),
    status integer,
    tx_hash character varying(66)
);


--
-- Name: receipts_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE receipts_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: receipts_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE receipts_id_seq OWNED BY receipts.id;


--
-- Name: schema_migrations; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE schema_migrations (
    version bigint NOT NULL,
    dirty boolean NOT NULL
);


--
-- Name: transactions; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE transactions (
    id integer NOT NULL,
    hash character varying(66),
    nonce numeric,
    tx_to character varying(66),
    gaslimit numeric,
    gasprice numeric,
    value numeric,
    block_id integer NOT NULL,
    tx_from character varying(66),
    input_data character varying
);


--
-- Name: transactions_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE transactions_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: transactions_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE transactions_id_seq OWNED BY transactions.id;


--
-- Name: watched_contracts; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE watched_contracts (
    contract_id integer NOT NULL,
    contract_hash character varying(66),
    contract_abi json
);


--
-- Name: watched_contracts_contract_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE watched_contracts_contract_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: watched_contracts_contract_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE watched_contracts_contract_id_seq OWNED BY watched_contracts.contract_id;


--
-- Name: watched_event_logs; Type: VIEW; Schema: public; Owner: -
--

CREATE VIEW watched_event_logs AS
 SELECT log_filters.name,
    logs.id,
    logs.block_number,
    logs.address,
    logs.tx_hash,
    logs.index,
    logs.topic0,
    logs.topic1,
    logs.topic2,
    logs.topic3,
    logs.data,
    logs.receipt_id
   FROM ((log_filters
     CROSS JOIN block_stats)
     JOIN logs ON ((((logs.address)::text = (log_filters.address)::text) AND (logs.block_number >= COALESCE(log_filters.from_block, block_stats.min_block)) AND (logs.block_number <= COALESCE(log_filters.to_block, block_stats.max_block)))))
  WHERE ((((log_filters.topic0)::text = (logs.topic0)::text) OR (log_filters.topic0 IS NULL)) AND (((log_filters.topic1)::text = (logs.topic1)::text) OR (log_filters.topic1 IS NULL)) AND (((log_filters.topic2)::text = (logs.topic2)::text) OR (log_filters.topic2 IS NULL)) AND (((log_filters.topic3)::text = (logs.topic3)::text) OR (log_filters.topic3 IS NULL)));


SET search_path = oasis, pg_catalog;

--
-- Name: kill db_id; Type: DEFAULT; Schema: oasis; Owner: -
--

ALTER TABLE ONLY kill ALTER COLUMN db_id SET DEFAULT nextval('kill_db_id_seq'::regclass);


--
-- Name: log_make db_id; Type: DEFAULT; Schema: oasis; Owner: -
--

ALTER TABLE ONLY log_make ALTER COLUMN db_id SET DEFAULT nextval('log_make_db_id_seq'::regclass);


--
-- Name: log_take db_id; Type: DEFAULT; Schema: oasis; Owner: -
--

ALTER TABLE ONLY log_take ALTER COLUMN db_id SET DEFAULT nextval('log_take_db_id_seq'::regclass);


SET search_path = public, pg_catalog;

--
-- Name: blocks id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY blocks ALTER COLUMN id SET DEFAULT nextval('blocks_id_seq'::regclass);


--
-- Name: eth_nodes id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY eth_nodes ALTER COLUMN id SET DEFAULT nextval('nodes_id_seq'::regclass);


--
-- Name: log_filters id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY log_filters ALTER COLUMN id SET DEFAULT nextval('log_filters_id_seq'::regclass);


--
-- Name: logs id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY logs ALTER COLUMN id SET DEFAULT nextval('logs_id_seq'::regclass);


--
-- Name: receipts id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY receipts ALTER COLUMN id SET DEFAULT nextval('receipts_id_seq'::regclass);


--
-- Name: transactions id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY transactions ALTER COLUMN id SET DEFAULT nextval('transactions_id_seq'::regclass);


--
-- Name: watched_contracts contract_id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY watched_contracts ALTER COLUMN contract_id SET DEFAULT nextval('watched_contracts_contract_id_seq'::regclass);


SET search_path = oasis, pg_catalog;

--
-- Name: kill kill_id_key; Type: CONSTRAINT; Schema: oasis; Owner: -
--

ALTER TABLE ONLY kill
    ADD CONSTRAINT kill_id_key UNIQUE (id);


--
-- Name: kill kill_vulcanize_log_id_key; Type: CONSTRAINT; Schema: oasis; Owner: -
--

ALTER TABLE ONLY kill
    ADD CONSTRAINT kill_vulcanize_log_id_key UNIQUE (vulcanize_log_id);


--
-- Name: log_make log_make_id_key; Type: CONSTRAINT; Schema: oasis; Owner: -
--

ALTER TABLE ONLY log_make
    ADD CONSTRAINT log_make_id_key UNIQUE (id);


--
-- Name: log_make log_make_vulcanize_log_id_key; Type: CONSTRAINT; Schema: oasis; Owner: -
--

ALTER TABLE ONLY log_make
    ADD CONSTRAINT log_make_vulcanize_log_id_key UNIQUE (vulcanize_log_id);


--
-- Name: log_take log_take_vulcanize_log_id_key; Type: CONSTRAINT; Schema: oasis; Owner: -
--

ALTER TABLE ONLY log_take
    ADD CONSTRAINT log_take_vulcanize_log_id_key UNIQUE (vulcanize_log_id);


SET search_path = public, pg_catalog;

--
-- Name: blocks blocks_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY blocks
    ADD CONSTRAINT blocks_pkey PRIMARY KEY (id);


--
-- Name: watched_contracts contract_hash_uc; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY watched_contracts
    ADD CONSTRAINT contract_hash_uc UNIQUE (contract_hash);


--
-- Name: blocks eth_node_id_block_number_uc; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY blocks
    ADD CONSTRAINT eth_node_id_block_number_uc UNIQUE (number, eth_node_id);


--
-- Name: eth_nodes eth_node_uc; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY eth_nodes
    ADD CONSTRAINT eth_node_uc UNIQUE (genesis_block, network_id, eth_node_id);


--
-- Name: logs logs_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY logs
    ADD CONSTRAINT logs_pkey PRIMARY KEY (id);


--
-- Name: log_filters name_uc; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY log_filters
    ADD CONSTRAINT name_uc UNIQUE (name);


--
-- Name: eth_nodes nodes_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY eth_nodes
    ADD CONSTRAINT nodes_pkey PRIMARY KEY (id);


--
-- Name: receipts receipts_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY receipts
    ADD CONSTRAINT receipts_pkey PRIMARY KEY (id);


--
-- Name: schema_migrations schema_migrations_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY schema_migrations
    ADD CONSTRAINT schema_migrations_pkey PRIMARY KEY (version);


--
-- Name: transactions transactions_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY transactions
    ADD CONSTRAINT transactions_pkey PRIMARY KEY (id);


--
-- Name: watched_contracts watched_contracts_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY watched_contracts
    ADD CONSTRAINT watched_contracts_pkey PRIMARY KEY (contract_id);


SET search_path = oasis, pg_catalog;

--
-- Name: kill_id_index; Type: INDEX; Schema: oasis; Owner: -
--

CREATE INDEX kill_id_index ON kill USING btree (id);


--
-- Name: log_make_guy_index; Type: INDEX; Schema: oasis; Owner: -
--

CREATE INDEX log_make_guy_index ON log_make USING btree (guy);


--
-- Name: log_make_pair_index; Type: INDEX; Schema: oasis; Owner: -
--

CREATE INDEX log_make_pair_index ON log_make USING btree (pair);


--
-- Name: log_take_gal_index; Type: INDEX; Schema: oasis; Owner: -
--

CREATE INDEX log_take_gal_index ON log_take USING btree (gal);


--
-- Name: log_take_gem_index; Type: INDEX; Schema: oasis; Owner: -
--

CREATE INDEX log_take_gem_index ON log_take USING btree (gem);


--
-- Name: log_take_guy_index; Type: INDEX; Schema: oasis; Owner: -
--

CREATE INDEX log_take_guy_index ON log_take USING btree (guy);


--
-- Name: log_take_id_index; Type: INDEX; Schema: oasis; Owner: -
--

CREATE INDEX log_take_id_index ON log_take USING btree (id);


--
-- Name: log_take_pair_index; Type: INDEX; Schema: oasis; Owner: -
--

CREATE INDEX log_take_pair_index ON log_take USING btree (pair);


--
-- Name: log_take_pie_index; Type: INDEX; Schema: oasis; Owner: -
--

CREATE INDEX log_take_pie_index ON log_take USING btree (pie);


SET search_path = public, pg_catalog;

--
-- Name: block_id_index; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX block_id_index ON transactions USING btree (block_id);


--
-- Name: block_number_index; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX block_number_index ON blocks USING btree (number);


--
-- Name: node_id_index; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX node_id_index ON blocks USING btree (eth_node_id);


--
-- Name: transaction_id_index; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX transaction_id_index ON receipts USING btree (transaction_id);


--
-- Name: tx_from_index; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX tx_from_index ON transactions USING btree (tx_from);


--
-- Name: tx_to_index; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX tx_to_index ON transactions USING btree (tx_to);


SET search_path = oasis, pg_catalog;

--
-- Name: kill log_index_fk; Type: FK CONSTRAINT; Schema: oasis; Owner: -
--

ALTER TABLE ONLY kill
    ADD CONSTRAINT log_index_fk FOREIGN KEY (vulcanize_log_id) REFERENCES public.logs(id) ON DELETE CASCADE;


--
-- Name: log_make log_index_fk; Type: FK CONSTRAINT; Schema: oasis; Owner: -
--

ALTER TABLE ONLY log_make
    ADD CONSTRAINT log_index_fk FOREIGN KEY (vulcanize_log_id) REFERENCES public.logs(id) ON DELETE CASCADE;


--
-- Name: log_take log_index_fk; Type: FK CONSTRAINT; Schema: oasis; Owner: -
--

ALTER TABLE ONLY log_take
    ADD CONSTRAINT log_index_fk FOREIGN KEY (vulcanize_log_id) REFERENCES public.logs(id) ON DELETE CASCADE;


SET search_path = public, pg_catalog;

--
-- Name: transactions blocks_fk; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY transactions
    ADD CONSTRAINT blocks_fk FOREIGN KEY (block_id) REFERENCES blocks(id) ON DELETE CASCADE;


--
-- Name: blocks node_fk; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY blocks
    ADD CONSTRAINT node_fk FOREIGN KEY (eth_node_id) REFERENCES eth_nodes(id) ON DELETE CASCADE;


--
-- Name: logs receipts_fk; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY logs
    ADD CONSTRAINT receipts_fk FOREIGN KEY (receipt_id) REFERENCES receipts(id) ON DELETE CASCADE;


--
-- Name: receipts transaction_fk; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY receipts
    ADD CONSTRAINT transaction_fk FOREIGN KEY (transaction_id) REFERENCES transactions(id) ON DELETE CASCADE;


--
-- PostgreSQL database dump complete
--

