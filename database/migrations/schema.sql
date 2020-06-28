CREATE SEQUENCE id_certificado_seq START WITH 1 INCREMENT BY 1 NO MINVALUE MAXVALUE 123456789 CACHE 1;
CREATE TABLE IF NOT EXISTS certificado ( 
    id_certificado integer DEFAULT nextval('public.id_certificado_seq'::regclass) PRIMARY KEY, 
    texto text NOT NULL, 
    id_usuario integer DEFAULT 1 NOT NULL,
    id_curso integer NOT NULL );
