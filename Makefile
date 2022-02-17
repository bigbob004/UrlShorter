run_in_memory:
	go run cmd/main.go -InMemory

docker_and_migrate:
	docker run --name=urls -e POSTGRES_PASSWORD='qwerty' -p 5432:5432 -d --rm postgres
	sleep 2
	migrate -path ./schema -database 'postgres://postgres:qwerty@localhost:5432/postgres?sslmode=disable' up

run_in_db:
	go run cmd/main.go -InDb

