#!/usr/bin/env bash

ENVIRONMENT=dev

revel build github.com/Zeloid/keeper-chat/api out ${ENVIRONMENT}
docker-compose build
