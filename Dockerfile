FROM mysql:8.0.33
COPY ./db /docker-entrypoint-initdb.d/

