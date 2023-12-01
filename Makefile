start:
	docker compose up -d
stop:
	docker compose down
dev:
	go run cmd/api/main.go


