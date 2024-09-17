--
-- PostgreSQL database cluster dump
--

-- Started on 2024-06-11 18:23:42

SET default_transaction_read_only = off;

SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;

--
-- Roles
--

CREATE ROLE postgres;
ALTER ROLE postgres WITH SUPERUSER INHERIT CREATEROLE CREATEDB LOGIN REPLICATION BYPASSRLS;

--
-- User Configurations
--








--
-- Databases
--

--
-- Database "template1" dump
--

\connect template1

--
-- PostgreSQL database dump
--

-- Dumped from database version 16.1
-- Dumped by pg_dump version 16.1

-- Started on 2024-06-11 18:23:42

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

-- Completed on 2024-06-11 18:23:42

--
-- PostgreSQL database dump complete
--

--
-- Database "fismed" dump
--

--
-- PostgreSQL database dump
--

-- Dumped from database version 16.1
-- Dumped by pg_dump version 16.1

-- Started on 2024-06-11 18:23:42

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

--
-- TOC entry 4858 (class 1262 OID 16398)
-- Name: fismed; Type: DATABASE; Schema: -; Owner: postgres
--

CREATE DATABASE fismed WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE_PROVIDER = libc LOCALE = 'English_Indonesia.1252';


ALTER DATABASE fismed OWNER TO postgres;

\connect fismed

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
-- TOC entry 216 (class 1259 OID 16413)
-- Name: user_category; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.user_category (
    id integer NOT NULL,
    role character varying(50) NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    created_by character varying(255),
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    updated_by character varying(255)
);


ALTER TABLE public.user_category OWNER TO postgres;

--
-- TOC entry 215 (class 1259 OID 16412)
-- Name: user_category_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.user_category_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.user_category_id_seq OWNER TO postgres;

--
-- TOC entry 4859 (class 0 OID 0)
-- Dependencies: 215
-- Name: user_category_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.user_category_id_seq OWNED BY public.user_category.id;


--
-- TOC entry 218 (class 1259 OID 16424)
-- Name: users; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.users (
    id integer NOT NULL,
    username character varying(50) NOT NULL,
    password character varying(255),
    role_id integer,
    token text NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    created_by character varying(255),
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    updated_by character varying(255)
);


ALTER TABLE public.users OWNER TO postgres;

--
-- TOC entry 217 (class 1259 OID 16423)
-- Name: user_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.user_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.user_id_seq OWNER TO postgres;

--
-- TOC entry 4860 (class 0 OID 0)
-- Dependencies: 217
-- Name: user_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.user_id_seq OWNED BY public.users.id;


--
-- TOC entry 4693 (class 2604 OID 16416)
-- Name: user_category id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.user_category ALTER COLUMN id SET DEFAULT nextval('public.user_category_id_seq'::regclass);


--
-- TOC entry 4696 (class 2604 OID 16427)
-- Name: users id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.user_id_seq'::regclass);


--
-- TOC entry 4850 (class 0 OID 16413)
-- Dependencies: 216
-- Data for Name: user_category; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.user_category (id, role, created_at, created_by, updated_at, updated_by) FROM stdin;
1	SALES	2024-06-11 17:11:16.111105	system	2024-06-11 17:11:16.111105	system
2	ADMIN	2024-06-11 17:12:52.189158	system	2024-06-11 17:12:52.189158	system
3	LOGISTIK	2024-06-11 17:12:52.189158	system	2024-06-11 17:12:52.189158	system
4	KEUANGAN	2024-06-11 17:12:52.189158	system	2024-06-11 17:12:52.189158	system
5	SUPER ADMIN	2024-06-11 17:12:52.189158	system	2024-06-11 17:12:52.189158	system
\.


--
-- TOC entry 4852 (class 0 OID 16424)
-- Dependencies: 218
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.users (id, username, password, role_id, token, created_at, created_by, updated_at, updated_by) FROM stdin;
1	sales1	0ad80eb119d9bf7775aa23786b05b391	1	eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTgxMDcxNDAsImlhdCI6MTcxODEwMzU0MCwic3ViIjoiMSJ9.9guuPCNghxCswpvyugzOwhHAkJyqvqA0LR_2am3WELI	2024-06-11 17:15:48.925181	system	2024-06-11 17:15:48.925181	system
\.


--
-- TOC entry 4861 (class 0 OID 0)
-- Dependencies: 215
-- Name: user_category_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.user_category_id_seq', 5, true);


--
-- TOC entry 4862 (class 0 OID 0)
-- Dependencies: 217
-- Name: user_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.user_id_seq', 1, true);


--
-- TOC entry 4700 (class 2606 OID 16422)
-- Name: user_category user_category_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.user_category
    ADD CONSTRAINT user_category_pkey PRIMARY KEY (id);


--
-- TOC entry 4702 (class 2606 OID 16433)
-- Name: users user_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT user_pkey PRIMARY KEY (id);


--
-- TOC entry 4704 (class 2606 OID 16435)
-- Name: users user_username_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT user_username_key UNIQUE (username);


--
-- TOC entry 4705 (class 2606 OID 16436)
-- Name: users user_role_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT user_role_id_fkey FOREIGN KEY (role_id) REFERENCES public.user_category(id);


-- Completed on 2024-06-11 18:23:43

--
-- PostgreSQL database dump complete
--

-- Completed on 2024-06-11 18:23:43

--
-- PostgreSQL database cluster dump complete
--

