--
-- PostgreSQL database dump
--

SET statement_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SET check_function_bodies = false;
SET client_min_messages = warning;

--
-- Name: plpgsql; Type: EXTENSION; Schema: -; Owner: -
--

CREATE EXTENSION IF NOT EXISTS plpgsql WITH SCHEMA pg_catalog;


--
-- Name: EXTENSION plpgsql; Type: COMMENT; Schema: -; Owner: -
--

COMMENT ON EXTENSION plpgsql IS 'PL/pgSQL procedural language';


--
-- Name: hstore; Type: EXTENSION; Schema: -; Owner: -
--

CREATE EXTENSION IF NOT EXISTS hstore WITH SCHEMA public;


--
-- Name: EXTENSION hstore; Type: COMMENT; Schema: -; Owner: -
--

COMMENT ON EXTENSION hstore IS 'data type for storing sets of (key, value) pairs';


SET search_path = public, pg_catalog;

--
-- Name: event_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE event_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


SET default_tablespace = '';

SET default_with_oids = false;

--
-- Name: event; Type: TABLE; Schema: public; Owner: -; Tablespace: 
--

CREATE TABLE event (
    id bigint DEFAULT nextval('event_id_seq'::regclass) NOT NULL,
    created timestamp without time zone,
    updated timestamp without time zone
);


--
-- Name: event_event_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE event_event_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: event_event_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE event_event_seq OWNED BY event.id;


--
-- Name: event_player; Type: TABLE; Schema: public; Owner: -; Tablespace: 
--

CREATE TABLE event_player (
    id bigint NOT NULL,
    player_id integer NOT NULL
);


--
-- Name: event_player_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE event_player_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: event_player_id_seq1; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE event_player_id_seq1
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: event_player_id_seq1; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE event_player_id_seq1 OWNED BY event_player.id;


--
-- Name: event_role_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE event_role_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: event_role; Type: TABLE; Schema: public; Owner: -; Tablespace: 
--

CREATE TABLE event_role (
    id integer DEFAULT nextval('event_role_id_seq'::regclass) NOT NULL,
    event_type_id integer NOT NULL,
    name character varying(64) NOT NULL,
    description text
);


--
-- Name: event_tag; Type: TABLE; Schema: public; Owner: -; Tablespace: 
--

CREATE TABLE event_tag (
    id bigint NOT NULL,
    tag_id integer NOT NULL
);


--
-- Name: event_tag_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE event_tag_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: event_tag_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE event_tag_id_seq OWNED BY event_tag.id;


--
-- Name: event_to_game; Type: TABLE; Schema: public; Owner: -; Tablespace: 
--

CREATE TABLE event_to_game (
    id bigint NOT NULL,
    event_id integer,
    game_id integer
);


--
-- Name: event_to_game_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE event_to_game_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: event_to_game_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE event_to_game_id_seq OWNED BY event_to_game.id;


--
-- Name: event_to_player; Type: TABLE; Schema: public; Owner: -; Tablespace: 
--

CREATE TABLE event_to_player (
    id bigint NOT NULL,
    event_id integer,
    player_id integer,
    event_role_id integer
);


--
-- Name: event_to_player_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE event_to_player_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: event_to_player_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE event_to_player_id_seq OWNED BY event_to_player.id;


--
-- Name: event_type_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE event_type_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: event_type; Type: TABLE; Schema: public; Owner: -; Tablespace: 
--

CREATE TABLE event_type (
    id integer DEFAULT nextval('event_type_id_seq'::regclass) NOT NULL,
    name character varying(255) NOT NULL,
    description text,
    table_name character varying(255)
);


--
-- Name: game; Type: TABLE; Schema: public; Owner: -; Tablespace: 
--

CREATE TABLE game (
    id integer NOT NULL,
    name character varying(255),
    slug character varying(255),
    organization_id integer NOT NULL,
    timezone character varying(64),
    registration_start_time timestamp without time zone,
    registration_end_time timestamp without time zone,
    running_start_time timestamp without time zone,
    running_end_time timestamp without time zone,
    created timestamp without time zone,
    updated timestamp without time zone,
    CONSTRAINT game_registration_start_before_end CHECK ((registration_start_time < registration_end_time)),
    CONSTRAINT game_registration_start_before_running_end CHECK ((registration_start_time <= running_end_time)),
    CONSTRAINT game_running_start_before_end CHECK ((running_start_time < running_end_time))
);


--
-- Name: game_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE game_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: game_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE game_id_seq OWNED BY game.id;


--
-- Name: game_options; Type: TABLE; Schema: public; Owner: -; Tablespace: 
--

