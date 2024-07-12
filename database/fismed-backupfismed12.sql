--
-- PostgreSQL database dump
--

-- Dumped from database version 14.12
-- Dumped by pg_dump version 16.1

-- Started on 2024-07-12 14:47:47

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
-- TOC entry 5 (class 2615 OID 16412)
-- Name: public; Type: SCHEMA; Schema: -; Owner: pg_database_owner
--

CREATE SCHEMA public;


ALTER SCHEMA public OWNER TO pg_database_owner;

--
-- TOC entry 3535 (class 0 OID 0)
-- Dependencies: 5
-- Name: SCHEMA public; Type: COMMENT; Schema: -; Owner: pg_database_owner
--

COMMENT ON SCHEMA public IS 'standard public schema';


SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- TOC entry 209 (class 1259 OID 16443)
-- Name: customer; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.customer (
    id integer NOT NULL,
    name character varying,
    address_company text,
    npwp_address text,
    npwp character varying,
    ipak_number character varying,
    facture_address character varying,
    city_facture character varying,
    zip_code_facture character varying,
    number_phone_facture character varying,
    email_facture character varying,
    fax_facture character varying,
    pic_facture character varying,
    item_address character varying,
    city_item character varying,
    zip_code_item character varying,
    number_phone_item character varying,
    email_item character varying,
    fax_item character varying,
    pic_item character varying,
    contact_person character varying,
    tax_code_id integer,
    top character varying,
    created_at timestamp without time zone,
    created_by character varying,
    updated_at timestamp without time zone,
    updated_by character varying,
    handphone character varying
);


ALTER TABLE public.customer OWNER TO postgres;

--
-- TOC entry 210 (class 1259 OID 16448)
-- Name: customer_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.customer_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.customer_id_seq OWNER TO postgres;

--
-- TOC entry 3537 (class 0 OID 0)
-- Dependencies: 210
-- Name: customer_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.customer_id_seq OWNED BY public.customer.id;


--
-- TOC entry 211 (class 1259 OID 16450)
-- Name: divisi; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.divisi (
    id integer NOT NULL,
    name character varying NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    created_by character varying NOT NULL,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_by character varying NOT NULL
);


ALTER TABLE public.divisi OWNER TO postgres;

--
-- TOC entry 212 (class 1259 OID 16457)
-- Name: divisi_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.divisi_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.divisi_id_seq OWNER TO postgres;

--
-- TOC entry 3538 (class 0 OID 0)
-- Dependencies: 212
-- Name: divisi_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.divisi_id_seq OWNED BY public.divisi.id;


--
-- TOC entry 234 (class 1259 OID 16661)
-- Name: galery; Type: TABLE; Schema: public; Owner: fismed-user
--

CREATE TABLE public.galery (
    id integer NOT NULL,
    nama character varying,
    iamge text
);


ALTER TABLE public.galery OWNER TO "fismed-user";

--
-- TOC entry 233 (class 1259 OID 16660)
-- Name: galery_id_seq; Type: SEQUENCE; Schema: public; Owner: fismed-user
--

CREATE SEQUENCE public.galery_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.galery_id_seq OWNER TO "fismed-user";

--
-- TOC entry 3539 (class 0 OID 0)
-- Dependencies: 233
-- Name: galery_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: fismed-user
--

ALTER SEQUENCE public.galery_id_seq OWNED BY public.galery.id;


--
-- TOC entry 232 (class 1259 OID 16652)
-- Name: hitang; Type: TABLE; Schema: public; Owner: fismed-user
--

CREATE TABLE public.hitang (
    id integer NOT NULL,
    nama character varying,
    nominal character varying,
    amount character varying,
    tanggal timestamp without time zone
);


ALTER TABLE public.hitang OWNER TO "fismed-user";

--
-- TOC entry 231 (class 1259 OID 16651)
-- Name: hitang_id_seq; Type: SEQUENCE; Schema: public; Owner: fismed-user
--

CREATE SEQUENCE public.hitang_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.hitang_id_seq OWNER TO "fismed-user";

--
-- TOC entry 3540 (class 0 OID 0)
-- Dependencies: 231
-- Name: hitang_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: fismed-user
--

ALTER SEQUENCE public.hitang_id_seq OWNED BY public.hitang.id;


--
-- TOC entry 238 (class 1259 OID 16698)
-- Name: item_buyer; Type: TABLE; Schema: public; Owner: fismed-user
--

CREATE TABLE public.item_buyer (
    id integer NOT NULL,
    po_id integer,
    name character varying,
    quantity integer,
    price character varying,
    discount character varying,
    amount character varying
);


ALTER TABLE public.item_buyer OWNER TO "fismed-user";

--
-- TOC entry 237 (class 1259 OID 16697)
-- Name: item_buyer_id_seq; Type: SEQUENCE; Schema: public; Owner: fismed-user
--

CREATE SEQUENCE public.item_buyer_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.item_buyer_id_seq OWNER TO "fismed-user";

--
-- TOC entry 3541 (class 0 OID 0)
-- Dependencies: 237
-- Name: item_buyer_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: fismed-user
--

ALTER SEQUENCE public.item_buyer_id_seq OWNED BY public.item_buyer.id;


--
-- TOC entry 213 (class 1259 OID 16458)
-- Name: order_items; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.order_items (
    id integer NOT NULL,
    pi_id integer,
    name character varying,
    quantity text,
    price character varying,
    discount character varying,
    sub_total character varying,
    kat character varying,
    created_at timestamp without time zone,
    created_by character varying,
    update_at timestamp without time zone,
    updated_by character varying
);


ALTER TABLE public.order_items OWNER TO postgres;

--
-- TOC entry 214 (class 1259 OID 16463)
-- Name: order_items_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.order_items_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.order_items_id_seq OWNER TO postgres;

--
-- TOC entry 3542 (class 0 OID 0)
-- Dependencies: 214
-- Name: order_items_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.order_items_id_seq OWNED BY public.order_items.id;


--
-- TOC entry 226 (class 1259 OID 16625)
-- Name: pemasukan; Type: TABLE; Schema: public; Owner: fismed-user
--

CREATE TABLE public.pemasukan (
    id integer NOT NULL,
    nama character varying,
    nominal character varying,
    amount character varying,
    tanggal timestamp without time zone
);


ALTER TABLE public.pemasukan OWNER TO "fismed-user";

--
-- TOC entry 225 (class 1259 OID 16624)
-- Name: pemasukan_id_seq; Type: SEQUENCE; Schema: public; Owner: fismed-user
--

CREATE SEQUENCE public.pemasukan_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.pemasukan_id_seq OWNER TO "fismed-user";

--
-- TOC entry 3543 (class 0 OID 0)
-- Dependencies: 225
-- Name: pemasukan_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: fismed-user
--

ALTER SEQUENCE public.pemasukan_id_seq OWNED BY public.pemasukan.id;


--
-- TOC entry 230 (class 1259 OID 16643)
-- Name: pengeluaran; Type: TABLE; Schema: public; Owner: fismed-user
--

CREATE TABLE public.pengeluaran (
    id integer NOT NULL,
    nama character varying,
    nominal character varying,
    amount character varying,
    tanggal timestamp without time zone
);


ALTER TABLE public.pengeluaran OWNER TO "fismed-user";

--
-- TOC entry 229 (class 1259 OID 16642)
-- Name: pengeluaran_id_seq; Type: SEQUENCE; Schema: public; Owner: fismed-user
--

CREATE SEQUENCE public.pengeluaran_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.pengeluaran_id_seq OWNER TO "fismed-user";

--
-- TOC entry 3544 (class 0 OID 0)
-- Dependencies: 229
-- Name: pengeluaran_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: fismed-user
--

ALTER SEQUENCE public.pengeluaran_id_seq OWNED BY public.pengeluaran.id;


--
-- TOC entry 215 (class 1259 OID 16464)
-- Name: performance_invoice; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.performance_invoice (
    id integer NOT NULL,
    customer_id integer,
    sub_total character varying,
    status character varying,
    divisi character varying,
    invoice_number character varying,
    po_number character varying,
    due_date character varying,
    doctor_name character varying,
    patient_name character varying,
    created_at timestamp without time zone,
    created_by character varying,
    update_at timestamp without time zone,
    updated_by character varying,
    total character varying,
    pajak character varying,
    tanggal_tindakan character varying,
    rm character varying,
    number_si character varying,
    reason text
);


ALTER TABLE public.performance_invoice OWNER TO postgres;

--
-- TOC entry 216 (class 1259 OID 16469)
-- Name: performance_invoice_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.performance_invoice_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.performance_invoice_id_seq OWNER TO postgres;

