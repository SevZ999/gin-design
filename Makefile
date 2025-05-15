run:
	go run ./cmd/main.go

git:
	@echo "Staging all changes..."
	git add .
	@read -p "Enter commit message: " msg; \
	git commit -m "$$msg"
	@echo "Commit completed!"