CREATE TABLE game_options (
    id integer NOT NULL,
    data hstore NOT NULL
);


--
-- Name: game_options_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE game_options_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: game_options_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE game_options_id_seq OWNED BY game_options.id;


--
-- Name: game_organization_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE game_organization_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: game_organization_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE game_organization_id_seq OWNED BY game.organization_id;


--
-- Name: human_code; Type: TABLE; Schema: public; Owner: -; Tablespace: 
--

CREATE TABLE human_code (
    id integer NOT NULL,
    code character varying(5) NOT NULL,
    created timestamp without time zone,
    updated timestamp without time zone
);


--
-- Name: human_code_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE human_code_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: human_code_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE human_code_id_seq OWNED BY human_code.id;


--
-- Name: member; Type: TABLE; Schema: public; Owner: -; Tablespace: 
--

CREATE TABLE member (
    id integer NOT NULL,
    user_id integer NOT NULL,
    organization_id integer NOT NULL,
    created timestamp without time zone,
    updated timestamp without time zone
);


--
-- Name: member_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE member_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: member_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE member_id_seq OWNED BY member.id;


--
-- Name: member_organization_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE member_organization_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: member_organization_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE member_organization_id_seq OWNED BY member.organization_id;


--
-- Name: member_user_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE member_user_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: member_user_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE member_user_id_seq OWNED BY member.user_id;


--
-- Name: organization; Type: TABLE; Schema: public; Owner: -; Tablespace: 
--

CREATE TABLE organization (
    id integer NOT NULL,
    name character varying(255),
    slug character varying(255),
    location character varying(255),
    default_timezone character varying(64),
    created timestamp without time zone,
    updated timestamp without time zone
);


--
-- Name: organization_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE organization_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: organization_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE organization_id_seq OWNED BY organization.id;


--
-- Name: oz; Type: TABLE; Schema: public; Owner: -; Tablespace: 
--

CREATE TABLE oz (
    id integer NOT NULL,
    confirmed boolean NOT NULL,
    created timestamp without time zone,
    updated timestamp without time zone
);


--
-- Name: oz_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE oz_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: oz_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE oz_id_seq OWNED BY oz.id;


--
-- Name: oz_pool; Type: TABLE; Schema: public; Owner: -; Tablespace: 
--

CREATE TABLE oz_pool (
    id integer NOT NULL,
    created timestamp without time zone,
    updated timestamp without time zone
);


--
-- Name: oz_pool_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE oz_pool_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: oz_pool_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE oz_pool_id_seq OWNED BY oz_pool.id;


--
-- Name: password_reset; Type: TABLE; Schema: public; Owner: -; Tablespace: 
--

CREATE TABLE password_reset (
    id integer NOT NULL,
    code character varying(36) NOT NULL,
    expires timestamp without time zone,
    created timestamp without time zone,
    updated timestamp without time zone
);


--
-- Name: password_reset_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE password_reset_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: password_reset_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE password_reset_id_seq OWNED BY password_reset.id;


--
-- Name: player; Type: TABLE; Schema: public; Owner: -; Tablespace: 
--

CREATE TABLE player (
    id integer NOT NULL,
    user_id integer NOT NULL,
    game_id integer NOT NULL,
    created timestamp without time zone,
    updated timestamp without time zone,
    last_fed timestamp(6) without time zone
);


--
-- Name: player_game_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE player_game_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: player_game_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE player_game_id_seq OWNED BY player.game_id;


--
-- Name: player_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE player_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: player_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE player_id_seq OWNED BY player.id;


--
-- Name: player_user_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE player_user_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: player_user_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE player_user_id_seq OWNED BY player.user_id;


--
-- Name: tag; Type: TABLE; Schema: public; Owner: -; Tablespace: 
--

CREATE TABLE tag (
    id integer NOT NULL,
    tagger_id integer NOT NULL,
    taggee_id integer NOT NULL,
    claimed timestamp without time zone,
    created timestamp without time zone,
    updated timestamp without time zone
);


--
-- Name: tag_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE tag_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: tag_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE tag_id_seq OWNED BY tag.id;


--
-- Name: tag_taggee_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE tag_taggee_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: tag_taggee_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE tag_taggee_id_seq OWNED BY tag.taggee_id;


--
-- Name: tag_tagger_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE tag_tagger_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: tag_tagger_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE tag_tagger_id_seq OWNED BY tag.tagger_id;


--
-- Name: user; Type: TABLE; Schema: public; Owner: -; Tablespace: 
--

CREATE TABLE "user" (
    id integer NOT NULL,
    email character varying(254),
    first_name character varying(255),
    last_name character varying(255),
    screen_name character varying(20),
    password character varying(60),
    api_key character varying(36),
    last_login timestamp without time zone,
    created timestamp without time zone DEFAULT now(),
    updated timestamp without time zone DEFAULT now()
);