--
-- TOC entry 3545 (class 0 OID 0)
-- Dependencies: 216
-- Name: performance_invoice_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.performance_invoice_id_seq OWNED BY public.performance_invoice.id;


--
-- TOC entry 228 (class 1259 OID 16634)
-- Name: piutang; Type: TABLE; Schema: public; Owner: fismed-user
--

CREATE TABLE public.piutang (
    id integer NOT NULL,
    nama character varying,
    nominal character varying,
    amount character varying,
    tanggal timestamp without time zone
);


ALTER TABLE public.piutang OWNER TO "fismed-user";

--
-- TOC entry 227 (class 1259 OID 16633)
-- Name: piutang_id_seq; Type: SEQUENCE; Schema: public; Owner: fismed-user
--

CREATE SEQUENCE public.piutang_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.piutang_id_seq OWNER TO "fismed-user";

--
-- TOC entry 3546 (class 0 OID 0)
-- Dependencies: 227
-- Name: piutang_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: fismed-user
--

ALTER SEQUENCE public.piutang_id_seq OWNED BY public.piutang.id;


--
-- TOC entry 236 (class 1259 OID 16689)
-- Name: purchase_order; Type: TABLE; Schema: public; Owner: fismed-user
--

CREATE TABLE public.purchase_order (
    id integer NOT NULL,
    nama_suplier character varying,
    nomor_po character varying,
    tanggal timestamp without time zone,
    catatan_po text,
    prepared_by character varying,
    prepared_jabatan character varying,
    approved_by character varying,
    approved_jabatan character varying,
    divisi character varying,
    created_at timestamp without time zone,
    created_by character varying,
    updated_at timestamp without time zone,
    updated_by character varying,
    sub_total character varying,
    pajak character varying,
    total character varying
);


ALTER TABLE public.purchase_order OWNER TO "fismed-user";

--
-- TOC entry 235 (class 1259 OID 16688)
-- Name: purchase_order_id_seq; Type: SEQUENCE; Schema: public; Owner: fismed-user
--

CREATE SEQUENCE public.purchase_order_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.purchase_order_id_seq OWNER TO "fismed-user";

--
-- TOC entry 3547 (class 0 OID 0)
-- Dependencies: 235
-- Name: purchase_order_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: fismed-user
--

ALTER SEQUENCE public.purchase_order_id_seq OWNED BY public.purchase_order.id;


--
-- TOC entry 217 (class 1259 OID 16470)
-- Name: stock_items; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.stock_items (
    id integer NOT NULL,
    name character varying NOT NULL,
    total text,
    price character varying,
    created_at timestamp without time zone,
    created_by character varying,
    updated_at timestamp without time zone,
    updated_by character varying
);


ALTER TABLE public.stock_items OWNER TO postgres;

--
-- TOC entry 218 (class 1259 OID 16475)
-- Name: stock_items_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.stock_items_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.stock_items_id_seq OWNER TO postgres;

--
-- TOC entry 3548 (class 0 OID 0)
-- Dependencies: 218
-- Name: stock_items_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.stock_items_id_seq OWNED BY public.stock_items.id;


--
-- TOC entry 219 (class 1259 OID 16476)
-- Name: tax_code; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.tax_code (
    id integer NOT NULL,
    tax character varying(50),
    created_at timestamp without time zone,
    created_by character varying,
    updated_at timestamp without time zone,
    updated_by character varying
);


ALTER TABLE public.tax_code OWNER TO postgres;

--
-- TOC entry 220 (class 1259 OID 16481)
-- Name: tax_code_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.tax_code_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.tax_code_id_seq OWNER TO postgres;

--
-- TOC entry 3549 (class 0 OID 0)
-- Dependencies: 220
-- Name: tax_code_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.tax_code_id_seq OWNED BY public.tax_code.id;


--
-- TOC entry 221 (class 1259 OID 16482)
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
-- TOC entry 222 (class 1259 OID 16489)
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
-- TOC entry 3550 (class 0 OID 0)
-- Dependencies: 222
-- Name: user_category_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.user_category_id_seq OWNED BY public.user_category.id;


--
-- TOC entry 223 (class 1259 OID 16490)
-- Name: users; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.users (
    id integer NOT NULL,
    username character varying(50) NOT NULL,
    password character varying(255),
    role_id integer,
    token text,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    created_by character varying(255),
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    updated_by character varying(255)
);


ALTER TABLE public.users OWNER TO postgres;

--
-- TOC entry 224 (class 1259 OID 16497)
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
-- TOC entry 3551 (class 0 OID 0)
-- Dependencies: 224
-- Name: user_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.user_id_seq OWNED BY public.users.id;


--
-- TOC entry 3303 (class 2604 OID 16498)
-- Name: customer id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.customer ALTER COLUMN id SET DEFAULT nextval('public.customer_id_seq'::regclass);


--
-- TOC entry 3304 (class 2604 OID 16499)
-- Name: divisi id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.divisi ALTER COLUMN id SET DEFAULT nextval('public.divisi_id_seq'::regclass);


--
-- TOC entry 3321 (class 2604 OID 16664)
-- Name: galery id; Type: DEFAULT; Schema: public; Owner: fismed-user
--

ALTER TABLE ONLY public.galery ALTER COLUMN id SET DEFAULT nextval('public.galery_id_seq'::regclass);


--
-- TOC entry 3320 (class 2604 OID 16655)
-- Name: hitang id; Type: DEFAULT; Schema: public; Owner: fismed-user
--

ALTER TABLE ONLY public.hitang ALTER COLUMN id SET DEFAULT nextval('public.hitang_id_seq'::regclass);


--
-- TOC entry 3323 (class 2604 OID 16701)
-- Name: item_buyer id; Type: DEFAULT; Schema: public; Owner: fismed-user
--

ALTER TABLE ONLY public.item_buyer ALTER COLUMN id SET DEFAULT nextval('public.item_buyer_id_seq'::regclass);


--
-- TOC entry 3307 (class 2604 OID 16500)
-- Name: order_items id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.order_items ALTER COLUMN id SET DEFAULT nextval('public.order_items_id_seq'::regclass);


--
-- TOC entry 3317 (class 2604 OID 16628)
-- Name: pemasukan id; Type: DEFAULT; Schema: public; Owner: fismed-user
--

ALTER TABLE ONLY public.pemasukan ALTER COLUMN id SET DEFAULT nextval('public.pemasukan_id_seq'::regclass);


--
-- TOC entry 3319 (class 2604 OID 16646)
-- Name: pengeluaran id; Type: DEFAULT; Schema: public; Owner: fismed-user
--

ALTER TABLE ONLY public.pengeluaran ALTER COLUMN id SET DEFAULT nextval('public.pengeluaran_id_seq'::regclass);


--
-- TOC entry 3308 (class 2604 OID 16501)
-- Name: performance_invoice id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.performance_invoice ALTER COLUMN id SET DEFAULT nextval('public.performance_invoice_id_seq'::regclass);


--
-- TOC entry 3318 (class 2604 OID 16637)
-- Name: piutang id; Type: DEFAULT; Schema: public; Owner: fismed-user
--

ALTER TABLE ONLY public.piutang ALTER COLUMN id SET DEFAULT nextval('public.piutang_id_seq'::regclass);


--
-- TOC entry 3322 (class 2604 OID 16692)
-- Name: purchase_order id; Type: DEFAULT; Schema: public; Owner: fismed-user
--

ALTER TABLE ONLY public.purchase_order ALTER COLUMN id SET DEFAULT nextval('public.purchase_order_id_seq'::regclass);


--
-- TOC entry 3309 (class 2604 OID 16502)
-- Name: stock_items id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.stock_items ALTER COLUMN id SET DEFAULT nextval('public.stock_items_id_seq'::regclass);


--
-- TOC entry 3310 (class 2604 OID 16503)
-- Name: tax_code id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.tax_code ALTER COLUMN id SET DEFAULT nextval('public.tax_code_id_seq'::regclass);


--
-- TOC entry 3311 (class 2604 OID 16504)
-- Name: user_category id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.user_category ALTER COLUMN id SET DEFAULT nextval('public.user_category_id_seq'::regclass);


--
-- TOC entry 3314 (class 2604 OID 16505)
-- Name: users id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.user_id_seq'::regclass);


