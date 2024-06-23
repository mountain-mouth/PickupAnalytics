all: docker start

docker start:
	@docker-compose up -d
	@sh migration.sh

migration start:
	@sh migration.sh