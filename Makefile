all:
	@echo "Invalid make cmd. Options available: ['web69-up', 'web69-down', 'web70-up', 'web70-down', 'web71-up', 'web71-down', 'lb-up', 'lb-down', 'docker-up', 'docker-down']"

web69-up:
	@docker compose up web69 --build

web69-down:
	@docker compose down web69

web70-up:
	@docker compose up web70 --build

web70-down:
	@docker compose down web70

web71-up:
	@docker compose up web71 --build

web71-down:
	@docker compose down web71

lb-up:
	@docker compose up lb --build

lb-down:
	@docker compose down lb

docker-up:
	@docker compose up --build

docker-down:
	@docker compose down