--
-- TOC entry 3500 (class 0 OID 16443)
-- Dependencies: 209
-- Data for Name: customer; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.customer VALUES (2, 'OETOMO', '456 Market St', '456 NPWP St', '0987654321', 'IPAK456', '456 Facture St', 'City B', '67890', '234-567-8901', 'emailB@company.com', '234-567-8902', 'PIC B', '456 Item St', 'City B', '67890', '234-567-8903', 'itemB@company.com', '234-567-8904', 'PIC Item B', 'Contact B', 2, 'Gak tau ini apa ', '2024-06-12 22:23:31.420925', 'admin', '2024-06-12 22:23:31.420925', 'admin', NULL);
INSERT INTO public.customer VALUES (3, 'BOROMEUS', '789 Broadway', '789 NPWP St', '1122334455', 'IPAK789', '789 Facture St', 'City C', '54321', '345-678-9012', 'emailC@company.com', '345-678-9013', 'PIC C', '789 Item St', 'City C', '54321', '345-678-9014', 'itemC@company.com', '345-678-9015', 'PIC Item C', 'Contact C', 3, 'Gak tau ini apa ', '2024-06-12 22:23:31.420925', 'admin', '2024-06-12 22:23:31.420925', 'admin', NULL);
INSERT INTO public.customer VALUES (15, 'Company A', '123 Main St', '123 NPWP St', '1234567890', 'IPAK123', '123 Facture St', 'City A', '12345', '123-456-7890', 'emailA@company.com', '123-456-7891', 'PIC A', '123 Item St', 'City A', '12345', '123-456-7892', 'itemA@company.com', '123-456-7893', 'PIC Item A', 'Contact A', 1, 'Gak tau ini apa ', '2024-06-22 13:37:57.063848', 'admin', '2024-06-22 13:37:57.063848', 'admin', NULL);
INSERT INTO public.customer VALUES (5, 'SANTO YUSUF', '202 Oak St', '202 NPWP St', '3344556677', 'IPAK202', '202 Facture St', 'City E', '87654', '567-890-1234', 'emailE@company.com', '567-890-1235', 'PIC E', '202 Item St', 'City E', '87654', '567-890-1236', 'itemE@company.com', '567-890-1237', 'PIC Item E', 'Contact E', 2, 'Gak tau ini apa ', '2024-06-12 22:23:31.420925', 'admin', '2024-06-12 22:23:31.420925', 'admin', NULL);
INSERT INTO public.customer VALUES (4, 'KARTINI', '101 State St', '101 NPWP St', '2233445566', 'IPAK101', '101 Facture St', 'City D', '98765', '456-789-0123', 'emailD@company.com', '456-789-0124', 'PIC D', '101 Item St', 'City D', '98765', '456-789-0125', 'itemD@company.com', '456-789-0126', 'PIC Item D', 'Contact D', 1, 'Gak tau ini apa ', '2024-06-12 22:23:31.420925', 'admin', '2024-06-12 22:23:31.420925', 'admin', NULL);
INSERT INTO public.customer VALUES (1, 'MAYAPADA', 'Bandung', '123 NPWP St', '1234567890', 'IPAK123', '123 Facture St', 'City A', '12345', '123-456-7890', 'emailA@company.com', '123-456-7891', 'PIC A', '123 Item St', 'City A', '12345', '123-456-7892', 'itemA@company.com', '123-456-7893', 'PIC Item A', 'Contact A', 1, 'Gak tau ini apa ', '2024-06-12 22:23:31.420925', 'admin', '2024-06-12 22:23:31.420925', 'admin', NULL);
INSERT INTO public.customer VALUES (16, 'Company A', '123 Main St', '123 NPWP St', '1234567890', 'IPAK123', '123 Facture St', 'City A', '12345', '123-456-7890', 'emailA@company.com', '123-456-7891', 'PIC A', '123 Item St', 'City A', '12345', '123-456-7892', 'itemA@company.com', '123-456-7893', 'PIC Item A', 'Contact A', 1, '30 days', '2024-06-24 16:04:05.916188', 'admin', '2024-06-24 16:04:05.916188', 'admin', NULL);
INSERT INTO public.customer VALUES (17, 'Company A', '123 Main St', '123 NPWP St', '1234567890', 'IPAK123', '123 Facture St', 'City A', '12345', '123-456-7890', 'emailA@company.com', '123-456-7891', 'PIC A', '123 Item St', 'City A', '12345', '123-456-7892', 'itemA@company.com', '123-456-7893', 'PIC Item A', 'Contact A', 1, '30 days', '2024-06-25 14:20:41.388399', 'admin', '2024-06-25 14:20:41.388399', 'admin', '09120310239');
INSERT INTO public.customer VALUES (18, 'Company A', '123 Main St', '123 NPWP St', '1234567890', 'IPAK123', '123 Facture St', 'City A', '12345', '123-456-7890', 'emailA@company.com', '123-456-7891', 'PIC A', '123 Item St', 'City A', '12345', '123-456-7892', 'itemA@company.com', '123-456-7893', 'PIC Item A', 'Contact A', 1, '30 days', '2024-06-25 14:24:38.703729', 'admin', '2024-06-25 14:24:38.703729', 'admin', '09120310239');
INSERT INTO public.customer VALUES (19, 'Company A', '123 Main St', '123 NPWP St', '1234567890', 'IPAK123', '123 Facture St', 'City A', '12345', '123-456-7890', 'emailA@company.com', '123-456-7891', 'PIC A', '123 Item St', 'City A', '12345', '123-456-7892', 'itemA@company.com', '123-456-7893', 'PIC Item A', 'Contact A', 1, '30 days', '2024-07-03 14:24:34.335348', 'admin', '2024-07-03 14:24:34.335348', 'admin', '09120310239');
INSERT INTO public.customer VALUES (20, 'Company A', '123 Main St', '123 NPWP St', '1234567890', 'IPAK123', '123 Facture St', 'City A', '12345', '123-456-7890', 'emailA@company.com', '123-456-7891', 'PIC A', '123 Item St', 'City A', '12345', '123-456-7892', 'itemA@company.com', '123-456-7893', 'PIC Item A', 'Contact A', 1, '30 days', '2024-07-06 15:38:28.911169', 'admin', '2024-07-06 15:38:28.911169', 'admin', '09120310239');
INSERT INTO public.customer VALUES (21, 'asdasdad', 'asda', 'dada', 'as', 'asdasd', 'asdas', 'asda', 'dasda', 'dasda', '', 'dad', 'dasd', 'dasd', 'asd', 'asdasd', 'asdas', '', 'dasdas', 'sdasd', 'asda', 2, 'dasd', '2024-07-10 15:28:56.986127', 'admin', '2024-07-10 15:28:56.986127', 'admin', 'asdasdasd');


--
-- TOC entry 3502 (class 0 OID 16450)
-- Dependencies: 211
-- Data for Name: divisi; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.divisi VALUES (1, 'RADIOLOGI', '2024-06-21 10:12:55.531313', 'admin', '2024-06-21 10:12:55.531313', 'admin');
INSERT INTO public.divisi VALUES (2, 'ORTOPEDI', '2024-06-21 10:12:55.531313', 'admin', '2024-06-21 10:12:55.531313', 'admin');
INSERT INTO public.divisi VALUES (3, 'KARDIOLOGI', '2024-06-21 10:12:55.531313', 'admin', '2024-06-21 10:12:55.531313', 'admin');


--
-- TOC entry 3525 (class 0 OID 16661)
-- Dependencies: 234
-- Data for Name: galery; Type: TABLE DATA; Schema: public; Owner: fismed-user
--