--
-- Name: user_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE user_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: user_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE user_id_seq OWNED BY "user".id;


--
-- Name: id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY event_player ALTER COLUMN id SET DEFAULT nextval('event_player_id_seq1'::regclass);


--
-- Name: id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY event_tag ALTER COLUMN id SET DEFAULT nextval('event_tag_id_seq'::regclass);


--
-- Name: id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY event_to_game ALTER COLUMN id SET DEFAULT nextval('event_to_game_id_seq'::regclass);


--
-- Name: id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY event_to_player ALTER COLUMN id SET DEFAULT nextval('event_to_player_id_seq'::regclass);


--
-- Name: id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY game ALTER COLUMN id SET DEFAULT nextval('game_id_seq'::regclass);


--
-- Name: organization_id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY game ALTER COLUMN organization_id SET DEFAULT nextval('game_organization_id_seq'::regclass);


--
-- Name: id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY game_options ALTER COLUMN id SET DEFAULT nextval('game_options_id_seq'::regclass);


--
-- Name: id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY human_code ALTER COLUMN id SET DEFAULT nextval('human_code_id_seq'::regclass);


--
-- Name: id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY member ALTER COLUMN id SET DEFAULT nextval('member_id_seq'::regclass);


--
-- Name: user_id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY member ALTER COLUMN user_id SET DEFAULT nextval('member_user_id_seq'::regclass);


--
-- Name: organization_id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY member ALTER COLUMN organization_id SET DEFAULT nextval('member_organization_id_seq'::regclass);


--
-- Name: id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY organization ALTER COLUMN id SET DEFAULT nextval('organization_id_seq'::regclass);


--
-- Name: id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY oz ALTER COLUMN id SET DEFAULT nextval('oz_id_seq'::regclass);


--
-- Name: id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY oz_pool ALTER COLUMN id SET DEFAULT nextval('oz_pool_id_seq'::regclass);


--
-- Name: id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY password_reset ALTER COLUMN id SET DEFAULT nextval('password_reset_id_seq'::regclass);


--
-- Name: id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY player ALTER COLUMN id SET DEFAULT nextval('player_id_seq'::regclass);


--
-- Name: user_id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY player ALTER COLUMN user_id SET DEFAULT nextval('player_user_id_seq'::regclass);


--
-- Name: game_id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY player ALTER COLUMN game_id SET DEFAULT nextval('player_game_id_seq'::regclass);


--
-- Name: id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY tag ALTER COLUMN id SET DEFAULT nextval('tag_id_seq'::regclass);


--
-- Name: tagger_id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY tag ALTER COLUMN tagger_id SET DEFAULT nextval('tag_tagger_id_seq'::regclass);


--
-- Name: taggee_id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY tag ALTER COLUMN taggee_id SET DEFAULT nextval('tag_taggee_id_seq'::regclass);


--
-- Name: id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY "user" ALTER COLUMN id SET DEFAULT nextval('user_id_seq'::regclass);


--
-- Name: constraintname; Type: CONSTRAINT; Schema: public; Owner: -; Tablespace: 
--

ALTER TABLE ONLY event_type
    ADD CONSTRAINT constraintname UNIQUE (table_name);


--
-- Name: event_pkey; Type: CONSTRAINT; Schema: public; Owner: -; Tablespace: 
--

ALTER TABLE ONLY event
    ADD CONSTRAINT event_pkey PRIMARY KEY (id);


--
-- Name: event_player_copy_pkey; Type: CONSTRAINT; Schema: public; Owner: -; Tablespace: 
--

ALTER TABLE ONLY event_tag
    ADD CONSTRAINT event_player_copy_pkey PRIMARY KEY (id);


--
-- Name: event_player_pkey; Type: CONSTRAINT; Schema: public; Owner: -; Tablespace: 
--

ALTER TABLE ONLY event_player
    ADD CONSTRAINT event_player_pkey PRIMARY KEY (id);


--
-- Name: event_role_pkey; Type: CONSTRAINT; Schema: public; Owner: -; Tablespace: 
--

ALTER TABLE ONLY event_role
    ADD CONSTRAINT event_role_pkey PRIMARY KEY (id);


--
-- Name: event_to_game_pkey; Type: CONSTRAINT; Schema: public; Owner: -; Tablespace: 
--

ALTER TABLE ONLY event_to_game
    ADD CONSTRAINT event_to_game_pkey PRIMARY KEY (id);


