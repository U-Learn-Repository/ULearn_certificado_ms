FROM postgres:12
MAINTAINER skoltun
ENV POSTGRES_HOST=ULearn_certificado_db
ENV POSTGRES_DB=ULearn_certificado_db
ENV POSTGRES_USER=certificado123
ENV POSTGRES_PASSWORD=certificado123
ENV POSTGRES_HOST_AUTH_METHOD=trust
COPY schema.sql /docker-entrypoint-initdb.d/