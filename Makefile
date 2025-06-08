run:
	go run ./cmd/main.go

git:
	@echo "Staging all changes..."
	git add .
	@read -p "Enter commit message: " msg; \
	git commit -m "$$msg"
	@echo "Commit completed!"

build:
	@echo "Building..."
	@read -p "Enter build version: " version; \
	docker build -t crpi-nmbc0yqbbslptbxi.cn-shanghai.personal.cr.aliyuncs.com/sevz98/gin-design:v"$$version" .
push:
	@read -p "Enter build version: " version; \
	docker push crpi-nmbc0yqbbslptbxi.cn-shanghai.personal.cr.aliyuncs.com/sevz98/gin-design:v"$$version"
	@echo "Push completed!"



start:
	docker-compose --compatibility up -d

wire:
	wire ./...

swag:
	swag init -g ./cmd/main.go


.PHONY: gen del
gen:
	@echo "正在准备生成代码..."
	@read -p "请输入模块名（小写字母）: " module_name; \
	if [ -z "$$module_name" ]; then \
		echo "错误：模块名不能为空"; \
		exit 1; \
	elif ! echo "$$module_name" | grep -qE '^[a-z]+$$'; then \
		echo "错误：模块名必须为小写字母"; \
		exit 1; \
	fi; \
	echo "正在生成模块：$$module_name"; \
	./generate.sh $$module_name

del:
	@echo "正在准备删除代码..."
	@if [ -z "$(MODULE)" ]; then \
		read -p "请输入要删除的模块名（小写字母）: " MODULE; \
	fi; \
	if [ -z "$$MODULE" ]; then \
		echo "错误：模块名不能为空"; \
		exit 1; \
	elif ! echo "$$MODULE" | grep -qE '^[a-z]+$$'; then \
		echo "错误：模块名必须为小写字母"; \
		exit 1; \
	fi; \
	echo "正在删除模块：$$MODULE"; \
	rm -f ./internal/app/controller/$$MODULE.go; \
	rm -f ./internal/app/service/$$MODULE.go; \
	rm -f ./internal/app/repo/$$MODULE.go; \
	rm -f ./internal/app/router/$$MODULE.go; \
	echo "删除完成！"