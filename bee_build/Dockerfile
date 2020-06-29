FROM golang:1.14.1-alpine

RUN apk add --no-cache git
RUN apk add --no-cache build-base


RUN set -x \
    # go get bee
    && export GO111MODULE=on \
    && go get github.com/astaxie/beego \
    && go get -u github.com/beego/bee 

WORKDIR /go/src


# #CMD cd /go/src/github.com/diagutierrezro/ULearn_certificado_ms
# CMD go get -v -u ./...
# CMD bee run -downdoc=true -gendoc=true