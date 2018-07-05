FROM centos

LABEL MAINTAINER="Marcos Tenrero"

COPY ./releases/atq-director-linux-amd64 /atq/atq-amd64
COPY controller-config.docker.yaml /controller-config.yaml

RUN mkdir -p /gluster/storage

EXPOSE 8080

RUN cd /atq

ENTRYPOINT [ "/atq/atq-amd64" ]