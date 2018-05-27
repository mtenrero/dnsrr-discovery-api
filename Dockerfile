FROM alpine

LABEL MAINTAINER="Marcos Tenrero"

COPY ./releases/ /dnsrr

ENTRYPOINT [ "./dnsrr/dnsrrResolver-linux-amd64" ]