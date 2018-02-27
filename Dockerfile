FROM golang

WORKDIR /go/src/app
COPY api .

RUN go get -d -v
RUN go install -v

CMD ["app"]