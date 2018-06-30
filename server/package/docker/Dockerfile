#FROM scratch
FROM centos

RUN mkdir -p /opt/promise
RUN touch /tmp/promise.log
WORKDIR /opt/promise
COPY server .
COPY rpm/conf ./conf
COPY rpm/static ./static
#ADD centos-7-docker.tar.xz /

# LABEL org.label-schema.schema-verison = "1.0" \
#     org.label-schema.name="CentOS Base Image" \
#     org.label-schema.vendor="CentOS" \
#     org.label-schema.license="GPLv2"

EXPOSE 80

# ENTRYPOINT ["/bin/bash"]
ENTRYPOINT ["./server"]