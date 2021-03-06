#!/bin/sh

set -e

# to make sure etcd is ready (election ended and leader elected)
while ! etcdctl endpoint health &>/dev/null; do :; done

exec etcdctl put /config/go_oauth2_server.json '{
  "Database": {
    "Type": "postgres",
    "Host": "db",
    "Port": 5432,
    "User": "oauth2",
    "Password": "",
    "DatabaseName": "oauth2",
    "MaxIdleConns": 5,
    "MaxOpenConns": 5
  },
  "Oauth": {
    "AccessTokenLifetime": 3600,
    "RefreshTokenLifetime": 1209600,
    "AuthCodeLifetime": 3600
  },
  "Session": {
      "Secret": "test_secret",
      "Path": "/",
      "MaxAge": 604800,
      "HTTPOnly": true
  },
  "IsDevelopment": true
}'

exec etcdctl put /config/go_api_server.json '{
  "Database": {
    "Type": "postgres",
    "Host": "db",
    "Port": 5432,
    "User": "keeper",
    "Password": "",
    "DatabaseName": "keeper",
    "MaxIdleConns": 5,
    "MaxOpenConns": 5
  },
  "IsDevelopment": true
}'
