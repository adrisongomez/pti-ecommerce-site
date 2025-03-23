.ONESHELL:

start-server:
	go run ./cmd/ecommerce/main.go

generate-clients:
	go run github.com/steebchen/prisma-client-go generate --schema ./backends/databases/schema.prisma

db-push:
	go run github.com/steebchen/prisma-client-go db push --schema ./backends/databases/schema.prisma

db-pull:
	go run github.com/steebchen/prisma-client-go db pull --schema ./backends/databases/schema.prisma

migrate-dev:
	go run github.com/steebchen/prisma-client-go migrate dev --schema ./backends/databases/schema.prisma

migrate-depoy:
	go run github.com/steebchen/prisma-client-go migrate deploy --schema ./backends/databases/schema.prisma

generate-svc:
	goa gen github.com/adrisongomez/pti-ecommerce-site/backends/design -o ./backends//internal

format-prisma:
	go run github.com/steebchen/prisma-client-go format --schema ./backends/databases/schema.prisma

start-external-svc:
	docker-compose -f ./docker/docker-compose.yml up -d

stop-docker:
	docker-compose -f ./docker/docker-compose.yml stop