--
-- Name: event_to_player_pkey; Type: CONSTRAINT; Schema: public; Owner: -; Tablespace: 
--

ALTER TABLE ONLY event_to_player
    ADD CONSTRAINT event_to_player_pkey PRIMARY KEY (id);


--
-- Name: event_type_name_key; Type: CONSTRAINT; Schema: public; Owner: -; Tablespace: 
--

ALTER TABLE ONLY event_type
    ADD CONSTRAINT event_type_name_key UNIQUE (name);


--
-- Name: event_type_pkey; Type: CONSTRAINT; Schema: public; Owner: -; Tablespace: 
--

ALTER TABLE ONLY event_type
    ADD CONSTRAINT event_type_pkey PRIMARY KEY (id);


--
-- Name: game_options_pkey; Type: CONSTRAINT; Schema: public; Owner: -; Tablespace: 
--

ALTER TABLE ONLY game_options
    ADD CONSTRAINT game_options_pkey PRIMARY KEY (id);


--
-- Name: game_pkey; Type: CONSTRAINT; Schema: public; Owner: -; Tablespace: 
--

ALTER TABLE ONLY game
    ADD CONSTRAINT game_pkey PRIMARY KEY (id);


--
-- Name: human_code_code_uniq; Type: CONSTRAINT; Schema: public; Owner: -; Tablespace: 
--

ALTER TABLE ONLY human_code
    ADD CONSTRAINT human_code_code_uniq UNIQUE (code);


--
-- Name: human_code_pkey; Type: CONSTRAINT; Schema: public; Owner: -; Tablespace: 
--

ALTER TABLE ONLY human_code
    ADD CONSTRAINT human_code_pkey PRIMARY KEY (id);


--
-- Name: member_pkey; Type: CONSTRAINT; Schema: public; Owner: -; Tablespace: 
--

ALTER TABLE ONLY member
    ADD CONSTRAINT member_pkey PRIMARY KEY (id);


--
-- Name: organization_pkey; Type: CONSTRAINT; Schema: public; Owner: -; Tablespace: 
--

ALTER TABLE ONLY organization
    ADD CONSTRAINT organization_pkey PRIMARY KEY (id);


--
-- Name: oz_pkey; Type: CONSTRAINT; Schema: public; Owner: -; Tablespace: 
--

ALTER TABLE ONLY oz
    ADD CONSTRAINT oz_pkey PRIMARY KEY (id);


--
-- Name: oz_pool_pkey; Type: CONSTRAINT; Schema: public; Owner: -; Tablespace: 
--

ALTER TABLE ONLY oz_pool
    ADD CONSTRAINT oz_pool_pkey PRIMARY KEY (id);


--
-- Name: password_reset_code_uniq; Type: CONSTRAINT; Schema: public; Owner: -; Tablespace: 
--

ALTER TABLE ONLY password_reset
    ADD CONSTRAINT password_reset_code_uniq UNIQUE (code);


--
-- Name: password_reset_pkey; Type: CONSTRAINT; Schema: public; Owner: -; Tablespace: 
--

ALTER TABLE ONLY password_reset
    ADD CONSTRAINT password_reset_pkey PRIMARY KEY (id);


--
-- Name: player_pkey; Type: CONSTRAINT; Schema: public; Owner: -; Tablespace: 
--

ALTER TABLE ONLY player
    ADD CONSTRAINT player_pkey PRIMARY KEY (id);


--
-- Name: tag_pkey; Type: CONSTRAINT; Schema: public; Owner: -; Tablespace: 
--

ALTER TABLE ONLY tag
    ADD CONSTRAINT tag_pkey PRIMARY KEY (id);


--
-- Name: user_pkey; Type: CONSTRAINT; Schema: public; Owner: -; Tablespace: 
--

ALTER TABLE ONLY "user"
    ADD CONSTRAINT user_pkey PRIMARY KEY (id);


--
-- Name: member_user_organization; Type: INDEX; Schema: public; Owner: -; Tablespace: 
--

CREATE UNIQUE INDEX member_user_organization ON member USING btree (user_id, organization_id);


--
-- Name: organization_name_idx; Type: INDEX; Schema: public; Owner: -; Tablespace: 
--

CREATE UNIQUE INDEX organization_name_idx ON organization USING btree (name);


--
-- Name: organization_slug_idx; Type: INDEX; Schema: public; Owner: -; Tablespace: 
--

CREATE UNIQUE INDEX organization_slug_idx ON organization USING btree (slug);


--
-- Name: player_user_game; Type: INDEX; Schema: public; Owner: -; Tablespace: 
--

CREATE UNIQUE INDEX player_user_game ON player USING btree (user_id, game_id);


