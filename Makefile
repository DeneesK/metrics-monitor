.PHONY: start_monitor

start_monitor:
	docker stop $(docker ps -q) && docker compose -f docker_compose.yml up -d --build
