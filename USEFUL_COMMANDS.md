# Docker commands


docker run -it --rm --link keeperchat_db_1:postgres --net keeperchat_default postgres psql -U keeper -h postgres
