FROM postgres
LABEL maintainer="tommychu2256@gmail.com"

ENV POSTGRES_PASSWORD='postgres'
ENV POSTGRES_USER='postgres'
ENV POSTGRES_DB='word_list'

ENV PORT=5432

EXPOSE $PORT
COPY ./word_list.sql /docker-entrypoint-initdb.d/word_list.sql
