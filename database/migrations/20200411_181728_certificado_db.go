package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CertificadoDb_20200411_181728 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CertificadoDb_20200411_181728{}
	m.Created = "20200411_181728"

	migration.Register("CertificadoDb_20200411_181728", m)
}

// Run the migrations
func (m *CertificadoDb_20200411_181728) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE SEQUENCE public.id_certificado_seq START WITH 1 INCREMENT BY 1 NO MINVALUE MAXVALUE 123456789 CACHE 1;")
	m.SQL("CREATE TABLE public.certificado ( id_certificado integer DEFAULT nextval('public.id_certificado_seq'::regclass) NOT NULL, texto text NOT NULL, id_usuario integer DEFAULT 1 NOT NULL, id_curso integer NOT NULL );")
	m.SQL("CREATE SEQUENCE public.id_datos_usuario_curso_seq START WITH 1 INCREMENT BY 1 NO MINVALUE MAXVALUE 123456789 CACHE 1;")
	m.SQL("CREATE TABLE public.datos_usuario_curso ( id_datos_usuario_curso integer DEFAULT nextval('public.id_datos_usuario_curso_seq'::regclass) NOT NULL, id_usuario integer NOT NULL, id_curso integer NOT NULL, nombre_curso character varying NOT NULL, nombre_usuario character varying NOT NULL, documento_usuario integer NOT NULL );")
	m.SQL("CREATE INDEX fki_id_datos_usuario_curso_fk ON public.certificado USING btree (id_usuario);")
	m.SQL("INSERT INTO public.datos_usuario_curso (id_usuario, id_curso, nombre_curso, nombre_usuario, documento_usuario) VALUES (1, 1, 'GO', 'Diego', 12345);")
	m.SQL("INSERT INTO public.datos_usuario_curso (id_usuario, id_curso, nombre_curso, nombre_usuario, documento_usuario) VALUES (2, 1, 'GO', 'Miguel', 54321);")
}

// Reverse the migrations
func (m *CertificadoDb_20200411_181728) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP SEQUENCE if exists public.id_datos_usuario_curso_seq CASCADE")
	m.SQL("DROP SEQUENCE if exists public.id_certificado_seq CASCADE")
	m.SQL("DROP TABLE if exists public.certificado CASCADE;")
	m.SQL("DROP TABLE if exists public.datos_usuario_curso CASCADE;")
}
