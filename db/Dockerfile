FROM postgres:13.3

COPY ./db/scripts/bootstrap_db.sh /docker-entrypoint-initdb.d
COPY ./db/ddl /ddl
