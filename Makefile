# Start the application
up: 
	docker-compose -f docker/docker-compose.yml up


# Start the application in detached mode
up-d:
	 docker-compose -f docker/docker-compose.yml up -d