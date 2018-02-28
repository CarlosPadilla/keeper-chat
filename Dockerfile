FROM golang

WORKDIR /opt/out
COPY out .
COPY wait-for-it.sh .
RUN chmod +x wait-for-it.sh

CMD ["./wait-for-it.sh", "db:5432", "--", "./run.sh"]