INSERT INTO public.galery VALUES (1, 'logo', '/9j/4AAQSkZJRgABAQAAAQABAAD//gA7Q1JFQVRPUjogZ2QtanBlZyB2MS4wICh1c2luZyBJSkcgSlBFRyB2NjIpLCBxdWFsaXR5ID0gODIK/9sAQwAGBAQFBAQGBQUFBgYGBwkOCQkICAkSDQ0KDhUSFhYVEhQUFxohHBcYHxkUFB0nHR8iIyUlJRYcKSwoJCshJCUk/9sAQwEGBgYJCAkRCQkRJBgUGCQkJCQkJCQkJCQkJCQkJCQkJCQkJCQkJCQkJCQkJCQkJCQkJCQkJCQkJCQkJCQkJCQk/8AAEQgAMgAyAwEiAAIRAQMRAf/EAB8AAAEFAQEBAQEBAAAAAAAAAAABAgMEBQYHCAkKC//EALUQAAIBAwMCBAMFBQQEAAABfQECAwAEEQUSITFBBhNRYQcicRQygZGhCCNCscEVUtHwJDNicoIJChYXGBkaJSYnKCkqNDU2Nzg5OkNERUZHSElKU1RVVldYWVpjZGVmZ2hpanN0dXZ3eHl6g4SFhoeIiYqSk5SVlpeYmZqio6Slpqeoqaqys7S1tre4ubrCw8TFxsfIycrS09TV1tfY2drh4uPk5ebn6Onq8fLz9PX29/j5+v/EAB8BAAMBAQEBAQEBAQEAAAAAAAABAgMEBQYHCAkKC//EALURAAIBAgQEAwQHBQQEAAECdwABAgMRBAUhMQYSQVEHYXETIjKBCBRCkaGxwQkjM1LwFWJy0QoWJDThJfEXGBkaJicoKSo1Njc4OTpDREVGR0hJSlNUVVZXWFlaY2RlZmdoaWpzdHV2d3h5eoKDhIWGh4iJipKTlJWWl5iZmqKjpKWmp6ipqrKztLW2t7i5usLDxMXGx8jJytLT1NXW19jZ2uLj5OXm5+jp6vLz9PX29/j5+v/aAAwDAQACEQMRAD8A+qKDwKR3VFLMwUDkknGKrHVLHH/H3b/9/BTs2tERKcY7s8q8YeNr/UtQmtbK4ktrOFyg8pirS44JJ649qyNI8U6to1ys0F3NIoOWhlcsj+xB/nWbeOpu5zuX/Wv3H941GCG4X5j6Dmvia9XE+1dR3Wp4ksQ5Tumew+F9YTxJrN3qcaSJDDbxQJu6bjln/wDZR+FdXXker+HvFejaFpSaJLNCWLzXflzLGVkbG3cSR8oAx+ea9D0XX7e/toI47hb64UCOeS2UtGsgA3fN0Az719nQoyjh4VJtNy3XY78Pi+abpTVmvx9DZoooqjvPDfip4vvNS1q40eC4ZLC1IjdEOBK+Mnd6gZxjpxmuACLkfKv5V2PxP8PXGjeJrm6MbfZL1/Njl/h3EfMufXIJx6EVy1pZT30yxQRs7McDAzk195g54ajhIzukran5Pmf1qrjZRldu+noZ6wtLN5cabmJIAArtdBfT/B0MepX0YuLr/WW9qOPMbtIx/hQdu5PIFST+HYvBWmLfaoiPfXOfs9m/U+rSeij+6OSSOa4q9vZr+5kuJ5DJJI25mPVjXh+w/tKXPNWpJ6f3n/kenzywWn2/y/4J2OpeINR8VbbnU7ppomO5Ldflij5OPl7n3PNWdD1+98P3aXFpKwUEeZFn5ZF9CP69qxNLH/Evg/3f6mtLT9OudVvI7O0jaSaQ4AA6D1PoB61+a5lia8sbKEW9HZI97CXajNfEz3y2v4Lq3injbKSoHU+oIyKKistMis7OC2UErDGsYJ9AMUV7ydbsfR3mN1XSItYga1ujutnGHj2qd34kHH4VUsPD+ieF7Z5rW0ht1ijJeU8sFAycsea2q5zx+858L3Nrasq3F60dpHu6ZkcIf0JruhHmkot6GVeMYRdVL3kjwnxXr83ifW7m/mJKOSkKn+CIH5R/X6msfYn91fyrovGPg298JX5SVHks3OYbjHDD0PoR6Vk6Vpd5rV7HZafA1xPIcBV6D3J7D3r9Ew9TDqgpQa5Uj8nxNLFPEyjNPmbPZvhz4Y0bUPB2nXF1pttNK6tudkBJ+c12WnaNp+khxY2cFvv5by0Az9ar+F9G/wCEf0Cy03fva3iCs3q3U/qTWrX53XjTlWlUit2z9WwVH2dCEZLVJBRRRUnWFY+vKrXOkBgCPtoOCPSNyKKKqG5lW+A05Y0ljKSIrqRghhkGo7O0t7VStvbxQgnJEaBf5UUVtH+Gzln/ABkWaKKK5zvCiiigD//Z');


--
-- TOC entry 3523 (class 0 OID 16652)
-- Dependencies: 232
-- Data for Name: hitang; Type: TABLE DATA; Schema: public; Owner: fismed-user
--



--
-- TOC entry 3529 (class 0 OID 16698)
-- Dependencies: 238
-- Data for Name: item_buyer; Type: TABLE DATA; Schema: public; Owner: fismed-user
--



