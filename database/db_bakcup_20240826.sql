--
-- PostgreSQL database dump
--

-- Dumped from database version 14.12
-- Dumped by pg_dump version 16.1

-- Started on 2024-08-26 10:25:08

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
-- TOC entry 5 (class 2615 OID 16555)
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
-- TOC entry 209 (class 1259 OID 16556)
-- Name: customer; Type: TABLE; Schema: public; Owner: fismed-user
--

CREATE TABLE public.customer (
    id integer NOT NULL,
    nama_company character varying,
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
    handphone character varying,
    nama character varying,
    kategori character varying
);


ALTER TABLE public.customer OWNER TO "fismed-user";

--
-- TOC entry 210 (class 1259 OID 16561)
-- Name: customer_id_seq; Type: SEQUENCE; Schema: public; Owner: fismed-user
--

CREATE SEQUENCE public.customer_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.customer_id_seq OWNER TO "fismed-user";

--
-- TOC entry 3537 (class 0 OID 0)
-- Dependencies: 210
-- Name: customer_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: fismed-user
--

ALTER SEQUENCE public.customer_id_seq OWNED BY public.customer.id;


--
-- TOC entry 211 (class 1259 OID 16562)
-- Name: divisi; Type: TABLE; Schema: public; Owner: fismed-user
--

CREATE TABLE public.divisi (
    id integer NOT NULL,
    name character varying NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    created_by character varying NOT NULL,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_by character varying NOT NULL
);


ALTER TABLE public.divisi OWNER TO "fismed-user";

--
-- TOC entry 212 (class 1259 OID 16569)
-- Name: divisi_id_seq; Type: SEQUENCE; Schema: public; Owner: fismed-user
--

CREATE SEQUENCE public.divisi_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.divisi_id_seq OWNER TO "fismed-user";

--
-- TOC entry 3538 (class 0 OID 0)
-- Dependencies: 212
-- Name: divisi_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: fismed-user
--

ALTER SEQUENCE public.divisi_id_seq OWNED BY public.divisi.id;


--
-- TOC entry 213 (class 1259 OID 16570)
-- Name: galery; Type: TABLE; Schema: public; Owner: fismed-user
--

CREATE TABLE public.galery (
    id integer NOT NULL,
    nama character varying,
    iamge text
);


ALTER TABLE public.galery OWNER TO "fismed-user";

--
-- TOC entry 214 (class 1259 OID 16575)
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
-- Dependencies: 214
-- Name: galery_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: fismed-user
--

ALTER SEQUENCE public.galery_id_seq OWNED BY public.galery.id;


--
-- TOC entry 238 (class 1259 OID 16739)
-- Name: hutang; Type: TABLE; Schema: public; Owner: fismed-user
--

CREATE TABLE public.hutang (
    id integer NOT NULL,
    nama character varying,
    nominal character varying,
    amount character varying
);


ALTER TABLE public.hutang OWNER TO "fismed-user";

--
-- TOC entry 237 (class 1259 OID 16738)
-- Name: hutang_id_seq; Type: SEQUENCE; Schema: public; Owner: fismed-user
--

CREATE SEQUENCE public.hutang_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.hutang_id_seq OWNER TO "fismed-user";

--
-- TOC entry 3540 (class 0 OID 0)
-- Dependencies: 237
-- Name: hutang_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: fismed-user
--

ALTER SEQUENCE public.hutang_id_seq OWNED BY public.hutang.id;


--
-- TOC entry 215 (class 1259 OID 16582)
-- Name: item_buyer; Type: TABLE; Schema: public; Owner: fismed-user
--

CREATE TABLE public.item_buyer (
    id integer NOT NULL,
    po_id integer,
    name character varying,
    quantity character varying,
    price character varying,
    discount character varying,
    amount character varying
);


ALTER TABLE public.item_buyer OWNER TO "fismed-user";

--
-- TOC entry 216 (class 1259 OID 16587)
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
-- Dependencies: 216
-- Name: item_buyer_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: fismed-user
--

ALTER SEQUENCE public.item_buyer_id_seq OWNED BY public.item_buyer.id;


--
-- TOC entry 217 (class 1259 OID 16588)
-- Name: order_items; Type: TABLE; Schema: public; Owner: fismed-user
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


ALTER TABLE public.order_items OWNER TO "fismed-user";

--
-- TOC entry 218 (class 1259 OID 16593)
-- Name: order_items_id_seq; Type: SEQUENCE; Schema: public; Owner: fismed-user
--

CREATE SEQUENCE public.order_items_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.order_items_id_seq OWNER TO "fismed-user";

--
-- TOC entry 3542 (class 0 OID 0)
-- Dependencies: 218
-- Name: order_items_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: fismed-user
--

ALTER SEQUENCE public.order_items_id_seq OWNED BY public.order_items.id;


--
-- TOC entry 219 (class 1259 OID 16594)
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
-- TOC entry 220 (class 1259 OID 16599)
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
-- Dependencies: 220
-- Name: pemasukan_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: fismed-user
--

ALTER SEQUENCE public.pemasukan_id_seq OWNED BY public.pemasukan.id;


--
-- TOC entry 221 (class 1259 OID 16600)
-- Name: pengeluaran; Type: TABLE; Schema: public; Owner: fismed-user
--

CREATE TABLE public.pengeluaran (
    id integer NOT NULL,
    nama character varying,
    sub_total character varying,
    pajak character varying,
    tanggal character varying,
    total character varying
);


ALTER TABLE public.pengeluaran OWNER TO "fismed-user";

--
-- TOC entry 222 (class 1259 OID 16605)
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
-- Dependencies: 222
-- Name: pengeluaran_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: fismed-user
--

ALTER SEQUENCE public.pengeluaran_id_seq OWNED BY public.pengeluaran.id;


--
-- TOC entry 223 (class 1259 OID 16606)
-- Name: performance_invoice; Type: TABLE; Schema: public; Owner: fismed-user
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


ALTER TABLE public.performance_invoice OWNER TO "fismed-user";

--
-- TOC entry 224 (class 1259 OID 16611)
-- Name: performance_invoice_id_seq; Type: SEQUENCE; Schema: public; Owner: fismed-user
--

CREATE SEQUENCE public.performance_invoice_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.performance_invoice_id_seq OWNER TO "fismed-user";

--
-- TOC entry 3545 (class 0 OID 0)
-- Dependencies: 224
-- Name: performance_invoice_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: fismed-user
--

ALTER SEQUENCE public.performance_invoice_id_seq OWNED BY public.performance_invoice.id;


--
-- TOC entry 225 (class 1259 OID 16612)
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
-- TOC entry 226 (class 1259 OID 16617)
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
-- Dependencies: 226
-- Name: piutang_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: fismed-user
--

ALTER SEQUENCE public.piutang_id_seq OWNED BY public.piutang.id;


--
-- TOC entry 227 (class 1259 OID 16618)
-- Name: purchase_order; Type: TABLE; Schema: public; Owner: fismed-user
--

CREATE TABLE public.purchase_order (
    id integer NOT NULL,
    nama_suplier character varying,
    nomor_po character varying,
    tanggal character varying,
    catatan_po text,
    prepared_by character varying,
    prepared_jabatan character varying,
    approved_by character varying,
    approved_jabatan character varying,
    status character varying,
    created_at timestamp without time zone,
    created_by character varying,
    updated_at timestamp without time zone,
    updated_by character varying,
    sub_total character varying,
    pajak character varying,
    total character varying,
    reason text
);


ALTER TABLE public.purchase_order OWNER TO "fismed-user";

--
-- TOC entry 228 (class 1259 OID 16623)
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
-- Dependencies: 228
-- Name: purchase_order_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: fismed-user
--

ALTER SEQUENCE public.purchase_order_id_seq OWNED BY public.purchase_order.id;


--
-- TOC entry 229 (class 1259 OID 16624)
-- Name: stock_items; Type: TABLE; Schema: public; Owner: fismed-user
--

CREATE TABLE public.stock_items (
    id integer NOT NULL,
    name character varying NOT NULL,
    total text,
    price character varying,
    created_at timestamp without time zone,
    created_by character varying,
    updated_at timestamp without time zone,
    updated_by character varying,
    katalog character varying,
    gudang character varying
);


ALTER TABLE public.stock_items OWNER TO "fismed-user";

--
-- TOC entry 230 (class 1259 OID 16629)
-- Name: stock_items_id_seq; Type: SEQUENCE; Schema: public; Owner: fismed-user
--

CREATE SEQUENCE public.stock_items_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.stock_items_id_seq OWNER TO "fismed-user";

