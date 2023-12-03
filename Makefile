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

dev:
	go run cmd/api/main.go