--
-- TOC entry 3504 (class 0 OID 16458)
-- Dependencies: 213
-- Data for Name: order_items; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.order_items VALUES (2, 1, 'Item A2', '3', '100', '10', '50000000', 'Category1', '2024-06-01 10:00:00', 'admin', '2024-06-01 10:00:00', 'admin');
INSERT INTO public.order_items VALUES (3, 2, 'Item A1', '2', '50', '0', '10000000', 'Category1', '2024-06-01 10:00:00', 'admin', '2024-06-01 10:00:00', 'admin');
INSERT INTO public.order_items VALUES (4, 2, 'Item A2', '2', '50', '20', '20000000', 'Category1', '2024-06-01 10:00:00', 'admin', '2024-06-01 10:00:00', 'admin');
INSERT INTO public.order_items VALUES (5, 2, 'Item A3', '2', '50', '10', '30000000', 'Category1', '2024-06-01 10:00:00', 'admin', '2024-06-01 10:00:00', 'admin');
INSERT INTO public.order_items VALUES (6, 7, 'Alat Medis A', '10', '500000', '5', 'Rp. 4.750.000', 'K001', '2024-06-22 12:50:26.278799', 'sales', '2024-06-22 12:50:26.278799', 'sales');
INSERT INTO public.order_items VALUES (7, 7, 'Alat Medis B', '20', '300000', '10', 'Rp. 5.400.000', 'K002', '2024-06-22 12:50:26.278799', 'sales', '2024-06-22 12:50:26.278799', 'sales');
INSERT INTO public.order_items VALUES (8, 8, 'Alat Medis A', '10', '500000', '5', 'Rp. 4.750.000', 'K001', '2024-06-22 12:57:33.828504', 'sales', '2024-06-22 12:57:33.828504', 'sales');
INSERT INTO public.order_items VALUES (9, 8, 'Alat Medis B', '20', '300000', '10', 'Rp. 5.400.000', 'K002', '2024-06-22 12:57:33.828504', 'sales', '2024-06-22 12:57:33.828504', 'sales');
INSERT INTO public.order_items VALUES (11, 9, 'Alat Medis B', '20', '300000', '10', 'Rp. 5.400.000', 'K002', '2024-06-22 13:06:10.582396', 'sales', '2024-06-22 13:06:10.582396', 'sales');
INSERT INTO public.order_items VALUES (12, 10, 'Alat Medis A', '10', '500000', '5', 'Rp. 4.750.000', 'K001', '2024-06-22 13:21:49.952658', 'sales', '2024-06-22 13:21:49.952658', 'sales');
INSERT INTO public.order_items VALUES (14, 13, 'Alat Medis A', '10', '500000', '5', 'Rp. 4.750.000', 'K001', '2024-06-23 00:36:20.235691', 'sales', '2024-06-23 00:36:20.235691', 'sales');
INSERT INTO public.order_items VALUES (15, 13, 'Alat Medis B', '20', '300000', '10', 'Rp. 5.400.000', 'K002', '2024-06-23 00:36:20.235691', 'sales', '2024-06-23 00:36:20.235691', 'sales');
INSERT INTO public.order_items VALUES (16, 14, 'Alat Medis A', '10', '500000', '5', 'Rp. 4.750.000', 'K001', '2024-06-23 00:36:51.305974', 'sales', '2024-06-23 00:36:51.305974', 'sales');
INSERT INTO public.order_items VALUES (17, 14, 'Alat Medis B', '20', '300000', '10', 'Rp. 5.400.000', 'K002', '2024-06-23 00:36:51.305974', 'sales', '2024-06-23 00:36:51.305974', 'sales');
INSERT INTO public.order_items VALUES (52, 31, '', '30', '4000', '20', '', '123123', '2024-07-11 12:08:17.638114', 'sales', '2024-07-11 12:08:17.638114', 'sales');
INSERT INTO public.order_items VALUES (53, 32, '', '20000', '5000', '10', '', '123123', '2024-07-11 12:15:27.70179', 'sales', '2024-07-11 12:15:27.70179', 'sales');
INSERT INTO public.order_items VALUES (18, 10, 'Alat Medis B', '20', '300000', '10', 'Rp. 5.400.000', 'K003', '2024-06-26 04:22:45.508295', 'sales', '2024-06-26 04:22:45.508295', 'sales');
INSERT INTO public.order_items VALUES (19, 10, 'Alat Medis B', '20', '300000', '10', 'Rp. 5.400.000', 'K003', '2024-06-26 04:25:16.398501', 'sales', '2024-06-26 04:25:16.398501', 'sales');
INSERT INTO public.order_items VALUES (20, 1, 'Alat Medis B', '20', '300000', '10', 'Rp. 5.400.000', 'K003', '2024-06-26 04:28:28.244233', 'sales', '2024-06-26 04:28:28.244233', 'sales');
INSERT INTO public.order_items VALUES (10, 9, 'Alat Medis B', '50', '310000', '10', 'Rp. 13.950.000', 'K002', '2024-06-22 13:06:10.582396', 'sales', '2024-06-26 11:29:28.49443', 'sales');
INSERT INTO public.order_items VALUES (21, 10, 'Alat Medis B', '20', '300000', '10', 'Rp. 5.400.000', 'K003', '2024-06-26 04:29:27.74466', 'sales', '2024-06-26 04:29:27.74466', 'sales');
INSERT INTO public.order_items VALUES (22, 1, 'Alat Medis B', '20', '300000', '10', 'Rp. 5.400.000', 'K003', '2024-06-26 04:31:26.608391', 'sales', '2024-06-26 04:31:26.608391', 'sales');
INSERT INTO public.order_items VALUES (1, 1, 'Alat Medis B', '50', '310000', '10', 'Rp. 13.950.000', 'K002', '2024-06-01 10:00:00', 'admin', '2024-06-26 11:32:46.037018', 'sales');
INSERT INTO public.order_items VALUES (23, 1, 'Alat Medis B', '20', '300000', '10', 'Rp. 5.400.000', 'K003', '2024-06-26 04:32:45.307122', 'sales', '2024-06-26 04:32:45.307122', 'sales');
INSERT INTO public.order_items VALUES (13, 10, 'Alat Medis B', '20123', '30123000', '10', 'Rp. 545.548.616.100', 'K002', '2024-06-22 13:21:49.952658', 'sales', '2024-06-26 11:36:25.623949', 'sales');
INSERT INTO public.order_items VALUES (24, 15, 'Alat Medis A', '10', '500000', '5', 'Rp. 4.750.000', 'K001', '2024-07-02 14:30:00.70458', 'sales', '2024-07-02 14:30:00.70458', 'sales');
INSERT INTO public.order_items VALUES (25, 15, 'Alat Medis B', '20', '300000', '10', 'Rp. 5.400.000', 'K002', '2024-07-02 14:30:00.70458', 'sales', '2024-07-02 14:30:00.70458', 'sales');
INSERT INTO public.order_items VALUES (26, 16, 'Alat Medis A', '10', '500000', '5', 'Rp. 4.750.000', 'K001', '2024-07-02 14:30:45.681076', 'sales', '2024-07-02 14:30:45.681076', 'sales');
INSERT INTO public.order_items VALUES (27, 16, 'Alat Medis B', '20', '300000', '10', 'Rp. 5.400.000', 'K002', '2024-07-02 14:30:45.681076', 'sales', '2024-07-02 14:30:45.681076', 'sales');
INSERT INTO public.order_items VALUES (28, 17, 'Alat Medis A', '10', '500000', '5', 'Rp. 4.750.000', 'K001', '2024-07-03 14:21:38.082248', 'sales', '2024-07-03 14:21:38.082248', 'sales');
INSERT INTO public.order_items VALUES (29, 17, 'Alat Medis B', '20', '300000', '10', 'Rp. 5.400.000', 'K002', '2024-07-03 14:21:38.082248', 'sales', '2024-07-03 14:21:38.082248', 'sales');
INSERT INTO public.order_items VALUES (30, 18, 'Alat Medis A', '10', '500000', '5', 'Rp. 4.750.000', 'K001', '2024-07-06 07:24:40.554983', 'sales', '2024-07-06 07:24:40.554983', 'sales');
INSERT INTO public.order_items VALUES (31, 18, 'Alat Medis B', '20', '300000', '10', 'Rp. 5.400.000', 'K002', '2024-07-06 07:24:40.554983', 'sales', '2024-07-06 07:24:40.554983', 'sales');
INSERT INTO public.order_items VALUES (32, 19, 'Alat Medis A', '10', '500000', '5', 'Rp. 4.750.000', 'K001', '2024-07-06 07:25:03.063931', 'sales', '2024-07-06 07:25:03.063931', 'sales');
INSERT INTO public.order_items VALUES (33, 19, 'Alat Medis B', '20', '300000', '10', 'Rp. 5.400.000', 'K002', '2024-07-06 07:25:03.063931', 'sales', '2024-07-06 07:25:03.063931', 'sales');
INSERT INTO public.order_items VALUES (34, 20, 'gyuikjb', '10', '300000', '5', '', 'okad', '2024-07-06 12:00:32.332196', 'sales', '2024-07-06 12:00:32.332196', 'sales');
INSERT INTO public.order_items VALUES (35, 21, 'ADITYA COBA FRONTEND', '23', '2000000', '5', '', '123', '2024-07-06 14:55:07.071272', 'sales', '2024-07-06 14:55:07.071272', 'sales');
INSERT INTO public.order_items VALUES (36, 22, 'Suntik', '1', '2000', '20', '', '1', '2024-07-08 10:08:48.342944', 'sales', '2024-07-08 10:08:48.342944', 'sales');
INSERT INTO public.order_items VALUES (43, 26, 'Alat Medis A', '10', '500000', '5', 'Rp. 4.750.000', 'K001', '2024-07-09 03:37:29.564957', 'sales', '2024-07-09 03:37:29.564957', 'sales');
INSERT INTO public.order_items VALUES (44, 26, 'Alat Medis B', '20', '300000', '10', 'Rp. 5.400.000', 'K002', '2024-07-09 03:37:29.564957', 'sales', '2024-07-09 03:37:29.564957', 'sales');
INSERT INTO public.order_items VALUES (45, 27, 'Item C', '10', '500000', '5', 'Rp. 4.750.000', 'K001', '2024-07-09 03:38:43.191594', 'sales', '2024-07-09 03:38:43.191594', 'sales');
INSERT INTO public.order_items VALUES (46, 27, 'Item B', '20', '300000', '10', 'Rp. 5.400.000', 'K002', '2024-07-09 03:38:43.191594', 'sales', '2024-07-09 03:38:43.191594', 'sales');
INSERT INTO public.order_items VALUES (47, 28, 'Alat Medis A', '10', '500000', '5', 'Rp. 4.750.000', 'K001', '2024-07-09 04:58:00.201573', 'sales', '2024-07-09 04:58:00.201573', 'sales');
INSERT INTO public.order_items VALUES (48, 28, 'Alat Medis B', '20', '300000', '10', 'Rp. 5.400.000', 'K002', '2024-07-09 04:58:00.201573', 'sales', '2024-07-09 04:58:00.201573', 'sales');
INSERT INTO public.order_items VALUES (49, 29, 'Alat Medis A', '10', '500000', '5', 'Rp. 4.750.000', 'K001', '2024-07-09 04:58:02.921971', 'sales', '2024-07-09 04:58:02.921971', 'sales');
INSERT INTO public.order_items VALUES (50, 29, 'Alat Medis B', '20', '300000', '10', 'Rp. 5.400.000', 'K002', '2024-07-09 04:58:02.921971', 'sales', '2024-07-09 04:58:02.921971', 'sales');
INSERT INTO public.order_items VALUES (51, 30, '', '10', '17293', '5', '', 'K002', '2024-07-10 12:42:11.099732', 'sales', '2024-07-10 12:42:11.099732', 'sales');


--
-- TOC entry 3517 (class 0 OID 16625)
-- Dependencies: 226
-- Data for Name: pemasukan; Type: TABLE DATA; Schema: public; Owner: fismed-user
--

INSERT INTO public.pemasukan VALUES (1, 'MAYAPADA HOSPITAL', '11266500', 'Rp. 11.266.500', '2024-07-09 03:37:29.564957');
INSERT INTO public.pemasukan VALUES (2, 'MAYAPADA HOSPITAL', '11266500', 'Rp. 11.266.500', '2024-07-09 03:38:43.191594');
INSERT INTO public.pemasukan VALUES (3, 'MAYAPADA HOSPITAL', '11266500', 'Rp. 11.266.500', '2024-07-09 04:58:00.201573');
INSERT INTO public.pemasukan VALUES (4, 'MAYAPADA HOSPITAL', '11266500', 'Rp. 11.266.500', '2024-07-09 04:58:02.921971');
INSERT INTO public.pemasukan VALUES (5, 'MAYAPADA', '106560', 'Rp. 106.560', '2024-07-11 12:08:17.638114');
INSERT INTO public.pemasukan VALUES (6, 'MAYAPADA', '99900000', 'Rp. 99.900.000', '2024-07-11 12:15:27.70179');


