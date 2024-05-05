--
-- PostgreSQL database dump
--

-- Dumped from database version 14.7 (Homebrew)
-- Dumped by pg_dump version 14.7 (Homebrew)

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
-- Name: update_timestamp_businesses(); Type: FUNCTION; Schema: public; Owner: orololuwa
--

CREATE FUNCTION public.update_timestamp_businesses() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$;


ALTER FUNCTION public.update_timestamp_businesses() OWNER TO orololuwa;

--
-- Name: update_timestamp_kyc(); Type: FUNCTION; Schema: public; Owner: orololuwa
--

CREATE FUNCTION public.update_timestamp_kyc() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$;


ALTER FUNCTION public.update_timestamp_kyc() OWNER TO orololuwa;

--
-- Name: update_timestamp_users(); Type: FUNCTION; Schema: public; Owner: orololuwa
--

CREATE FUNCTION public.update_timestamp_users() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$;


ALTER FUNCTION public.update_timestamp_users() OWNER TO orololuwa;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: businesses; Type: TABLE; Schema: public; Owner: orololuwa
--

CREATE TABLE public.businesses (
    id integer NOT NULL,
    name character varying(255) NOT NULL,
    email character varying(255) NOT NULL,
    description character varying(255),
    sector character varying(255) NOT NULL,
    is_corporate_affairs boolean DEFAULT false NOT NULL,
    logo character varying(255),
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    user_id integer
);


ALTER TABLE public.businesses OWNER TO orololuwa;

--
-- Name: businesses_id_seq; Type: SEQUENCE; Schema: public; Owner: orololuwa
--

CREATE SEQUENCE public.businesses_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.businesses_id_seq OWNER TO orololuwa;

--
-- Name: businesses_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: orololuwa
--

ALTER SEQUENCE public.businesses_id_seq OWNED BY public.businesses.id;


--
-- Name: kyc; Type: TABLE; Schema: public; Owner: orololuwa
--

CREATE TABLE public.kyc (
    id integer NOT NULL,
    certificate_of_registration character varying(255),
    proof_of_address character varying(255),
    bvn character varying(255),
    business_address text,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    business_id integer
);


ALTER TABLE public.kyc OWNER TO orololuwa;

--
-- Name: kyc_id_seq; Type: SEQUENCE; Schema: public; Owner: orololuwa
--

CREATE SEQUENCE public.kyc_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.kyc_id_seq OWNER TO orololuwa;

--
-- Name: kyc_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: orololuwa
--

ALTER SEQUENCE public.kyc_id_seq OWNED BY public.kyc.id;


--
-- Name: schema_migration; Type: TABLE; Schema: public; Owner: orololuwa
--

CREATE TABLE public.schema_migration (
    version character varying(14) NOT NULL
);


ALTER TABLE public.schema_migration OWNER TO orololuwa;

--
-- Name: users; Type: TABLE; Schema: public; Owner: orololuwa
--

CREATE TABLE public.users (
    id integer NOT NULL,
    email character varying(255) NOT NULL,
    first_name character varying(255) NOT NULL,
    last_name character varying(255) NOT NULL,
    phone character varying(255) NOT NULL,
    password character varying(60),
    avatar character varying(255),
    gender character varying(255),
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL
);


ALTER TABLE public.users OWNER TO orololuwa;

--
-- Name: users_id_seq; Type: SEQUENCE; Schema: public; Owner: orololuwa
--

CREATE SEQUENCE public.users_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.users_id_seq OWNER TO orololuwa;

--
-- Name: users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: orololuwa
--

ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;


--
-- Name: businesses id; Type: DEFAULT; Schema: public; Owner: orololuwa
--

ALTER TABLE ONLY public.businesses ALTER COLUMN id SET DEFAULT nextval('public.businesses_id_seq'::regclass);


--
-- Name: kyc id; Type: DEFAULT; Schema: public; Owner: orololuwa
--

