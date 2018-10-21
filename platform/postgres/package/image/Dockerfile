FROM postgres:latest

COPY ./create-multiple-postgresql-databases.sh /docker-entrypoint-initdb.d/create-multiple-postgresql-databases.sh
RUN chmod a+x /docker-entrypoint-initdb.d/create-multiple-postgresql-databases.sh