--
-- TOC entry 3521 (class 0 OID 16643)
-- Dependencies: 230
-- Data for Name: pengeluaran; Type: TABLE DATA; Schema: public; Owner: fismed-user
--



--
-- TOC entry 3506 (class 0 OID 16464)
-- Dependencies: 215
-- Data for Name: performance_invoice; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.performance_invoice VALUES (2, 2, 'Rp. 60,000,000', 'Diterima', 'Radiologi', '1002', '2002', '2024-08-01', 'Dr. Jones', 'Jane Doe', '2024-06-02 10:00:00', 'Sales', '2024-06-02 10:00:00', 'admin', 'Rp. 66,600,000', 'Rp. 6,600,000', NULL, NULL, NULL, NULL);
INSERT INTO public.performance_invoice VALUES (9, 1, 'Rp. 10.150.000', 'Diproses', 'Ortopedi', 'PI/581/7263Z4X2Z3', 'PO/581/7263C4A2C3', '2024-07-01', NULL, NULL, '2024-06-22 13:06:10.582396', 'sales', '2024-06-22 13:06:10.582396', 'sales', 'Rp. 11.266.500', 'Rp. 1.116.500', NULL, NULL, 'SI/581/63X79X5915', NULL);
INSERT INTO public.performance_invoice VALUES (7, 1, 'Rp. 10.150.000', 'Diproses', 'Ortopedi', 'PI/581/7263Z4X2Z3', 'PO/581/7263C4A2C3', '2024-07-01', NULL, NULL, '2024-06-22 12:50:26.278799', 'sales', '2024-06-22 12:50:26.278799', 'sales', 'Rp. 11.266.500', 'Rp. 1.116.500', NULL, NULL, 'SI/581/63X79X5915', NULL);
INSERT INTO public.performance_invoice VALUES (8, 1, 'Rp. 10.150.000', 'Diproses', 'Ortopedi', 'PI/581/7263Z4X2Z3', 'PO/581/7263C4A2C3', '2024-07-01', NULL, NULL, '2024-06-22 12:57:33.828504', 'sales', '2024-06-22 12:57:33.828504', 'sales', 'Rp. 11.266.500', 'Rp. 1.116.500', NULL, NULL, 'SI/581/63X79X5915', NULL);
INSERT INTO public.performance_invoice VALUES (11, 1, 'Rp. 60,000,000', 'Ditolak', 'Ortopedi', '1001', '2001', '2024-07-02', '', '', '2024-06-02 10:00:00', 'Sales', '2024-06-01 10:00:00', 'admin', 'Rp. 66,600,000', 'Rp. 6,600,000', NULL, NULL, 'SI/581/63X79X5915', NULL);
INSERT INTO public.performance_invoice VALUES (14, 1, 'Rp. 10.150.000', 'Diproses', 'Ortopedi', 'PI/525/Y46509Z603', 'PO/525/B46509C603', '2024-07-02', 'DR. AHMAD EFENDI', 'REHAN LATUCONSINA', '2024-06-23 00:36:51.305974', 'sales', '2024-06-23 00:36:51.305974', 'sales', 'Rp. 11.266.500', 'Rp. 1.116.500', '2024-08-12', '1092882', 'SI/525/Y46509Z603', NULL);
INSERT INTO public.performance_invoice VALUES (1, 1, 'Rp. 24.100.000', 'Diterima', 'Radiologi', 'PI/525/Y46509Z603100', 'PO/525/B46509C603100', '2024-07-02', 'DR. AHMAD EFENDI100', 'REHAN LATUCONSINA100', '2024-06-02 10:00:00', 'Sales', '2024-06-27 07:02:35.2185', 'admin', 'Rp. 26.751.000', 'Rp. 2.651.000', '2024-08-12100', '1092882100', 'SI/525/Y46509Z603100', '1. Ditolak, 2. Ditolak, 3. Ditolak');
INSERT INTO public.performance_invoice VALUES (12, 1, 'Rp. 60,000,000', 'Ditolak', 'Ortopedi', '1001', '2001', '2024-07-02', '', '', '2024-06-02 10:00:00', 'Sales', '2024-06-01 10:00:00', 'admin', 'Rp. 66,600,000', 'Rp. 6,600,000', NULL, NULL, 'SI/581/63X79X5915', '1. Ditolak, 2. Ditolak, 3. Ditolak');
INSERT INTO public.performance_invoice VALUES (15, 1, 'Rp. 10.150.000', 'Diproses', 'Ortopedi', 'PI/581/7263Z4X2Z3', 'PO/581/7263C4A2C3', '2024-07-01', NULL, NULL, '2024-07-02 14:30:00.70458', 'sales', '2024-07-02 14:30:00.70458', 'sales', 'Rp. 11.266.500', 'Rp. 1.116.500', NULL, NULL, 'SI/581/63X79X5915', NULL);
INSERT INTO public.performance_invoice VALUES (16, 1, 'Rp. 10.150.000', 'Diproses', 'Ortopedi', 'PI/581/7263Z4X2Z3', 'PO/581/7263C4A2C3', '2024-07-01', NULL, NULL, '2024-07-02 14:30:45.681076', 'sales', '2024-07-02 14:30:45.681076', 'sales', 'Rp. 11.266.500', 'Rp. 1.116.500', NULL, NULL, 'SI/581/63X79X5915', NULL);
INSERT INTO public.performance_invoice VALUES (17, 1, 'Rp. 10.150.000', 'Diproses', 'Ortopedi', 'PI/525/Y46509Z603', 'PO/525/B46509C603', '2024-07-02', 'DR. AHMAD EFENDI', 'REHAN LATUCONSINA', '2024-07-03 14:21:38.082248', 'sales', '2024-07-03 14:21:38.082248', 'sales', 'Rp. 11.266.500', 'Rp. 1.116.500', '2024-08-12', '1092882', 'SI/525/Y46509Z603', NULL);
INSERT INTO public.performance_invoice VALUES (18, 1, 'Rp. 10.150.000', 'Diproses', 'Ortopedi', 'PI/525/Y46509Z603', 'PO/525/B46509C603', '2024-07-02', 'DR. AHMAD EFENDI', 'REHAN LATUCONSINA', '2024-07-06 07:24:40.554983', 'sales', '2024-07-06 07:24:40.554983', 'sales', 'Rp. 11.266.500', 'Rp. 1.116.500', '2024-08-12', '1092882', 'SI/525/Y46509Z603', NULL);
INSERT INTO public.performance_invoice VALUES (19, 1, 'Rp. 10.150.000', 'Diproses', 'Ortopedi', 'PI/525/Y46509Z603', 'PO/525/B46509C603', '2024-07-02', 'DR. AHMAD EFENDI', 'REHAN LATUCONSINA', '2024-07-06 07:25:03.063931', 'sales', '2024-07-06 07:25:03.063931', 'sales', 'Rp. 11.266.500', 'Rp. 1.116.500', '2024-08-12', '1092882', 'SI/525/Y46509Z603', NULL);
INSERT INTO public.performance_invoice VALUES (20, 3, 'Rp. 2.850.000', 'Diproses', 'Ortopedi', 'PI/417/69Y4XX36X4', 'PO/417/69B4AA36A4', '2024-07-06', '', '', '2024-07-06 12:00:32.332196', 'sales', '2024-07-06 12:00:32.332196', 'sales', 'Rp. 3.163.500', 'Rp. 313.500', '', '', 'SI/417/69Y4XX36X4', NULL);
INSERT INTO public.performance_invoice VALUES (21, 1, 'Rp. 43.700.000', 'Diproses', 'Ortopedi', 'PI/142/ZY48YZY77Z', 'PO/142/CB48BCB77C', '2024-07-06', NULL, NULL, '2024-07-06 14:55:07.071272', 'sales', '2024-07-06 14:55:07.071272', 'sales', 'Rp. 48.507.000', 'Rp. 4.807.000', NULL, NULL, 'SI/142/2344596879', NULL);
INSERT INTO public.performance_invoice VALUES (22, 1, 'Rp. 1.600', 'Diproses', 'Ortopedi', 'PI/992/7039X5465X', 'PO/992/7039A5465A', '2024-07-08', NULL, NULL, '2024-07-08 10:08:48.342944', 'sales', '2024-07-08 10:08:48.342944', 'sales', 'Rp. 1.776', 'Rp. 176', NULL, NULL, 'SI/992/7039X5465X', NULL);
INSERT INTO public.performance_invoice VALUES (26, 1, 'Rp. 10.150.000', 'Diproses', 'Ortopedi', 'PI/581/7263Z4X2Z3', 'PO/581/7263C4A2C3', '2024-07-01', NULL, NULL, '2024-07-09 03:37:29.564957', 'sales', '2024-07-09 03:37:29.564957', 'sales', 'Rp. 11.266.500', 'Rp. 1.116.500', NULL, NULL, 'SI/581/63X79X5915', NULL);
INSERT INTO public.performance_invoice VALUES (10, 1, 'Rp. 24.100.000', 'Diproses', 'Radiologi', 'PI/525/Y46509Z603', 'PO/525/B46509C603', '2024-07-02', 'DR. AHMAD EFENDI', 'REHAN LATUCONSINA', '2024-06-22 13:21:49.952658', 'sales', '2024-06-26 04:29:27.74466', 'sales', 'Rp. 26.751.000', 'Rp. 2.651.000', '2024-08-12', '1092882', 'SI/525/Y46509Z603', NULL);
INSERT INTO public.performance_invoice VALUES (27, 1, 'Rp. 10.150.000', 'Diproses', 'Ortopedi', 'PI/581/7263Z4X2Z3', 'PO/581/7263C4A2C3', '2024-07-01', NULL, NULL, '2024-07-09 03:38:43.191594', 'sales', '2024-07-09 03:38:43.191594', 'sales', 'Rp. 11.266.500', 'Rp. 1.116.500', NULL, NULL, 'SI/581/63X79X5915', NULL);
INSERT INTO public.performance_invoice VALUES (28, 1, 'Rp. 10.150.000', 'Diproses', 'Ortopedi', 'PI/525/Y46509Z603', 'PO/525/B46509C603', '2024-07-02', 'DR. AHMAD EFENDI', 'REHAN LATUCONSINA', '2024-07-09 04:58:00.201573', 'sales', '2024-07-09 04:58:00.201573', 'sales', 'Rp. 11.266.500', 'Rp. 1.116.500', '2024-08-12', '1092882', 'SI/525/Y46509Z603', NULL);
INSERT INTO public.performance_invoice VALUES (29, 1, 'Rp. 10.150.000', 'Diproses', 'Ortopedi', 'PI/525/Y46509Z603', 'PO/525/B46509C603', '2024-07-02', 'DR. AHMAD EFENDI', 'REHAN LATUCONSINA', '2024-07-09 04:58:02.921971', 'sales', '2024-07-09 04:58:02.921971', 'sales', 'Rp. 11.266.500', 'Rp. 1.116.500', '2024-08-12', '1092882', 'SI/525/Y46509Z603', NULL);
INSERT INTO public.performance_invoice VALUES (13, 1, 'Rp. 5.355.156.488.650', 'Diterima', 'Ortopedi', 'PI/581/7263Z4X2Z3', 'PO/581/7263C4A2C3', '2024-07-01', NULL, NULL, '2024-06-23 00:36:20.235691', 'sales', '2024-06-26 04:36:24.715556', 'sales', 'Rp. 5.944.223.702.401', 'Rp. 589.067.213.751', NULL, NULL, 'SI/581/63X79X5915', NULL);
INSERT INTO public.performance_invoice VALUES (30, 3, 'Rp. 164.284', 'Diproses', 'Ortopedi', 'PI/173/0X74117561', 'PO/173/0A74117561', '2024-07-10', NULL, NULL, '2024-07-10 12:42:11.099732', 'sales', '2024-07-10 12:42:11.099732', 'sales', 'Rp. 182.355', 'Rp. 18.071', NULL, NULL, 'SI/173/0X74117561', NULL);
INSERT INTO public.performance_invoice VALUES (31, 1, 'Rp. 96.000', 'Diproses', 'Ortopedi', 'PI/104/9Y77209604', 'PO/490/9B77209604', '2024-07-11', '', '', '2024-07-11 12:08:17.638114', 'sales', '2024-07-11 12:08:17.638114', 'sales', 'Rp. 106.560', 'Rp. 10.560', '', '', 'SI/104/9Y77209604', NULL);
INSERT INTO public.performance_invoice VALUES (32, 1, 'Rp. 90.000.000', 'Diproses', 'Ortopedi', 'PI/063/768Y84247Z', 'PO/063/768B84247C', '2024-07-11', '', '', '2024-07-11 12:15:27.70179', 'sales', '2024-07-11 12:15:27.70179', 'sales', 'Rp. 99.900.000', 'Rp. 9.900.000', '', '', 'SI/063/768Y84247Z', NULL);


