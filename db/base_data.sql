--
-- PostgreSQL database dump
--

SET statement_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SET check_function_bodies = false;
SET client_min_messages = warning;

SET search_path = public, pg_catalog;

--
-- Data for Name: event_type; Type: TABLE DATA; Schema: public; Owner: -
--

COPY event_type (id, name, description, table_name) FROM stdin;
1	tag	A zombie tagged a human	event_tag
\.


--
-- Data for Name: event_role; Type: TABLE DATA; Schema: public; Owner: -
--

COPY event_role (id, event_type_id, name, description) FROM stdin;
1	1	tagger	the zombie who tagged a human
2	1	taggee	the human tagged by a zombie
\.


--
-- PostgreSQL database dump complete
--

