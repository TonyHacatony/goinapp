run:
	cd goapp && go run main.go

dev:
	@echo "Starting development server with live reload..."
	cd goapp && air

build:
	docker-compose up --build -d

re-build:
	docker-compose down -v
	docker-compose up --build -d

clean:
	docker-compose down -v
	docker system prune -f