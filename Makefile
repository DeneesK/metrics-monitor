.PHONY: start_monitor

start_monitor:
	docker compose -f docker_compose.yml up -d --build
