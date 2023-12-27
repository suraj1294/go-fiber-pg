start-dev:
	 docker compose -f docker-compose.dev.yml  up -d --build

stop-dev:
	 docker compose -f docker-compose.dev.yml down

start-db:
	 docker compose -f docker-compose.db.yml  up -d --build

stop-db:
	docker compose -f docker-compose.db.yml down

docker-build:
	docker build -t server . --no-cache

docker-run:
	docker run -p 8080:8080 --env-file .env  -d server

docker-exec:
	docker exec -it $(id) sh 
#make id=ed docker-exec

docker-dev-build:
	docker-compose -f docker-compose.dev.yml build

dev:
	go run cmd/api/main.go


