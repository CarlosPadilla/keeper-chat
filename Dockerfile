FROM golang

COPY out/ ./

RUN apt-get update -q
RUN apt-get install postgresql-client -q -y

COPY wait-for-postgres.sh .
RUN chmod +x wait-for-postgres.sh

COPY api-entrypoint.sh .
RUN chmod +x api-entrypoint.sh

RUN go get -u -d github.com/mattes/migrate/cli github.com/lib/pq
RUN go build -tags 'postgres' -o /usr/local/bin/migrate github.com/mattes/migrate/cli

COPY fixtures/* fixtures/

CMD ["./wait-for-postgres.sh", "./api-entrypoint.sh"]