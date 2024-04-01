all:
	@echo "Invalid make cmd. Options available: ['docker-up', 'docker-down', 'ping']"

docker-up:
	@docker compose up --build

docker-down:
	@docker compose down

ping:
	@curl http://localhost:7070/ping
