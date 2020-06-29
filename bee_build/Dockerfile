FROM golang:1.14.1-alpine

RUN apk add --no-cache git
RUN apk add --no-cache build-base


ENV CERTIFICADO_CRUD_HTTP_PORT=6671
ENV API_NAME=ULearn_certificado_ms
ENV CERTIFICADO_CRUD_DB_USER=certificado123
ENV CERTIFICADO_CRUD_DB_PASSWORD=certificado123
ENV CERTIFICADO_CRUD_DB_URL=ULearn_certificado_db
ENV CERTIFICADO_CRUD_DB_NAME=ULearn_certificado_db
ENV RUN_MODE=dev

WORKDIR /go/src

RUN set -x \
    # go get bee
    && export GO111MODULE=on \
    && go get github.com/astaxie/beego \
    && go get -u github.com/beego/bee 



#CMD cd /go/src/github.com/diagutierrezro/ULearn_certificado_ms
CMD go get -v -u ./...
CMD bee run -downdoc=true -gendoc=true