ALTER TABLE ONLY public.kyc ALTER COLUMN id SET DEFAULT nextval('public.kyc_id_seq'::regclass);


--
-- Name: users id; Type: DEFAULT; Schema: public; Owner: orololuwa
--

ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);


--
-- Name: businesses businesses_pkey; Type: CONSTRAINT; Schema: public; Owner: orololuwa
--

ALTER TABLE ONLY public.businesses
    ADD CONSTRAINT businesses_pkey PRIMARY KEY (id);


--
-- Name: businesses businesses_user_id_key; Type: CONSTRAINT; Schema: public; Owner: orololuwa
--

ALTER TABLE ONLY public.businesses
    ADD CONSTRAINT businesses_user_id_key UNIQUE (user_id);


--
-- Name: kyc kyc_pkey; Type: CONSTRAINT; Schema: public; Owner: orololuwa
--

ALTER TABLE ONLY public.kyc
    ADD CONSTRAINT kyc_pkey PRIMARY KEY (id);


--
-- Name: schema_migration schema_migration_pkey; Type: CONSTRAINT; Schema: public; Owner: orololuwa
--

ALTER TABLE ONLY public.schema_migration
    ADD CONSTRAINT schema_migration_pkey PRIMARY KEY (version);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: orololuwa
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- Name: businesses_email_idx; Type: INDEX; Schema: public; Owner: orololuwa
--

CREATE UNIQUE INDEX businesses_email_idx ON public.businesses USING btree (email);


--
-- Name: schema_migration_version_idx; Type: INDEX; Schema: public; Owner: orololuwa
--

CREATE UNIQUE INDEX schema_migration_version_idx ON public.schema_migration USING btree (version);


--
-- Name: users_email_idx; Type: INDEX; Schema: public; Owner: orololuwa
--

CREATE UNIQUE INDEX users_email_idx ON public.users USING btree (email);


--
-- Name: users_phone_idx; Type: INDEX; Schema: public; Owner: orololuwa
--

CREATE UNIQUE INDEX users_phone_idx ON public.users USING btree (phone);


--
-- Name: businesses update_timestamp_businesses_trigger; Type: TRIGGER; Schema: public; Owner: orololuwa
--

CREATE TRIGGER update_timestamp_businesses_trigger BEFORE UPDATE ON public.businesses FOR EACH ROW WHEN ((new.updated_at IS DISTINCT FROM old.updated_at)) EXECUTE FUNCTION public.update_timestamp_businesses();


--
-- Name: users update_timestamp_businesses_trigger; Type: TRIGGER; Schema: public; Owner: orololuwa
--

CREATE TRIGGER update_timestamp_businesses_trigger BEFORE UPDATE ON public.users FOR EACH ROW WHEN ((new.updated_at IS DISTINCT FROM old.updated_at)) EXECUTE FUNCTION public.update_timestamp_users();


--
-- Name: kyc update_timestamp_kyc_trigger; Type: TRIGGER; Schema: public; Owner: orololuwa
--

CREATE TRIGGER update_timestamp_kyc_trigger BEFORE UPDATE ON public.kyc FOR EACH ROW WHEN ((new.updated_at IS DISTINCT FROM old.updated_at)) EXECUTE FUNCTION public.update_timestamp_kyc();


--
-- Name: kyc fk_business_id; Type: FK CONSTRAINT; Schema: public; Owner: orololuwa
--

ALTER TABLE ONLY public.kyc
    ADD CONSTRAINT fk_business_id FOREIGN KEY (business_id) REFERENCES public.businesses(id) ON DELETE CASCADE;


--
-- Name: businesses fk_user_id; Type: FK CONSTRAINT; Schema: public; Owner: orololuwa
--

ALTER TABLE ONLY public.businesses
    ADD CONSTRAINT fk_user_id FOREIGN KEY (user_id) REFERENCES public.users(id) ON DELETE CASCADE;


--
-- PostgreSQL database dump complete
--

