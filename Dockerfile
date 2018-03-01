FROM golang

ADD out .

COPY wait-for-it.sh .
RUN chmod +x wait-for-it.sh

COPY api-entrypoint.sh .
RUN chmod +x api-entrypoint.sh

RUN go get -u -d github.com/mattes/migrate/cli github.com/lib/pq
RUN go build -tags 'postgres' -o /usr/local/bin/migrate github.com/mattes/migrate/cli

COPY fixtures/* fixtures/

CMD ["./wait-for-it.sh", "db:5432", "--", "./api-entrypoint.sh"]