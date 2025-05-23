run:
	go run ./cmd/main.go

git:
	@echo "Staging all changes..."
	git add .
	@read -p "Enter commit message: " msg; \
	git commit -m "$$msg"
	@echo "Commit completed!"

build:
	docker build -t loan-admin .

start:
	docker compose up -d

wire:
	wire ./...