.PHONY: build up down clean

build:
	@echo "Building Docker images..."
	docker-compose build

up:
	@echo "Starting services..."
	docker-compose up -d

down:
	@echo "Stopping services..."
	docker-compose down

clean: down
	@echo "Removing volumes..."
	docker-compose down -v
