run:
	go run cmd/main.go

proto_gen:
	go generate gen.go

up:
	docker-compose -f docker-compose.yml up -d --build

down:
	docker-compose -f docker-compose.yml down
