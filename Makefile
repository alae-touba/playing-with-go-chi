# Start the application
up:
	docker-compose --env-file .env -f docker/docker-compose.yml up

# Start the application in detached mode
up-d:
	docker-compose --env-file .env -f docker/docker-compose.yml up -d

# Stop and remove all containers and volumes
down:
	docker-compose --env-file .env -f docker/docker-compose.yml down -v

# Reset the entire environment
reset: down
	docker-compose --env-file .env -f docker/docker-compose.yml up --build

# generate ent files
ent:
	go generate ./repositories

hurl:
	hurl api-tests.hurl