#-------------------------------SYSTEM-----------------------------------
include .env
export
#-------------------------------------------------------------------------

#-------------------------------Development-------------------------------
reboot: env-init docker-down-clear docker-up

restart: docker-down docker-up

env-init:
	rm -f .env && cp -n .env.example .env

docker-down-clear:
	docker-compose down -v --remove-orphans

docker-down:
	docker-compose down

docker-up:
	docker-compose up -d --build --remove-orphans

sh:
	docker exec -it SERVER_GO sh

migration:
	docker-compose run --rm migration goose postgres "$(DB_STR)" $(filter-out $@,$(MAKECMDGOALS))
#-------------------------------------------------------------------------