--
-- TOC entry 3519 (class 0 OID 16634)
-- Dependencies: 228
-- Data for Name: piutang; Type: TABLE DATA; Schema: public; Owner: fismed-user
--

INSERT INTO public.piutang VALUES (2, 'MAYAPADA HOSPITAL', '11266500', 'Rp. 11.266.500', '2024-07-09 03:37:29.564957');
INSERT INTO public.piutang VALUES (3, 'MAYAPADA HOSPITAL', '11266500', 'Rp. 11.266.500', '2024-07-09 03:38:43.191594');
INSERT INTO public.piutang VALUES (4, 'MAYAPADA HOSPITAL', '11266500', 'Rp. 11.266.500', '2024-07-09 04:58:00.201573');
INSERT INTO public.piutang VALUES (5, 'MAYAPADA HOSPITAL', '11266500', 'Rp. 11.266.500', '2024-07-09 04:58:02.921971');
INSERT INTO public.piutang VALUES (6, 'MAYAPADA', '106560', 'Rp. 106.560', '2024-07-11 12:08:17.638114');
INSERT INTO public.piutang VALUES (7, 'MAYAPADA', '99900000', 'Rp. 99.900.000', '2024-07-11 12:15:27.70179');


--
-- TOC entry 3527 (class 0 OID 16689)
-- Dependencies: 236
-- Data for Name: purchase_order; Type: TABLE DATA; Schema: public; Owner: fismed-user
--



--
-- TOC entry 3508 (class 0 OID 16470)
-- Dependencies: 217
-- Data for Name: stock_items; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.stock_items VALUES (5, 'Item A', '100', '10000', '2024-06-15 02:53:32.450216', 'admin', '2024-06-15 02:53:32.450216', 'admin');
INSERT INTO public.stock_items VALUES (3, 'Item C', '140', '15000', '2024-06-03 14:00:00', 'admin', '2024-06-07 17:00:00', 'admin');
INSERT INTO public.stock_items VALUES (2, 'Item B', '180', '20000', '2024-06-02 13:00:00', 'admin', '2024-06-06 16:00:00', 'admin');
INSERT INTO public.stock_items VALUES (1, 'Item B', '9880', '100000', '2024-06-01 12:00:00', 'admin', '2024-06-05 15:00:00', 'admin');


--
-- TOC entry 3510 (class 0 OID 16476)
-- Dependencies: 219
-- Data for Name: tax_code; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.tax_code VALUES (1, 'SWASTA', '2024-06-12 22:21:26.726137', 'admin', '2024-06-12 22:21:26.726137', 'admin');
INSERT INTO public.tax_code VALUES (2, 'PEMERINTAH', '2024-06-12 22:21:26.726137', 'admin', '2024-06-12 22:21:26.726137', 'admin');
INSERT INTO public.tax_code VALUES (3, 'PEMERINTAH NON', '2024-06-12 22:21:26.726137', 'admin', '2024-06-12 22:21:26.726137', 'admin');


--
-- TOC entry 3512 (class 0 OID 16482)
-- Dependencies: 221
-- Data for Name: user_category; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.user_category VALUES (1, 'SALES', '2024-06-11 17:11:16.111105', 'system', '2024-06-11 17:11:16.111105', 'system');
INSERT INTO public.user_category VALUES (2, 'ADMIN', '2024-06-11 17:12:52.189158', 'system', '2024-06-11 17:12:52.189158', 'system');
INSERT INTO public.user_category VALUES (3, 'LOGISTIK', '2024-06-11 17:12:52.189158', 'system', '2024-06-11 17:12:52.189158', 'system');
INSERT INTO public.user_category VALUES (4, 'KEUANGAN', '2024-06-11 17:12:52.189158', 'system', '2024-06-11 17:12:52.189158', 'system');
INSERT INTO public.user_category VALUES (5, 'SUPER ADMIN', '2024-06-11 17:12:52.189158', 'system', '2024-06-11 17:12:52.189158', 'system');


