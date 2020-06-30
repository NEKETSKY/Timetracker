.PHONY: start stop db_update restart all
all: start db_update
start:
		docker-compose build
		docker-compose up -d

stop:
		docker-compose down

db_update:
		docker-compose run --rm migrator update
		docker-compose run --rm migrator changelogSync

restart: stop start