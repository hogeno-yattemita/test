up:
	docker compose up
down:
	docker compose down 
exec:
	docker exec -it go-api-test-go-1 ash
build:
	docker compose build
go-build:
	cd app && go build -v ./...
test:
	cd app && go test -v ./...
