all: docker start

project start:
	@docker-compose up -d
	@sh migration.sh

migration start:
	@sh migration.sh