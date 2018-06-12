FROM scratch

ADD centos-7-docker.tar.xz /

RUN touch /tmp/promise.log
RUN mkdir -p /opt/promise/apps/server/conf
COPY ./rpm/* /opt/promise/apps/server/

LABEL org.label-schema.schema-verison = "1.0" \
    org.label-schema.name="CentOS Base Image" \
    org.label-schema.vendor="CentOS" \
    org.label-schema.license="GPLv2"

EXPOSE 80

ENTRYPOINT ["/bin/bash"]