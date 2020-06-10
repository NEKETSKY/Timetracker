start:
		docker-compose build
		docker-compose up -d

stop:
		docker-compose down

db_update:
		docker-compose run --rm migrator update
		docker-compose run --rm migrator changelogSync

restart:
		docker-compose down
		docker-compose build
		docker-compose up -d