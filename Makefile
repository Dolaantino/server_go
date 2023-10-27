docker-compose:
				docker compose up -d
main:
	go run main.go
api:
	go build -v ./apiserver