--
-- TOC entry 3548 (class 0 OID 0)
-- Dependencies: 230
-- Name: stock_items_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: fismed-user
--

ALTER SEQUENCE public.stock_items_id_seq OWNED BY public.stock_items.id;


--
-- TOC entry 231 (class 1259 OID 16630)
-- Name: tax_code; Type: TABLE; Schema: public; Owner: fismed-user
--

CREATE TABLE public.tax_code (
    id integer NOT NULL,
    tax character varying(50),
    created_at timestamp without time zone,
    created_by character varying,
    updated_at timestamp without time zone,
    updated_by character varying
);


ALTER TABLE public.tax_code OWNER TO "fismed-user";

--
-- TOC entry 232 (class 1259 OID 16635)
-- Name: tax_code_id_seq; Type: SEQUENCE; Schema: public; Owner: fismed-user
--

CREATE SEQUENCE public.tax_code_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.tax_code_id_seq OWNER TO "fismed-user";

--
-- TOC entry 3549 (class 0 OID 0)
-- Dependencies: 232
-- Name: tax_code_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: fismed-user
--

ALTER SEQUENCE public.tax_code_id_seq OWNED BY public.tax_code.id;


--
-- TOC entry 233 (class 1259 OID 16636)
-- Name: user_category; Type: TABLE; Schema: public; Owner: fismed-user
--

CREATE TABLE public.user_category (
    id integer NOT NULL,
    role character varying(50) NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    created_by character varying(255),
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    updated_by character varying(255)
);


ALTER TABLE public.user_category OWNER TO "fismed-user";

--
-- TOC entry 234 (class 1259 OID 16643)
-- Name: user_category_id_seq; Type: SEQUENCE; Schema: public; Owner: fismed-user
--

CREATE SEQUENCE public.user_category_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.user_category_id_seq OWNER TO "fismed-user";

--
-- TOC entry 3550 (class 0 OID 0)
-- Dependencies: 234
-- Name: user_category_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: fismed-user
--

ALTER SEQUENCE public.user_category_id_seq OWNED BY public.user_category.id;


--
-- TOC entry 235 (class 1259 OID 16644)
-- Name: users; Type: TABLE; Schema: public; Owner: fismed-user
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


ALTER TABLE public.users OWNER TO "fismed-user";

--
-- TOC entry 236 (class 1259 OID 16651)
-- Name: user_id_seq; Type: SEQUENCE; Schema: public; Owner: fismed-user
--

CREATE SEQUENCE public.user_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.user_id_seq OWNER TO "fismed-user";

--
-- TOC entry 3551 (class 0 OID 0)
-- Dependencies: 236
-- Name: user_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: fismed-user
--

ALTER SEQUENCE public.user_id_seq OWNED BY public.users.id;


--
-- TOC entry 3303 (class 2604 OID 16652)
-- Name: customer id; Type: DEFAULT; Schema: public; Owner: fismed-user
--

ALTER TABLE ONLY public.customer ALTER COLUMN id SET DEFAULT nextval('public.customer_id_seq'::regclass);


--
-- TOC entry 3304 (class 2604 OID 16653)
-- Name: divisi id; Type: DEFAULT; Schema: public; Owner: fismed-user
--

ALTER TABLE ONLY public.divisi ALTER COLUMN id SET DEFAULT nextval('public.divisi_id_seq'::regclass);


--
-- TOC entry 3307 (class 2604 OID 16654)
-- Name: galery id; Type: DEFAULT; Schema: public; Owner: fismed-user
--

ALTER TABLE ONLY public.galery ALTER COLUMN id SET DEFAULT nextval('public.galery_id_seq'::regclass);


--
-- TOC entry 3323 (class 2604 OID 16742)
-- Name: hutang id; Type: DEFAULT; Schema: public; Owner: fismed-user
--

ALTER TABLE ONLY public.hutang ALTER COLUMN id SET DEFAULT nextval('public.hutang_id_seq'::regclass);


--
-- TOC entry 3308 (class 2604 OID 16656)
-- Name: item_buyer id; Type: DEFAULT; Schema: public; Owner: fismed-user
--

ALTER TABLE ONLY public.item_buyer ALTER COLUMN id SET DEFAULT nextval('public.item_buyer_id_seq'::regclass);


--
-- TOC entry 3309 (class 2604 OID 16657)
-- Name: order_items id; Type: DEFAULT; Schema: public; Owner: fismed-user
--

ALTER TABLE ONLY public.order_items ALTER COLUMN id SET DEFAULT nextval('public.order_items_id_seq'::regclass);


--
-- TOC entry 3310 (class 2604 OID 16658)
-- Name: pemasukan id; Type: DEFAULT; Schema: public; Owner: fismed-user
--

ALTER TABLE ONLY public.pemasukan ALTER COLUMN id SET DEFAULT nextval('public.pemasukan_id_seq'::regclass);


--
-- TOC entry 3311 (class 2604 OID 16659)
-- Name: pengeluaran id; Type: DEFAULT; Schema: public; Owner: fismed-user
--

ALTER TABLE ONLY public.pengeluaran ALTER COLUMN id SET DEFAULT nextval('public.pengeluaran_id_seq'::regclass);


--
-- TOC entry 3312 (class 2604 OID 16660)
-- Name: performance_invoice id; Type: DEFAULT; Schema: public; Owner: fismed-user
--

ALTER TABLE ONLY public.performance_invoice ALTER COLUMN id SET DEFAULT nextval('public.performance_invoice_id_seq'::regclass);


--
-- TOC entry 3313 (class 2604 OID 16661)
-- Name: piutang id; Type: DEFAULT; Schema: public; Owner: fismed-user
--

ALTER TABLE ONLY public.piutang ALTER COLUMN id SET DEFAULT nextval('public.piutang_id_seq'::regclass);


--
-- TOC entry 3314 (class 2604 OID 16662)
-- Name: purchase_order id; Type: DEFAULT; Schema: public; Owner: fismed-user
--

ALTER TABLE ONLY public.purchase_order ALTER COLUMN id SET DEFAULT nextval('public.purchase_order_id_seq'::regclass);


--
-- TOC entry 3315 (class 2604 OID 16663)
-- Name: stock_items id; Type: DEFAULT; Schema: public; Owner: fismed-user
--

ALTER TABLE ONLY public.stock_items ALTER COLUMN id SET DEFAULT nextval('public.stock_items_id_seq'::regclass);


--
-- TOC entry 3316 (class 2604 OID 16664)
-- Name: tax_code id; Type: DEFAULT; Schema: public; Owner: fismed-user
--

ALTER TABLE ONLY public.tax_code ALTER COLUMN id SET DEFAULT nextval('public.tax_code_id_seq'::regclass);


--
-- TOC entry 3317 (class 2604 OID 16665)
-- Name: user_category id; Type: DEFAULT; Schema: public; Owner: fismed-user
--

ALTER TABLE ONLY public.user_category ALTER COLUMN id SET DEFAULT nextval('public.user_category_id_seq'::regclass);


--
-- TOC entry 3320 (class 2604 OID 16666)
-- Name: users id; Type: DEFAULT; Schema: public; Owner: fismed-user
--

ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.user_id_seq'::regclass);


--
-- TOC entry 3500 (class 0 OID 16556)
-- Dependencies: 209
-- Data for Name: customer; Type: TABLE DATA; Schema: public; Owner: fismed-user
--