--
-- Name: screen_name_idx; Type: INDEX; Schema: public; Owner: -; Tablespace: 
--

CREATE UNIQUE INDEX screen_name_idx ON "user" USING btree (screen_name);


--
-- Name: tag_claimed_idx; Type: INDEX; Schema: public; Owner: -; Tablespace: 
--

CREATE INDEX tag_claimed_idx ON tag USING btree (claimed);


--
-- Name: tag_taggee; Type: INDEX; Schema: public; Owner: -; Tablespace: 
--

CREATE UNIQUE INDEX tag_taggee ON tag USING btree (taggee_id);


--
-- Name: user_email_idx; Type: INDEX; Schema: public; Owner: -; Tablespace: 
--

CREATE UNIQUE INDEX user_email_idx ON "user" USING btree (email);


--
-- Name: event_role_event_type_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY event_role
    ADD CONSTRAINT event_role_event_type_id_fkey FOREIGN KEY (event_type_id) REFERENCES event_type(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: event_tag_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY event_tag
    ADD CONSTRAINT event_tag_id_fkey FOREIGN KEY (id) REFERENCES event(id);


--
-- Name: event_tag_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY event_player
    ADD CONSTRAINT event_tag_id_fkey FOREIGN KEY (id) REFERENCES event(id);


--
-- Name: event_tag_player_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY event_player
    ADD CONSTRAINT event_tag_player_id_fkey FOREIGN KEY (player_id) REFERENCES player(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: event_tag_tag_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY event_tag
    ADD CONSTRAINT event_tag_tag_id_fkey FOREIGN KEY (tag_id) REFERENCES tag(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: event_to_game_event_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY event_to_game
    ADD CONSTRAINT event_to_game_event_id_fkey FOREIGN KEY (event_id) REFERENCES event(id);


--
-- Name: event_to_game_game_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY event_to_game
    ADD CONSTRAINT event_to_game_game_id_fkey FOREIGN KEY (game_id) REFERENCES game(id);


--
-- Name: event_to_player_event_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY event_to_player
    ADD CONSTRAINT event_to_player_event_id_fkey FOREIGN KEY (event_id) REFERENCES event(id);


--
-- Name: event_to_player_event_role_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY event_to_player
    ADD CONSTRAINT event_to_player_event_role_id_fkey FOREIGN KEY (event_role_id) REFERENCES event_role(id);


--
-- Name: event_to_player_player_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY event_to_player
    ADD CONSTRAINT event_to_player_player_id_fkey FOREIGN KEY (player_id) REFERENCES player(id);


--
-- Name: game_options_fk_game; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY game_options
    ADD CONSTRAINT game_options_fk_game FOREIGN KEY (id) REFERENCES player(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: human_code_fk_player; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY human_code
    ADD CONSTRAINT human_code_fk_player FOREIGN KEY (id) REFERENCES player(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: member_fk_organization; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY member
    ADD CONSTRAINT member_fk_organization FOREIGN KEY (organization_id) REFERENCES organization(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: member_fk_user; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY member
    ADD CONSTRAINT member_fk_user FOREIGN KEY (user_id) REFERENCES "user"(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: oz_fk_player; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY oz
    ADD CONSTRAINT oz_fk_player FOREIGN KEY (id) REFERENCES player(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: oz_pool_fk_player; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY oz_pool
    ADD CONSTRAINT oz_pool_fk_player FOREIGN KEY (id) REFERENCES player(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: password_reset_fk_user; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY password_reset
    ADD CONSTRAINT password_reset_fk_user FOREIGN KEY (id) REFERENCES "user"(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: player_fk_game; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY player
    ADD CONSTRAINT player_fk_game FOREIGN KEY (game_id) REFERENCES game(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: player_fk_user; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY player
    ADD CONSTRAINT player_fk_user FOREIGN KEY (user_id) REFERENCES "user"(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: tag_fk_player_taggee; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY tag
    ADD CONSTRAINT tag_fk_player_taggee FOREIGN KEY (taggee_id) REFERENCES player(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: tag_fk_player_tagger; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY tag
    ADD CONSTRAINT tag_fk_player_tagger FOREIGN KEY (tagger_id) REFERENCES player(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: public; Type: ACL; Schema: -; Owner: -
--

REVOKE ALL ON SCHEMA public FROM PUBLIC;
REVOKE ALL ON SCHEMA public FROM postgres;
GRANT ALL ON SCHEMA public TO postgres;
GRANT ALL ON SCHEMA public TO PUBLIC;


--
-- PostgreSQL database dump complete
--

