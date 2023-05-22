--
-- PostgreSQL database dump
--

-- Dumped from database version 14.5 (Debian 14.5-1.pgdg110+1)
-- Dumped by pg_dump version 14.5 (Homebrew)

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: gombs; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.gombs (
    id integer NOT NULL,
    gombs_title character varying(255),
    created_at timestamp without time zone,
    updated_at timestamp without time zone
);


--
-- Name: gombs_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

ALTER TABLE public.gombs ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.gombs_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Name: foods; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.foods (
    id integer NOT NULL,
    fdb_id character varying(32),
    food_title character varying(128),
    food_description text,
    food_image character varying(255),
    created_at timestamp without time zone,
    updated_at timestamp without time zone
);


--
-- Name: foods_gombs; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.foods_gombs (
    id integer NOT NULL,
    foods_id integer,
    gombs_id integer
);


--
-- Name: foods_gombs_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

ALTER TABLE public.foods_gombs ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.foods_gombs_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Name: foods_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

ALTER TABLE public.foods ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.foods_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Name: users; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.users (
    id integer NOT NULL,
    first_name character varying(255),
    last_name character varying(255),
    email character varying(255),
    password character varying(255),
    created_at timestamp without time zone,
    updated_at timestamp without time zone
);


--
-- Name: users_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

ALTER TABLE public.users ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.users_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Data for Name: gombs; Type: TABLE DATA; Schema: public; Owner: -
--

COPY public.gombs (id, gombs_title, created_at, updated_at) FROM stdin;
1	Greens	2023-05-20 00:00:00	2023-05-20 00:00:00
2	Onions	2023-05-20 00:00:00	2023-05-20 00:00:00
3	Mushrooms	2023-05-20 00:00:00	2023-05-20 00:00:00
4	Beans	2023-05-20 00:00:00	2023-05-20 00:00:00
5	Berries	2023-05-20 00:00:00	2023-05-20 00:00:00
6	Seeds	2023-05-20 00:00:00	2023-05-20 00:00:00
7   Fruits	2023-05-20 00:00:00	2023-05-20 00:00:00
8	Grains	2023-05-20 00:00:00	2023-05-20 00:00:00
9	AntiNutri	2023-05-20 00:00:00	2023-05-20 00:00:00
\.


--
-- Data for Name: foods; Type: TABLE DATA; Schema: public; Owner: -
--

COPY public.foods (id, fdb_id, food_title, food_description, food_image, created_at, updated_at) FROM stdin;
1	72119190    Kale_raw    Kale_raw    /homepic.jpg    2023-05-20 00:00:00	2023-05-20 00:00:00
2   72119211    Kale_fresh_raw  Kale freshraw   /homepic.jpg    2023-05-20 00:00:00	2023-05-20 00:00:00
3   75103000    Cabbage_green_raw   Cabbage, green _raw /homepic.jpg    2023-05-20 00:00:00	2023-05-20 00:00:00
4   75115000    Mushrooms   Mushrooms raw_  /homepic.jpg    2023-05-20 00:00:00	2023-05-20 00:00:00
\.


--
-- Data for Name: foods_gombs; Type: TABLE DATA; Schema: public; Owner: -
--

COPY public.foods_gombs (id, foods_id, gombs_id) FROM stdin;
1	1	1
2	2	1
3	3	1
4	4	3
\.


--
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: -
--

COPY public.users (id, first_name, last_name, email, password, created_at, updated_at) FROM stdin;
1	Admin	User	admin@nutrintel.com	$2a$14$wVsaPvJnJJsomWArouWCtusem6S/.Gauq/GjOIEHpyh2DAMmso1wy	2023-05-20 00:00:00	2023-05-20 00:00:00
\.


--
-- Name: gombs_id_seq; Type: SEQUENCE SET; Schema: public; Owner: -
--

SELECT pg_catalog.setval('public.gombs_id_seq', 9, true);


--
-- Name: foods_gombs_id_seq; Type: SEQUENCE SET; Schema: public; Owner: -
--

SELECT pg_catalog.setval('public.foods_gombs_id_seq', 4, true);


--
-- Name: foods_id_seq; Type: SEQUENCE SET; Schema: public; Owner: -
--

SELECT pg_catalog.setval('public.foods_id_seq', 4, true);


--
-- Name: users_id_seq; Type: SEQUENCE SET; Schema: public; Owner: -
--

SELECT pg_catalog.setval('public.users_id_seq', 1, true);


--
-- Name: gombs gombs_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.gombs
    ADD CONSTRAINT gombs_pkey PRIMARY KEY (id);


--
-- Name: foods_gombs foods_gombs_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.foods_gombs
    ADD CONSTRAINT foods_gombs_pkey PRIMARY KEY (id);


--
-- Name: foods foods_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.foods
    ADD CONSTRAINT foods_pkey PRIMARY KEY (id);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- Name: foods_gombs foods_gombs_gomb_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.foods_gombs
    ADD CONSTRAINT foods_gombs_gomb_id_fkey FOREIGN KEY (gombs_id) REFERENCES public.gombs(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: foods_gombs foods_gombs_food_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.foods_gombs
    ADD CONSTRAINT foods_gombs_food_id_fkey FOREIGN KEY (foods_id) REFERENCES public.foods(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- PostgreSQL database dump complete
--