INSERT INTO public.customer VALUES (2, 'OETOMO', '456 Market St', '456 NPWP St', '0987654321', 'IPAK456', '456 Facture St', 'City B', '67890', '234-567-8901', 'emailB@company.com', '234-567-8902', 'PIC B', '456 Item St', 'City B', '67890', '234-567-8903', 'itemB@company.com', '234-567-8904', 'PIC Item B', 'Contact B', 2, 'Gak tau ini apa ', '2024-06-12 22:23:31.420925', 'admin', '2024-06-12 22:23:31.420925', 'admin', NULL, NULL, 'RS');
INSERT INTO public.customer VALUES (3, 'BOROMEUS', '789 Broadway', '789 NPWP St', '1122334455', 'IPAK789', '789 Facture St', 'City C', '54321', '345-678-9012', 'emailC@company.com', '345-678-9013', 'PIC C', '789 Item St', 'City C', '54321', '345-678-9014', 'itemC@company.com', '345-678-9015', 'PIC Item C', 'Contact C', 3, 'Gak tau ini apa ', '2024-06-12 22:23:31.420925', 'admin', '2024-06-12 22:23:31.420925', 'admin', NULL, NULL, 'RS');
INSERT INTO public.customer VALUES (35, 'MAYAPIDI', 'bandung', 'bandung', '123', '123', 'bandung', 'bandung', '40164', '081519419720', '', '', 'adit', 'bandung', 'bandung', '40164', '081519419720', '', '', 'rehan', 'Gilbert Ionisiys', 1, '1', '2024-08-01 04:17:37.316695', 'admin', '2024-08-01 04:17:37.316695', 'admin', '123', 'Gilbert Ionisiys', 'RS');
INSERT INTO public.customer VALUES (5, 'SANTO YUSUF', '202 Oak St', '202 NPWP St', '3344556677', 'IPAK202', '202 Facture St', 'City E', '87654', '567-890-1234', 'emailE@company.com', '567-890-1235', 'PIC E', '202 Item St', 'City E', '87654', '567-890-1236', 'itemE@company.com', '567-890-1237', 'PIC Item E', 'Contact E', 2, 'Gak tau ini apa ', '2024-06-12 22:23:31.420925', 'admin', '2024-06-12 22:23:31.420925', 'admin', NULL, NULL, 'RS');
INSERT INTO public.customer VALUES (4, 'KARTINI', '101 State St', '101 NPWP St', '2233445566', 'IPAK101', '101 Facture St', 'City D', '98765', '456-789-0123', 'emailD@company.com', '456-789-0124', 'PIC D', '101 Item St', 'City D', '98765', '456-789-0125', 'itemD@company.com', '456-789-0126', 'PIC Item D', 'Contact D', 1, 'Gak tau ini apa ', '2024-06-12 22:23:31.420925', 'admin', '2024-06-12 22:23:31.420925', 'admin', NULL, NULL, 'RS');
INSERT INTO public.customer VALUES (1, 'MAYAPADA', 'Bandung', '123 NPWP St', '1234567890', 'IPAK123', '123 Facture St', 'City A', '12345', '123-456-7890', 'emailA@company.com', '123-456-7891', 'PIC A', '123 Item St', 'City A', '12345', '123-456-7892', 'itemA@company.com', '123-456-7893', 'PIC Item A', 'Contact A', 1, 'Gak tau ini apa ', '2024-06-12 22:23:31.420925', 'admin', '2024-06-12 22:23:31.420925', 'admin', NULL, NULL, 'RS');
INSERT INTO public.customer VALUES (36, 'OETOMO', 'Jl. Raya Bojongsoang No.156, Lengkong, Kec. Bojongsoang, Kabupaten Bandung, Jawa Barat 40287', '123', '123', '123', '23', '1231', '123', '123', '', '123', '123', '231', '1231', '123', '123', '', '123', '123', '123', 1, '123', '2024-08-08 16:45:51.793781', 'admin', '2024-08-08 16:45:51.793781', 'admin', '23', NULL, 'SP');
INSERT INTO public.customer VALUES (37, 'GEPPA', '123', '123', '123', '123', '123', '23', '123', '123', '', '123', '123', '123', '123', '123', '123', '', '123', '123', '23', 1, '123', '2024-08-08 16:46:31.102276', 'admin', '2024-08-08 16:46:31.102276', 'admin', '23', NULL, 'SP');
INSERT INTO public.customer VALUES (38, 'Telkom University', 'asd', 'asd', 'qwwe', '123', 'asdasd test', 'Kota Bandung', '40287', '123123123123', '', '123', '132', '123', '123', '123', '123', '', '123', '123', '123', 2, '123', '2024-08-26 03:08:09.803872', 'admin', '2024-08-26 03:08:09.803872', 'admin', '123', NULL, 'SP');
INSERT INTO public.customer VALUES (27, 'RSUD CIAWI', 'Jl. Raya Bojongsoang No.156, Lengkong, Kec. Bojongsoang, Kabupaten Bandung, Jawa Barat 40287', 'Jl. Raya Bojongsoang No.156, Lengkong, Kec. Bojongsoang, Kabupaten Bandung, Jawa Barat 40287', '123', '12345', 'Jl. Raya Bojongsoang No.156, Lengkong, Kec. Bojongsoang, Kabupaten Bandung, Jawa Barat 40287', 'BOGOR', '8888', '08745485655', 'emailA@company.com', '02254885', 'DEDEN', 'GG IMPRESS', 'BOGOR', '66666', '08123456789', 'emailA@company.com', '02254885', 'ADIT', '085234949686', 1, '30', '2024-07-16 15:28:41.063686', 'admin', '2024-07-16 15:28:41.063686', 'admin', '085234379548', 'emailA@company.com', 'SP');
INSERT INTO public.customer VALUES (33, '1', '1', '1', '1', '1', '1', '1', '1', '1', '', '1', '1', '', '1', '1', '1', '', '1', '1', '1', 1, '1', '2024-07-19 05:42:45.695699', 'admin', '2024-07-19 05:42:45.695699', 'admin', '1', '1', 'RS');
INSERT INTO public.customer VALUES (34, 'Rumah Sakit A', 'Rumah Sakit A', 'Rumah Sakit A', 'Rumah Sakit A', 'Rumah Sakit A', 'Rumah Sakit A', 'Rumah Sakit A', 'Rumah Sakit A', 'Rumah Sakit A', '', 'Rumah Sakit A', 'Rumah Sakit A', 'Rumah Sakit A', 'Rumah Sakit A', 'Rumah Sakit A', 'Rumah Sakit A', '', 'Rumah Sakit A', 'Rumah Sakit A', 'Rumah Sakit A', 1, 'Rumah Sakit A', '2024-07-19 08:20:25.786686', 'admin', '2024-07-19 08:20:25.786686', 'admin', 'Rumah Sakit A', NULL, 'SP');
INSERT INTO public.customer VALUES (39, 'Sedang Bekerja di perusahaan lain', 'asdwdwd', 'asdwdwd', '123123', '123123', 'Raditya Kost, Bojongsoang, Kabupaten Bandung', 'Bandung', '40287', '082113252387', '', 'ADITYAPATTY24@GMAIL.COM', '132', 'Bekasi, Jalan Tambora raya nomor 3 Jawabarat', 'KOTA BEKASI', '123', '0105390001', '', '1151a.p@gmail.com', '123', 'Aditya Patty', 1, '123', '2024-08-26 03:17:56.562697', 'admin', '2024-08-26 03:17:56.562697', 'admin', '123', NULL, 'SP');


--
-- TOC entry 3502 (class 0 OID 16562)
-- Dependencies: 211
-- Data for Name: divisi; Type: TABLE DATA; Schema: public; Owner: fismed-user
--

INSERT INTO public.divisi VALUES (1, 'RADIOLOGI', '2024-06-21 10:12:55.531313', 'admin', '2024-06-21 10:12:55.531313', 'admin');
INSERT INTO public.divisi VALUES (2, 'ORTOPEDI', '2024-06-21 10:12:55.531313', 'admin', '2024-06-21 10:12:55.531313', 'admin');
INSERT INTO public.divisi VALUES (3, 'KARDIOLOGI', '2024-06-21 10:12:55.531313', 'admin', '2024-06-21 10:12:55.531313', 'admin');


--
-- TOC entry 3504 (class 0 OID 16570)
-- Dependencies: 213
-- Data for Name: galery; Type: TABLE DATA; Schema: public; Owner: fismed-user
--

