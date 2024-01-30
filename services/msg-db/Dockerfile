FROM postgres

COPY msg-db-schema.sql /docker-entrypoint-initdb.d/
COPY users-db-init.sh /docker-entrypoint-initdb.d/