run: 
	@templ generate
	@go run cmd/server/main.go

docker: 
	@docker compose up --build -d