INSERT INTO public.galery VALUES (1, 'logo', '/9j/4AAQSkZJRgABAQAAAQABAAD//gA7Q1JFQVRPUjogZ2QtanBlZyB2MS4wICh1c2luZyBJSkcgSlBFRyB2NjIpLCBxdWFsaXR5ID0gODIK/9sAQwAGBAQFBAQGBQUFBgYGBwkOCQkICAkSDQ0KDhUSFhYVEhQUFxohHBcYHxkUFB0nHR8iIyUlJRYcKSwoJCshJCUk/9sAQwEGBgYJCAkRCQkRJBgUGCQkJCQkJCQkJCQkJCQkJCQkJCQkJCQkJCQkJCQkJCQkJCQkJCQkJCQkJCQkJCQkJCQk/8AAEQgAMgAyAwEiAAIRAQMRAf/EAB8AAAEFAQEBAQEBAAAAAAAAAAABAgMEBQYHCAkKC//EALUQAAIBAwMCBAMFBQQEAAABfQECAwAEEQUSITFBBhNRYQcicRQygZGhCCNCscEVUtHwJDNicoIJChYXGBkaJSYnKCkqNDU2Nzg5OkNERUZHSElKU1RVVldYWVpjZGVmZ2hpanN0dXZ3eHl6g4SFhoeIiYqSk5SVlpeYmZqio6Slpqeoqaqys7S1tre4ubrCw8TFxsfIycrS09TV1tfY2drh4uPk5ebn6Onq8fLz9PX29/j5+v/EAB8BAAMBAQEBAQEBAQEAAAAAAAABAgMEBQYHCAkKC//EALURAAIBAgQEAwQHBQQEAAECdwABAgMRBAUhMQYSQVEHYXETIjKBCBRCkaGxwQkjM1LwFWJy0QoWJDThJfEXGBkaJicoKSo1Njc4OTpDREVGR0hJSlNUVVZXWFlaY2RlZmdoaWpzdHV2d3h5eoKDhIWGh4iJipKTlJWWl5iZmqKjpKWmp6ipqrKztLW2t7i5usLDxMXGx8jJytLT1NXW19jZ2uLj5OXm5+jp6vLz9PX29/j5+v/aAAwDAQACEQMRAD8A+qKDwKR3VFLMwUDkknGKrHVLHH/H3b/9/BTs2tERKcY7s8q8YeNr/UtQmtbK4ktrOFyg8pirS44JJ649qyNI8U6to1ys0F3NIoOWhlcsj+xB/nWbeOpu5zuX/Wv3H941GCG4X5j6Dmvia9XE+1dR3Wp4ksQ5Tumew+F9YTxJrN3qcaSJDDbxQJu6bjln/wDZR+FdXXker+HvFejaFpSaJLNCWLzXflzLGVkbG3cSR8oAx+ea9D0XX7e/toI47hb64UCOeS2UtGsgA3fN0Az719nQoyjh4VJtNy3XY78Pi+abpTVmvx9DZoooqjvPDfip4vvNS1q40eC4ZLC1IjdEOBK+Mnd6gZxjpxmuACLkfKv5V2PxP8PXGjeJrm6MbfZL1/Njl/h3EfMufXIJx6EVy1pZT30yxQRs7McDAzk195g54ajhIzukran5Pmf1qrjZRldu+noZ6wtLN5cabmJIAArtdBfT/B0MepX0YuLr/WW9qOPMbtIx/hQdu5PIFST+HYvBWmLfaoiPfXOfs9m/U+rSeij+6OSSOa4q9vZr+5kuJ5DJJI25mPVjXh+w/tKXPNWpJ6f3n/kenzywWn2/y/4J2OpeINR8VbbnU7ppomO5Ldflij5OPl7n3PNWdD1+98P3aXFpKwUEeZFn5ZF9CP69qxNLH/Evg/3f6mtLT9OudVvI7O0jaSaQ4AA6D1PoB61+a5lia8sbKEW9HZI97CXajNfEz3y2v4Lq3injbKSoHU+oIyKKistMis7OC2UErDGsYJ9AMUV7ydbsfR3mN1XSItYga1ujutnGHj2qd34kHH4VUsPD+ieF7Z5rW0ht1ijJeU8sFAycsea2q5zx+858L3Nrasq3F60dpHu6ZkcIf0JruhHmkot6GVeMYRdVL3kjwnxXr83ifW7m/mJKOSkKn+CIH5R/X6msfYn91fyrovGPg298JX5SVHks3OYbjHDD0PoR6Vk6Vpd5rV7HZafA1xPIcBV6D3J7D3r9Ew9TDqgpQa5Uj8nxNLFPEyjNPmbPZvhz4Y0bUPB2nXF1pttNK6tudkBJ+c12WnaNp+khxY2cFvv5by0Az9ar+F9G/wCEf0Cy03fva3iCs3q3U/qTWrX53XjTlWlUit2z9WwVH2dCEZLVJBRRRUnWFY+vKrXOkBgCPtoOCPSNyKKKqG5lW+A05Y0ljKSIrqRghhkGo7O0t7VStvbxQgnJEaBf5UUVtH+Gzln/ABkWaKKK5zvCiiigD//Z');


--
-- TOC entry 3529 (class 0 OID 16739)
-- Dependencies: 238
-- Data for Name: hutang; Type: TABLE DATA; Schema: public; Owner: fismed-user
--



--
-- TOC entry 3506 (class 0 OID 16582)
-- Dependencies: 215
-- Data for Name: item_buyer; Type: TABLE DATA; Schema: public; Owner: fismed-user
--

INSERT INTO public.item_buyer VALUES (48, 37, 'Kalung', '900', 'Rp. 8.000', '12%', 'Rp. 6.336.000');
INSERT INTO public.item_buyer VALUES (49, 37, 'Kayu', '800', 'Rp. 9.000', '13%', 'Rp. 6.264.000');
INSERT INTO public.item_buyer VALUES (50, 37, 'Kamoceng', '700', 'Rp. 7.000', '14%', 'Rp. 4.214.000');
INSERT INTO public.item_buyer VALUES (51, 37, 'Kasur Operasi Tulang', '10', 'Rp. 1.000.000', '0%', 'Rp. 10.000.000');
INSERT INTO public.item_buyer VALUES (54, 39, 'Kayu', '100', 'Rp. 9.000', '90%', 'Rp. 90.000');
INSERT INTO public.item_buyer VALUES (52, 38, 'Kayu', '240', 'Rp. 9.000', '10%', 'Rp. 1.944.000');
INSERT INTO public.item_buyer VALUES (55, 40, 'Kayu', '240', 'Rp. 9.000', '10%', 'Rp. 1.944.000');
INSERT INTO public.item_buyer VALUES (57, 42, 'kancut', '20', 'Rp. 20.000', '10%', 'Rp. 360.000');


--
-- TOC entry 3508 (class 0 OID 16588)
-- Dependencies: 217
-- Data for Name: order_items; Type: TABLE DATA; Schema: public; Owner: fismed-user
--

INSERT INTO public.order_items VALUES (58, 38, 'Kasur Operasi Tulang', '30', '9000', '20', 'Rp. 216.000', 'KAT123', '2024-07-19 08:15:27.659667', 'sales', '2024-08-01 03:53:50.162449', 'sales');
INSERT INTO public.order_items VALUES (73, 38, 'Kalung', '40', '8000', '70', 'Rp. 96.000', 'kat404', '2024-08-01 03:53:50.159842', 'sales', '2024-08-01 03:53:50.159842', 'sales');
INSERT INTO public.order_items VALUES (74, 48, 'Kasur Operasi Tulang', '1', '10000', '0', 'Rp. 10.000', '123', '2024-08-01 04:29:47.943909', 'sales', '2024-08-01 04:29:47.943909', 'sales');
INSERT INTO public.order_items VALUES (59, 39, 'Kasur Operasi Tulang', '30', '1000000', '20', '', 'KAT12441', '2024-07-19 08:16:10.070013', 'sales', '2024-07-19 08:16:10.070013', 'sales');
INSERT INTO public.order_items VALUES (60, 40, 'Kasur Operasi Tulang', '5', '2000', '10', '', '123', '2024-07-19 09:13:33.862922', 'sales', '2024-07-19 09:13:33.862922', 'sales');
INSERT INTO public.order_items VALUES (61, 41, 'Kayu', '123123', '9000', '13', 'Rp. 964.053.090', 'qwe', '2024-07-28 16:13:21.024471', 'sales', '2024-07-28 16:13:21.024471', 'sales');
INSERT INTO public.order_items VALUES (62, 41, 'Kamoceng', '123', '7000', '13', 'Rp. 749.070', 'qwe', '2024-07-28 16:13:21.024471', 'sales', '2024-07-28 16:13:21.024471', 'sales');
INSERT INTO public.order_items VALUES (63, 42, 'Kayu', '123', '9000', '13', 'Rp. 963.090', 'qwe', '2024-07-28 16:30:35.508411', 'sales', '2024-07-28 16:30:35.508411', 'sales');
INSERT INTO public.order_items VALUES (64, 43, 'Kayu', '123', '9000', '13', 'Rp. 963.090', 'qwe', '2024-07-28 16:31:33.460915', 'sales', '2024-07-28 16:31:33.460915', 'sales');
INSERT INTO public.order_items VALUES (65, 44, 'Kayu', '123', '9000', '13', 'Rp. 963.090', 'qwe', '2024-07-28 16:34:14.354907', 'sales', '2024-07-28 16:34:14.354907', 'sales');
INSERT INTO public.order_items VALUES (66, 45, 'Kayu', '123', '9000', '13', 'Rp. 963.090', 'qwe', '2024-07-28 16:47:06.64366', 'sales', '2024-07-28 16:47:06.64366', 'sales');
INSERT INTO public.order_items VALUES (67, 46, 'Kayu', '123', '9000', '13', 'Rp. 963.090', 'qwe', '2024-07-28 16:47:46.555433', 'sales', '2024-07-28 16:47:46.555433', 'sales');
INSERT INTO public.order_items VALUES (68, 47, 'SEMPAK', '50', '10000', '50', 'Rp. 250.000', '123', '2024-07-29 14:29:07.450351', 'sales', '2024-07-29 14:29:07.450351', 'sales');
INSERT INTO public.order_items VALUES (69, 42, 'Kamoceng', '123', '7000', '12', 'Rp. 757.680', 'asd', '2024-07-31 09:40:49.284479', 'sales', '2024-07-31 09:40:49.284479', 'sales');
INSERT INTO public.order_items VALUES (70, 42, 'Kalung', '1414', '8000', '12', 'Rp. 9.954.560', 'qet', '2024-07-31 09:41:42.66245', 'sales', '2024-07-31 09:41:42.66245', 'sales');
INSERT INTO public.order_items VALUES (71, 45, 'Kamoceng', '12', '7000000', '13', 'Rp. 73.080.000', 'qwe', '2024-08-01 03:23:05.729356', 'sales', '2024-08-01 03:23:05.729356', 'sales');
INSERT INTO public.order_items VALUES (57, 37, 'Kasur Operasi Tulang', '30', '1000000', '20', 'Rp. 24.000.000', 'Kat002003', '2024-07-19 08:11:23.758646', 'sales', '2024-08-01 10:48:35.010952', 'sales');
INSERT INTO public.order_items VALUES (72, 37, 'Kayu', '313', '9000', '13', 'Rp. 2.450.790', 'k123', '2024-08-01 03:31:16.123114', 'sales', '2024-08-01 10:48:35.328706', 'sales');