--
-- TOC entry 3514 (class 0 OID 16490)
-- Dependencies: 223
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.users VALUES (1, 'sales1', '0ad80eb119d9bf7775aa23786b05b391', 1, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTk5MTgzMzYsImlhdCI6MTcxOTg4OTUzNiwic3ViIjoiMSJ9.PET6zqfM6QePXXC5tPWgIeKDEdqroBdB6bCdogbiSnw', '2024-06-11 17:15:48.925181', 'system', '2024-06-11 17:15:48.925181', 'system');


--
-- TOC entry 3552 (class 0 OID 0)
-- Dependencies: 210
-- Name: customer_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.customer_id_seq', 21, true);


--
-- TOC entry 3553 (class 0 OID 0)
-- Dependencies: 212
-- Name: divisi_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.divisi_id_seq', 3, true);


--
-- TOC entry 3554 (class 0 OID 0)
-- Dependencies: 233
-- Name: galery_id_seq; Type: SEQUENCE SET; Schema: public; Owner: fismed-user
--

SELECT pg_catalog.setval('public.galery_id_seq', 1, true);


--
-- TOC entry 3555 (class 0 OID 0)
-- Dependencies: 231
-- Name: hitang_id_seq; Type: SEQUENCE SET; Schema: public; Owner: fismed-user
--

SELECT pg_catalog.setval('public.hitang_id_seq', 1, false);


--
-- TOC entry 3556 (class 0 OID 0)
-- Dependencies: 237
-- Name: item_buyer_id_seq; Type: SEQUENCE SET; Schema: public; Owner: fismed-user
--

SELECT pg_catalog.setval('public.item_buyer_id_seq', 1, false);


--
-- TOC entry 3557 (class 0 OID 0)
-- Dependencies: 214
-- Name: order_items_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.order_items_id_seq', 53, true);


--
-- TOC entry 3558 (class 0 OID 0)
-- Dependencies: 225
-- Name: pemasukan_id_seq; Type: SEQUENCE SET; Schema: public; Owner: fismed-user
--

SELECT pg_catalog.setval('public.pemasukan_id_seq', 6, true);


--
-- TOC entry 3559 (class 0 OID 0)
-- Dependencies: 229
-- Name: pengeluaran_id_seq; Type: SEQUENCE SET; Schema: public; Owner: fismed-user
--

SELECT pg_catalog.setval('public.pengeluaran_id_seq', 1, false);


--
-- TOC entry 3560 (class 0 OID 0)
-- Dependencies: 216
-- Name: performance_invoice_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.performance_invoice_id_seq', 32, true);


--
-- TOC entry 3561 (class 0 OID 0)
-- Dependencies: 227
-- Name: piutang_id_seq; Type: SEQUENCE SET; Schema: public; Owner: fismed-user
--

SELECT pg_catalog.setval('public.piutang_id_seq', 7, true);


--
-- TOC entry 3562 (class 0 OID 0)
-- Dependencies: 235
-- Name: purchase_order_id_seq; Type: SEQUENCE SET; Schema: public; Owner: fismed-user
--

SELECT pg_catalog.setval('public.purchase_order_id_seq', 1, false);


--
-- TOC entry 3563 (class 0 OID 0)
-- Dependencies: 218
-- Name: stock_items_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.stock_items_id_seq', 5, true);


--
-- TOC entry 3564 (class 0 OID 0)
-- Dependencies: 220
-- Name: tax_code_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.tax_code_id_seq', 3, true);


--
-- TOC entry 3565 (class 0 OID 0)
-- Dependencies: 222
-- Name: user_category_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.user_category_id_seq', 5, true);


--
-- TOC entry 3566 (class 0 OID 0)
-- Dependencies: 224
-- Name: user_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.user_id_seq', 1, true);


--
-- TOC entry 3325 (class 2606 OID 16507)
-- Name: customer customer_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.customer
    ADD CONSTRAINT customer_pkey PRIMARY KEY (id);


--
-- TOC entry 3327 (class 2606 OID 16509)
-- Name: divisi divisi_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.divisi
    ADD CONSTRAINT divisi_pkey PRIMARY KEY (id);


--
-- TOC entry 3351 (class 2606 OID 16668)
-- Name: galery galery_pkey; Type: CONSTRAINT; Schema: public; Owner: fismed-user
--

ALTER TABLE ONLY public.galery
    ADD CONSTRAINT galery_pkey PRIMARY KEY (id);


--
-- TOC entry 3349 (class 2606 OID 16659)
-- Name: hitang hitang_pkey; Type: CONSTRAINT; Schema: public; Owner: fismed-user
--

ALTER TABLE ONLY public.hitang
    ADD CONSTRAINT hitang_pkey PRIMARY KEY (id);


--
-- TOC entry 3355 (class 2606 OID 16705)
-- Name: item_buyer item_buyer_pkey; Type: CONSTRAINT; Schema: public; Owner: fismed-user
--

ALTER TABLE ONLY public.item_buyer
    ADD CONSTRAINT item_buyer_pkey PRIMARY KEY (id);


--
-- TOC entry 3329 (class 2606 OID 16511)
-- Name: order_items order_items_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.order_items
    ADD CONSTRAINT order_items_pkey PRIMARY KEY (id);


--
-- TOC entry 3343 (class 2606 OID 16632)
-- Name: pemasukan pemasukan_pkey; Type: CONSTRAINT; Schema: public; Owner: fismed-user
--

ALTER TABLE ONLY public.pemasukan
    ADD CONSTRAINT pemasukan_pkey PRIMARY KEY (id);


--
-- TOC entry 3347 (class 2606 OID 16650)
-- Name: pengeluaran pengeluaran_pkey; Type: CONSTRAINT; Schema: public; Owner: fismed-user
--

ALTER TABLE ONLY public.pengeluaran
    ADD CONSTRAINT pengeluaran_pkey PRIMARY KEY (id);


--
-- TOC entry 3331 (class 2606 OID 16513)
-- Name: performance_invoice performance_invoice_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.performance_invoice
    ADD CONSTRAINT performance_invoice_pkey PRIMARY KEY (id);


--
-- TOC entry 3345 (class 2606 OID 16641)
-- Name: piutang piutang_pkey; Type: CONSTRAINT; Schema: public; Owner: fismed-user
--

ALTER TABLE ONLY public.piutang
    ADD CONSTRAINT piutang_pkey PRIMARY KEY (id);


--
-- TOC entry 3353 (class 2606 OID 16696)
-- Name: purchase_order purchase_order_pkey; Type: CONSTRAINT; Schema: public; Owner: fismed-user
--

ALTER TABLE ONLY public.purchase_order
    ADD CONSTRAINT purchase_order_pkey PRIMARY KEY (id);


--
-- TOC entry 3333 (class 2606 OID 16515)
-- Name: stock_items stock_items_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.stock_items
    ADD CONSTRAINT stock_items_pkey PRIMARY KEY (id);


--
-- TOC entry 3335 (class 2606 OID 16517)
-- Name: tax_code tax_code_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.tax_code
    ADD CONSTRAINT tax_code_pkey PRIMARY KEY (id);


--
-- TOC entry 3337 (class 2606 OID 16519)
-- Name: user_category user_category_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.user_category
    ADD CONSTRAINT user_category_pkey PRIMARY KEY (id);


--
-- TOC entry 3339 (class 2606 OID 16521)
-- Name: users user_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT user_pkey PRIMARY KEY (id);


--
-- TOC entry 3341 (class 2606 OID 16523)
-- Name: users user_username_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT user_username_key UNIQUE (username);


--
-- TOC entry 3356 (class 2606 OID 16524)
-- Name: customer customer_tax_code_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.customer
    ADD CONSTRAINT customer_tax_code_id_fkey FOREIGN KEY (tax_code_id) REFERENCES public.tax_code(id);


--
-- TOC entry 3358 (class 2606 OID 16529)
-- Name: performance_invoice fk_customer; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.performance_invoice
    ADD CONSTRAINT fk_customer FOREIGN KEY (customer_id) REFERENCES public.customer(id);


--
-- TOC entry 3360 (class 2606 OID 16706)
-- Name: item_buyer item_buyer_po_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: fismed-user
--

ALTER TABLE ONLY public.item_buyer
    ADD CONSTRAINT item_buyer_po_id_fkey FOREIGN KEY (po_id) REFERENCES public.purchase_order(id);


--
-- TOC entry 3357 (class 2606 OID 16534)
-- Name: order_items order_items_pi_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.order_items
    ADD CONSTRAINT order_items_pi_id_fkey FOREIGN KEY (pi_id) REFERENCES public.performance_invoice(id);


--
-- TOC entry 3359 (class 2606 OID 16539)
-- Name: users user_role_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT user_role_id_fkey FOREIGN KEY (role_id) REFERENCES public.user_category(id);


--
-- TOC entry 3536 (class 0 OID 0)
-- Dependencies: 5
-- Name: SCHEMA public; Type: ACL; Schema: -; Owner: pg_database_owner
--

REVOKE USAGE ON SCHEMA public FROM PUBLIC;


-- Completed on 2024-07-12 14:47:57

--
-- PostgreSQL database dump complete
--

