FROM postgres:13.3

COPY ./scripts/bootstrap_db.sh /docker-entrypoint-initdb.d
COPY ./ddl /ddl