--
-- TOC entry 3510 (class 0 OID 16594)
-- Dependencies: 219
-- Data for Name: pemasukan; Type: TABLE DATA; Schema: public; Owner: fismed-user
--

INSERT INTO public.pemasukan VALUES (15, 'MAYAPADA', '26640000', 'Rp. 26.640.000', '2024-07-19 08:11:23.758646');
INSERT INTO public.pemasukan VALUES (16, 'MAYAPADA', '26640000', 'Rp. 26.640.000', '2024-07-19 08:15:27.659667');
INSERT INTO public.pemasukan VALUES (17, 'MAYAPADA', '26640000', 'Rp. 26.640.000', '2024-07-19 08:16:10.070013');
INSERT INTO public.pemasukan VALUES (18, 'MAYAPADA', '9990', 'Rp. 9.990', '2024-07-19 09:13:33.862922');
INSERT INTO public.pemasukan VALUES (19, 'SANTO YUSUF', '1070098929', 'Rp. 1.070.098.929', '2024-07-28 10:51:12.684529');
INSERT INTO public.pemasukan VALUES (20, 'SANTO YUSUF', '1070930397', 'Rp. 1.070.930.397', '2024-07-28 15:38:40.152277');
INSERT INTO public.pemasukan VALUES (21, 'SANTO YUSUF', '1070930397', 'Rp. 1.070.930.397', '2024-07-28 16:13:21.024471');
INSERT INTO public.pemasukan VALUES (22, 'SANTO YUSUF', '1069029', 'Rp. 1.069.029', '2024-07-28 16:30:35.508411');
INSERT INTO public.pemasukan VALUES (23, 'SANTO YUSUF', '1069029', 'Rp. 1.069.029', '2024-07-28 16:31:33.460915');
INSERT INTO public.pemasukan VALUES (24, 'SANTO YUSUF', '1069029', 'Rp. 1.069.029', '2024-07-28 16:34:14.354907');
INSERT INTO public.pemasukan VALUES (25, 'SANTO YUSUF', '1069029', 'Rp. 1.069.029', '2024-07-28 16:47:06.64366');
INSERT INTO public.pemasukan VALUES (26, 'SANTO YUSUF', '1069029', 'Rp. 1.069.029', '2024-07-28 16:47:46.555433');
INSERT INTO public.pemasukan VALUES (27, 'Rumah Sakit A', '277500', 'Rp. 277.500', '2024-07-29 14:29:07.450351');
INSERT INTO public.pemasukan VALUES (28, 'MAYAPIDI', '11100', 'Rp. 11.100', '2024-08-01 04:29:47.943909');


--
-- TOC entry 3512 (class 0 OID 16600)
-- Dependencies: 221
-- Data for Name: pengeluaran; Type: TABLE DATA; Schema: public; Owner: fismed-user
--

INSERT INTO public.pengeluaran VALUES (4, 'Ochistok Official', 'Rp. 160.000.000', 'Rp. 17.600.000', '2024-07-19', 'Rp. 177.600.000');
INSERT INTO public.pengeluaran VALUES (5, 'Ochistok Official', 'Rp. 160.000.000', 'Rp. 17.600.000', '2024-07-19', 'Rp. 177.600.000');
INSERT INTO public.pengeluaran VALUES (6, 'Ochistok Official', 'Rp. 160.000.000', 'Rp. 17.600.000', '2024-07-19', 'Rp. 177.600.000');
INSERT INTO public.pengeluaran VALUES (7, 'Ochistok Official', 'Rp. 160.000.000', 'Rp. 17.600.000', '2024-07-19', 'Rp. 177.600.000');
INSERT INTO public.pengeluaran VALUES (8, 'Ochistok Official', 'Rp. 160.000.000', 'Rp. 17.600.000', '2024-07-19', 'Rp. 177.600.000');
INSERT INTO public.pengeluaran VALUES (9, 'Ochistok Official', 'Rp. 154.000.000', 'Rp. 16.940.000', '2024-07-19', 'Rp. 170.940.000');
INSERT INTO public.pengeluaran VALUES (10, 'Ochistok Official', 'Rp. 154.000.000', 'Rp. 16.940.000', '2024-07-19', 'Rp. 170.940.000');
INSERT INTO public.pengeluaran VALUES (11, 'Ochistok Official', 'Rp. 160.000.000', 'Rp. 17.600.000', '2024-07-19', 'Rp. 177.600.000');
INSERT INTO public.pengeluaran VALUES (12, 'PT Global Indonesia', '15144129', '1665854', '2024-07-23', '16809983');
INSERT INTO public.pengeluaran VALUES (13, 'PTPT', '500000', '55000', '2024-07-26', '555000');
INSERT INTO public.pengeluaran VALUES (14, 'Aditya', '3845315854413', '422984743985', '2024-07-26', '4268300598398');
INSERT INTO public.pengeluaran VALUES (15, 'Aditya', '29300000', '3223000', '2024-07-27', '32523000');
INSERT INTO public.pengeluaran VALUES (16, 'adit', '400000', '44000', '2024-07-29', '444000');


--
-- TOC entry 3514 (class 0 OID 16606)
-- Dependencies: 223
-- Data for Name: performance_invoice; Type: TABLE DATA; Schema: public; Owner: fismed-user
--

INSERT INTO public.performance_invoice VALUES (47, 34, 'Rp. 250.000', 'Ditolak', 'Ortopedi', 'PI/405/913667043X', NULL, '2024-07-31', 'RADIT', 'IKBAL', '2024-07-29 14:29:07.450351', 'SALES', '2024-07-29 14:29:07.450351', 'SALES', 'Rp. 277.500', 'Rp. 27.500', '2024-07-25', '123', 'SI/248/X3145252Z8', 'REHAN GANTENG');
INSERT INTO public.performance_invoice VALUES (37, 1, 'Rp. 26.450.790', 'Diterima', 'Ortopedi', 'PI/753/6Z494X82Z0', 'PO/600/31C1148633', '2024-07-19', 'Rehan Sangat Tampan', 'Rehan Sangat Tampan', '2024-07-19 08:11:23.758646', 'sales', '2024-08-01 04:04:37.812907', 'admin', 'Rp. 29.360.376', 'Rp. 2.909.586', '2024-08-29', '123', 'SI/734/384Z51X2X3', 'REHAN GANTENG');
INSERT INTO public.performance_invoice VALUES (48, 35, 'Rp. 10.000', 'Diterima', 'Ortopedi', 'PI/373/8712004064', NULL, '2024-08-30', 'Gilbert', 'Reyhan', '2024-08-01 04:29:47.943909', 'SALES', '2024-08-01 04:31:19.294399', 'admin', 'Rp. 11.100', 'Rp. 1.100', '2024-08-29', NULL, 'SI/410/986YZ6Y984', NULL);
INSERT INTO public.performance_invoice VALUES (38, 1, 'Rp. 312.000', 'Diterima', 'Ortopedi', 'PI/394/69Z945X833', 'PO/644/9734207522', '2024-07-19', 'Rehan Sangat Tampan', 'Rehan Sangat Tampan', '2024-07-19 08:15:27.659995', 'sales', '2024-08-01 04:36:44.673028', 'admin', 'Rp. 346.320', 'Rp. 34.320', '2024-08-15', '123', 'SI/503/99XY65626X', 'REHAN GANTENG');
INSERT INTO public.performance_invoice VALUES (45, 5, 'Rp. 74.043.090', 'Ditolak', 'Radiologi', 'PI/338/062XZX0181', '', '2024-07-31', '', '', '2024-07-28 16:47:06.64366', 'SALES', '2024-08-01 03:23:05.729356', 'sales', 'Rp. 82.187.829', 'Rp. 8.144.739', '', '123', 'SI/338/062XZX0181', 'REHAN GANTENG');
INSERT INTO public.performance_invoice VALUES (41, 5, 'Rp. 964.802.160', 'Ditolak', 'Ortopedi', 'PI/730/Y718908213', '', '2024-07-30', 'ADITYA COBA APAHA BERHASIL', 'Ridwan Hanif', '2024-07-28 16:13:21.024471', 'SALES', '2024-08-01 03:24:30.307669', 'admin', 'Rp. 1.070.930.397', 'Rp. 106.128.237', '2024-08-29', '123', 'SI/730/Y718908213', 'REHAN GANTENG');
INSERT INTO public.performance_invoice VALUES (42, 5, 'Rp. 11.675.330', 'Ditolak', 'Ortopedi', 'PI/926/8Z14Y2X001', '', '2024-07-23', 'ADITYA COBA APAHA BERHASIL', 'ADITYA COBA APAHA BERHASIL', '2024-07-28 16:30:35.508411', 'SALES', '2024-08-01 03:24:57.302018', 'admin', 'Rp. 12.959.616', 'Rp. 1.284.286', '2024-07-18', '123', 'SI/926/8Z14Y2X001', 'REHAN GANTENG');
INSERT INTO public.performance_invoice VALUES (39, 1, 'Rp. 24.000.000', 'Diterima', 'Ortopedi', 'PI/470/5Z97337Z40', 'PO/619/B8A759778A', '2024-07-19', NULL, NULL, '2024-07-19 08:16:10.075165', 'sales', '2024-08-08 16:53:55.865033', 'admin', 'Rp. 26.640.000', 'Rp. 2.640.000', NULL, '123', 'SI/302/3Y15716930', 'REHAN GANTENG');
INSERT INTO public.performance_invoice VALUES (40, 1, 'Rp. 9.000', 'Diterima', 'Ortopedi', 'PI/154/Y41Z723978', 'PO/212/8AB74C5CB1', '2024-07-19', '', '', '2024-07-19 09:13:33.862922', 'sales', '2024-08-08 16:54:01.776015', 'admin', 'Rp. 9.990', 'Rp. 990', '', '123', 'SI/731/YZ90X446X7', 'REHAN GANTENG');
INSERT INTO public.performance_invoice VALUES (43, 5, 'Rp. 963.090', 'Ditolak', 'Ortopedi', 'PI/232/1Z7X674924', NULL, '2024-08-05', 'ADITYA COBA APAHA BERHASIL', 'ADITYA COBA APAHA BERHASIL', '2024-07-28 16:31:33.460915', 'SALES', '2024-07-28 16:31:33.460915', 'SALES', 'Rp. 1.069.029', 'Rp. 105.939', '2024-07-03', '123', 'SI/232/1Z7X674924', 'REHAN GANTENG');
INSERT INTO public.performance_invoice VALUES (44, 5, 'Rp. 963.090', 'Ditolak', 'Radiologi', 'PI/979/29Z04Y1536', NULL, '2024-07-25', NULL, NULL, '2024-07-28 16:34:14.354907', 'SALES', '2024-07-28 16:34:14.354907', 'SALES', 'Rp. 1.069.029', 'Rp. 105.939', NULL, '123', 'SI/979/29Z04Y1536', 'REHAN GANTENG');
INSERT INTO public.performance_invoice VALUES (46, 5, 'Rp. 963.090', 'Ditolak', 'Radiologi', 'PI/211/377097X701', NULL, '2024-07-31', NULL, NULL, '2024-07-28 16:47:46.555433', 'SALES', '2024-07-29 07:02:06.25787', 'admin', 'Rp. 1.069.029', 'Rp. 105.939', NULL, '123', 'SI/101/377097X701', 'REHAN GANTENG');


--
-- TOC entry 3516 (class 0 OID 16612)
-- Dependencies: 225
-- Data for Name: piutang; Type: TABLE DATA; Schema: public; Owner: fismed-user
--

INSERT INTO public.piutang VALUES (2, 'MAYAPADA HOSPITAL', '11266500', 'Rp. 11.266.500', '2024-07-09 03:37:29.564957');
INSERT INTO public.piutang VALUES (3, 'MAYAPADA HOSPITAL', '11266500', 'Rp. 11.266.500', '2024-07-09 03:38:43.191594');
INSERT INTO public.piutang VALUES (4, 'MAYAPADA HOSPITAL', '11266500', 'Rp. 11.266.500', '2024-07-09 04:58:00.201573');
INSERT INTO public.piutang VALUES (5, 'MAYAPADA HOSPITAL', '11266500', 'Rp. 11.266.500', '2024-07-09 04:58:02.921971');
INSERT INTO public.piutang VALUES (6, 'MAYAPADA', '106560', 'Rp. 106.560', '2024-07-11 12:08:17.638114');
INSERT INTO public.piutang VALUES (7, 'MAYAPADA', '99900000', 'Rp. 99.900.000', '2024-07-11 12:15:27.70179');
INSERT INTO public.piutang VALUES (8, '', '', '', '2024-07-18 06:34:18.057138');
INSERT INTO public.piutang VALUES (9, '', '', '', '2024-07-18 06:56:03.497361');
INSERT INTO public.piutang VALUES (10, '', 'Rp. 14.792.785', '', '2024-07-18 07:24:19.72401');
INSERT INTO public.piutang VALUES (11, '', 'Rp. 14.792.785', '', '2024-07-18 07:30:54.445419');
INSERT INTO public.piutang VALUES (12, 'RSUD CIAWI', '1998', 'Rp. 1.998', '2024-07-18 16:13:52.540507');
INSERT INTO public.piutang VALUES (13, 'MAYAPADA', '49950', 'Rp. 49.950', '2024-07-19 03:36:41.101644');
INSERT INTO public.piutang VALUES (14, '1', '0', 'Rp. 0', '2024-07-19 05:43:27.480143');
INSERT INTO public.piutang VALUES (15, '1', '99900', 'Rp. 99.900', '2024-07-19 05:53:17.427828');
INSERT INTO public.piutang VALUES (16, 'MAYAPADA', '26640000', 'Rp. 26.640.000', '2024-07-19 08:11:23.758646');
INSERT INTO public.piutang VALUES (17, 'MAYAPADA', '26640000', 'Rp. 26.640.000', '2024-07-19 08:15:27.659667');
INSERT INTO public.piutang VALUES (18, 'MAYAPADA', '26640000', 'Rp. 26.640.000', '2024-07-19 08:16:10.070013');
INSERT INTO public.piutang VALUES (19, 'MAYAPADA', '9990', 'Rp. 9.990', '2024-07-19 09:13:33.862922');
INSERT INTO public.piutang VALUES (20, 'SANTO YUSUF', '1070098929', 'Rp. 1.070.098.929', '2024-07-28 10:51:12.684529');
INSERT INTO public.piutang VALUES (21, 'SANTO YUSUF', '1070930397', 'Rp. 1.070.930.397', '2024-07-28 15:38:40.152277');
INSERT INTO public.piutang VALUES (22, 'SANTO YUSUF', '1070930397', 'Rp. 1.070.930.397', '2024-07-28 16:13:21.024471');
INSERT INTO public.piutang VALUES (23, 'SANTO YUSUF', '1069029', 'Rp. 1.069.029', '2024-07-28 16:30:35.508411');
INSERT INTO public.piutang VALUES (24, 'SANTO YUSUF', '1069029', 'Rp. 1.069.029', '2024-07-28 16:31:33.460915');
INSERT INTO public.piutang VALUES (25, 'SANTO YUSUF', '1069029', 'Rp. 1.069.029', '2024-07-28 16:34:14.354907');
INSERT INTO public.piutang VALUES (26, 'SANTO YUSUF', '1069029', 'Rp. 1.069.029', '2024-07-28 16:47:06.64366');
INSERT INTO public.piutang VALUES (27, 'SANTO YUSUF', '1069029', 'Rp. 1.069.029', '2024-07-28 16:47:46.555433');
INSERT INTO public.piutang VALUES (28, 'Rumah Sakit A', '277500', 'Rp. 277.500', '2024-07-29 14:29:07.450351');
INSERT INTO public.piutang VALUES (29, 'MAYAPIDI', '11100', 'Rp. 11.100', '2024-08-01 04:29:47.943909');


--
-- TOC entry 3518 (class 0 OID 16618)
-- Dependencies: 227
-- Data for Name: purchase_order; Type: TABLE DATA; Schema: public; Owner: fismed-user
--

INSERT INTO public.purchase_order VALUES (42, 'adit', '6A59248C76', '2024-07-29', 'KNTL', 'Patty', 'BABU', 'PRANGIN', 'ADMIN', 'DITERIMA', '2024-07-29 14:47:34.89501', 'ADMIN', '2024-07-29 14:48:38.876018', 'admin', 'Rp. 360.000', 'Rp. 39.600', 'Rp. 399.600', NULL);
INSERT INTO public.purchase_order VALUES (37, 'Aditya', '39C6659B8B', '2024-07-27', 'Purchase Order untuk restock barang ', 'Patty', 'Manager Product', 'Gilbert', 'Manager Project', 'DITERIMA', '2024-07-27 01:51:16.075479', 'ADMIN', '2024-07-26 19:35:37.467032', 'admin', 'Rp. 26.814.000', 'Rp. 2.949.540', 'Rp. 29.763.540', NULL);
INSERT INTO public.purchase_order VALUES (40, 'Perangin angin', '5656B20251', '2024-07-28', 'Purchase Order untuk restock barang ', 'Patty', 'Manager Product', 'Gilbert', 'Manager Project', 'DIPROSES', '2024-07-28 00:52:56.792746', 'ADMIN', '2024-07-27 17:54:44.249314', '2024-07-28', 'Rp. 1.944.000', 'Rp. 213.840', 'Rp. 2.157.840', '');
INSERT INTO public.purchase_order VALUES (39, 'Aditya', '2714555813', '2024-07-27', 'Purchase Order untuk restock barang ', 'Patty', 'Manager Product', 'Gilbert', 'Manager Project', 'DITOLAK', '2024-07-27 01:56:12.909183', 'ADMIN', '2024-07-27 17:22:43.376096', 'admin', 'Rp. 90.000', 'Rp. 9.900', 'Rp. 99.900', 'Harga Terlalu Murah');
INSERT INTO public.purchase_order VALUES (38, 'Akbar Verari', '38B71803AB', '2024-07-28', 'Purchase Order untuk restock barang ', 'Patty', 'Manager Product', 'Gilbert', 'Manager Project', 'DITOLAK', '2024-07-27 01:55:08.424106', 'ADMIN', '2024-07-28 16:49:44.637231', 'admin', 'Rp. 1.944.000', 'Rp. 213.840', 'Rp. 2.157.840', '1
1213123

123123
123
123
123
12312
3123
123');


--
-- TOC entry 3520 (class 0 OID 16624)
-- Dependencies: 229
-- Data for Name: stock_items; Type: TABLE DATA; Schema: public; Owner: fismed-user
--

INSERT INTO public.stock_items VALUES (14, 'Kalung', '105', '8000', '2024-07-19 08:06:47.525744', 'admin', '2024-07-19 08:06:47.525744', 'admin', 'K234', 'Gudang Utara');
INSERT INTO public.stock_items VALUES (15, 'Kamoceng', '123633', '7000', '2024-07-19 08:06:47.525744', 'admin', '2024-07-19 08:06:47.525744', 'admin', 'K345', 'Gudang Barat');
INSERT INTO public.stock_items VALUES (13, 'Kayu', '123633', '9000', '2024-07-19 08:06:47.525744', 'admin', '2024-07-19 08:06:47.525744', 'admin', 'K456', 'Gudang Timur');
INSERT INTO public.stock_items VALUES (12, 'Kasur Operasi Tulang', '104', '1000000', '2024-07-19 08:06:47.525744', 'admin', '2024-07-19 08:06:47.525744', 'admin', 'K123', 'Gudang Selatan ');


--
-- TOC entry 3522 (class 0 OID 16630)
-- Dependencies: 231
-- Data for Name: tax_code; Type: TABLE DATA; Schema: public; Owner: fismed-user
--

INSERT INTO public.tax_code VALUES (1, 'SWASTA', '2024-06-12 22:21:26.726137', 'admin', '2024-06-12 22:21:26.726137', 'admin');
INSERT INTO public.tax_code VALUES (2, 'PEMERINTAH', '2024-06-12 22:21:26.726137', 'admin', '2024-06-12 22:21:26.726137', 'admin');
INSERT INTO public.tax_code VALUES (3, 'PEMERINTAH NON', '2024-06-12 22:21:26.726137', 'admin', '2024-06-12 22:21:26.726137', 'admin');


--
-- TOC entry 3524 (class 0 OID 16636)
-- Dependencies: 233
-- Data for Name: user_category; Type: TABLE DATA; Schema: public; Owner: fismed-user
--

INSERT INTO public.user_category VALUES (1, 'SALES', '2024-06-11 17:11:16.111105', 'system', '2024-06-11 17:11:16.111105', 'system');
INSERT INTO public.user_category VALUES (2, 'ADMIN', '2024-06-11 17:12:52.189158', 'system', '2024-06-11 17:12:52.189158', 'system');
INSERT INTO public.user_category VALUES (3, 'LOGISTIK', '2024-06-11 17:12:52.189158', 'system', '2024-06-11 17:12:52.189158', 'system');
INSERT INTO public.user_category VALUES (4, 'KEUANGAN', '2024-06-11 17:12:52.189158', 'system', '2024-06-11 17:12:52.189158', 'system');
INSERT INTO public.user_category VALUES (5, 'SUPER ADMIN', '2024-06-11 17:12:52.189158', 'system', '2024-06-11 17:12:52.189158', 'system');


--
-- TOC entry 3526 (class 0 OID 16644)
-- Dependencies: 235
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: fismed-user
--

INSERT INTO public.users VALUES (1, 'sales1', '0ad80eb119d9bf7775aa23786b05b391', 1, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MjEzOTE3MzIsImlhdCI6MTcyMTM2MjkzMiwic3ViIjoiMSJ9.8NboMak9CeY1TgSIU4ZS0YsCbDdiaEtEzoZB2QqYqIs', '2024-06-11 17:15:48.925181', 'system', '2024-06-11 17:15:48.925181', 'system');


--
-- TOC entry 3552 (class 0 OID 0)
-- Dependencies: 210
-- Name: customer_id_seq; Type: SEQUENCE SET; Schema: public; Owner: fismed-user
--

SELECT pg_catalog.setval('public.customer_id_seq', 39, true);


--
-- TOC entry 3553 (class 0 OID 0)
-- Dependencies: 212
-- Name: divisi_id_seq; Type: SEQUENCE SET; Schema: public; Owner: fismed-user
--

SELECT pg_catalog.setval('public.divisi_id_seq', 3, true);


--
-- TOC entry 3554 (class 0 OID 0)
-- Dependencies: 214
-- Name: galery_id_seq; Type: SEQUENCE SET; Schema: public; Owner: fismed-user
--

SELECT pg_catalog.setval('public.galery_id_seq', 1, true);


--
-- TOC entry 3555 (class 0 OID 0)
-- Dependencies: 237
-- Name: hutang_id_seq; Type: SEQUENCE SET; Schema: public; Owner: fismed-user
--

SELECT pg_catalog.setval('public.hutang_id_seq', 1, false);


--
-- TOC entry 3556 (class 0 OID 0)
-- Dependencies: 216
-- Name: item_buyer_id_seq; Type: SEQUENCE SET; Schema: public; Owner: fismed-user
--

SELECT pg_catalog.setval('public.item_buyer_id_seq', 57, true);


--
-- TOC entry 3557 (class 0 OID 0)
-- Dependencies: 218
-- Name: order_items_id_seq; Type: SEQUENCE SET; Schema: public; Owner: fismed-user
--

SELECT pg_catalog.setval('public.order_items_id_seq', 74, true);


--
-- TOC entry 3558 (class 0 OID 0)
-- Dependencies: 220
-- Name: pemasukan_id_seq; Type: SEQUENCE SET; Schema: public; Owner: fismed-user
--

SELECT pg_catalog.setval('public.pemasukan_id_seq', 28, true);


--
-- TOC entry 3559 (class 0 OID 0)
-- Dependencies: 222
-- Name: pengeluaran_id_seq; Type: SEQUENCE SET; Schema: public; Owner: fismed-user
--

SELECT pg_catalog.setval('public.pengeluaran_id_seq', 16, true);


--
-- TOC entry 3560 (class 0 OID 0)
-- Dependencies: 224
-- Name: performance_invoice_id_seq; Type: SEQUENCE SET; Schema: public; Owner: fismed-user
--

SELECT pg_catalog.setval('public.performance_invoice_id_seq', 48, true);


--
-- TOC entry 3561 (class 0 OID 0)
-- Dependencies: 226
-- Name: piutang_id_seq; Type: SEQUENCE SET; Schema: public; Owner: fismed-user
--

SELECT pg_catalog.setval('public.piutang_id_seq', 29, true);


--
-- TOC entry 3562 (class 0 OID 0)
-- Dependencies: 228
-- Name: purchase_order_id_seq; Type: SEQUENCE SET; Schema: public; Owner: fismed-user
--

SELECT pg_catalog.setval('public.purchase_order_id_seq', 42, true);


--
-- TOC entry 3563 (class 0 OID 0)
-- Dependencies: 230
-- Name: stock_items_id_seq; Type: SEQUENCE SET; Schema: public; Owner: fismed-user
--

SELECT pg_catalog.setval('public.stock_items_id_seq', 15, true);


--
-- TOC entry 3564 (class 0 OID 0)
-- Dependencies: 232
-- Name: tax_code_id_seq; Type: SEQUENCE SET; Schema: public; Owner: fismed-user
--

SELECT pg_catalog.setval('public.tax_code_id_seq', 3, true);


--
-- TOC entry 3565 (class 0 OID 0)
-- Dependencies: 234
-- Name: user_category_id_seq; Type: SEQUENCE SET; Schema: public; Owner: fismed-user
--

SELECT pg_catalog.setval('public.user_category_id_seq', 5, true);


--
-- TOC entry 3566 (class 0 OID 0)
-- Dependencies: 236
-- Name: user_id_seq; Type: SEQUENCE SET; Schema: public; Owner: fismed-user
--

SELECT pg_catalog.setval('public.user_id_seq', 1, true);


--
-- TOC entry 3325 (class 2606 OID 16669)
-- Name: customer customer_pkey; Type: CONSTRAINT; Schema: public; Owner: fismed-user
--

ALTER TABLE ONLY public.customer
    ADD CONSTRAINT customer_pkey PRIMARY KEY (id);


--
-- TOC entry 3327 (class 2606 OID 16671)
-- Name: divisi divisi_pkey; Type: CONSTRAINT; Schema: public; Owner: fismed-user
--

ALTER TABLE ONLY public.divisi
    ADD CONSTRAINT divisi_pkey PRIMARY KEY (id);


--
-- TOC entry 3329 (class 2606 OID 16673)
-- Name: galery galery_pkey; Type: CONSTRAINT; Schema: public; Owner: fismed-user
--

ALTER TABLE ONLY public.galery
    ADD CONSTRAINT galery_pkey PRIMARY KEY (id);


--
-- TOC entry 3355 (class 2606 OID 16746)
-- Name: hutang hutang_pkey; Type: CONSTRAINT; Schema: public; Owner: fismed-user
--

ALTER TABLE ONLY public.hutang
    ADD CONSTRAINT hutang_pkey PRIMARY KEY (id);


--
-- TOC entry 3331 (class 2606 OID 16677)
-- Name: item_buyer item_buyer_pkey; Type: CONSTRAINT; Schema: public; Owner: fismed-user
--

ALTER TABLE ONLY public.item_buyer
    ADD CONSTRAINT item_buyer_pkey PRIMARY KEY (id);


--
-- TOC entry 3333 (class 2606 OID 16679)
-- Name: order_items order_items_pkey; Type: CONSTRAINT; Schema: public; Owner: fismed-user
--

ALTER TABLE ONLY public.order_items
    ADD CONSTRAINT order_items_pkey PRIMARY KEY (id);


--
-- TOC entry 3335 (class 2606 OID 16681)
-- Name: pemasukan pemasukan_pkey; Type: CONSTRAINT; Schema: public; Owner: fismed-user
--

ALTER TABLE ONLY public.pemasukan
    ADD CONSTRAINT pemasukan_pkey PRIMARY KEY (id);


--
-- TOC entry 3337 (class 2606 OID 16683)
-- Name: pengeluaran pengeluaran_pkey; Type: CONSTRAINT; Schema: public; Owner: fismed-user
--

ALTER TABLE ONLY public.pengeluaran
    ADD CONSTRAINT pengeluaran_pkey PRIMARY KEY (id);


--
-- TOC entry 3339 (class 2606 OID 16685)
-- Name: performance_invoice performance_invoice_pkey; Type: CONSTRAINT; Schema: public; Owner: fismed-user
--

ALTER TABLE ONLY public.performance_invoice
    ADD CONSTRAINT performance_invoice_pkey PRIMARY KEY (id);


--
-- TOC entry 3341 (class 2606 OID 16687)
-- Name: piutang piutang_pkey; Type: CONSTRAINT; Schema: public; Owner: fismed-user
--

ALTER TABLE ONLY public.piutang
    ADD CONSTRAINT piutang_pkey PRIMARY KEY (id);


--
-- TOC entry 3343 (class 2606 OID 16689)
-- Name: purchase_order purchase_order_pkey; Type: CONSTRAINT; Schema: public; Owner: fismed-user
--

ALTER TABLE ONLY public.purchase_order
    ADD CONSTRAINT purchase_order_pkey PRIMARY KEY (id);


--
-- TOC entry 3345 (class 2606 OID 16691)
-- Name: stock_items stock_items_pkey; Type: CONSTRAINT; Schema: public; Owner: fismed-user
--

ALTER TABLE ONLY public.stock_items
    ADD CONSTRAINT stock_items_pkey PRIMARY KEY (id);


--
-- TOC entry 3347 (class 2606 OID 16693)
-- Name: tax_code tax_code_pkey; Type: CONSTRAINT; Schema: public; Owner: fismed-user
--

ALTER TABLE ONLY public.tax_code
    ADD CONSTRAINT tax_code_pkey PRIMARY KEY (id);


--
-- TOC entry 3349 (class 2606 OID 16695)
-- Name: user_category user_category_pkey; Type: CONSTRAINT; Schema: public; Owner: fismed-user
--

ALTER TABLE ONLY public.user_category
    ADD CONSTRAINT user_category_pkey PRIMARY KEY (id);


--
-- TOC entry 3351 (class 2606 OID 16697)
-- Name: users user_pkey; Type: CONSTRAINT; Schema: public; Owner: fismed-user
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT user_pkey PRIMARY KEY (id);


--
-- TOC entry 3353 (class 2606 OID 16699)
-- Name: users user_username_key; Type: CONSTRAINT; Schema: public; Owner: fismed-user
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT user_username_key UNIQUE (username);


--
-- TOC entry 3356 (class 2606 OID 16700)
-- Name: customer customer_tax_code_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: fismed-user
--

ALTER TABLE ONLY public.customer
    ADD CONSTRAINT customer_tax_code_id_fkey FOREIGN KEY (tax_code_id) REFERENCES public.tax_code(id);


--
-- TOC entry 3359 (class 2606 OID 16705)
-- Name: performance_invoice fk_customer; Type: FK CONSTRAINT; Schema: public; Owner: fismed-user
--

ALTER TABLE ONLY public.performance_invoice
    ADD CONSTRAINT fk_customer FOREIGN KEY (customer_id) REFERENCES public.customer(id);


--
-- TOC entry 3357 (class 2606 OID 16710)
-- Name: item_buyer item_buyer_po_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: fismed-user
--

ALTER TABLE ONLY public.item_buyer
    ADD CONSTRAINT item_buyer_po_id_fkey FOREIGN KEY (po_id) REFERENCES public.purchase_order(id);


--
-- TOC entry 3358 (class 2606 OID 16715)
-- Name: order_items order_items_pi_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: fismed-user
--

ALTER TABLE ONLY public.order_items
    ADD CONSTRAINT order_items_pi_id_fkey FOREIGN KEY (pi_id) REFERENCES public.performance_invoice(id);


--
-- TOC entry 3360 (class 2606 OID 16720)
-- Name: users user_role_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: fismed-user
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT user_role_id_fkey FOREIGN KEY (role_id) REFERENCES public.user_category(id);


--
-- TOC entry 3536 (class 0 OID 0)
-- Dependencies: 5
-- Name: SCHEMA public; Type: ACL; Schema: -; Owner: pg_database_owner
--

REVOKE USAGE ON SCHEMA public FROM PUBLIC;


-- Completed on 2024-08-26 10:25:50

--
-- PostgreSQL database dump complete
--

