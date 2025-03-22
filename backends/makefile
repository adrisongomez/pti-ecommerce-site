.ONESHELL:

start-server:
	go run ./cmd/ecommerce/main.go

generate-clients:
	go run github.com/steebchen/prisma-client-go generate --schema ./databases/schema.prisma

db-push:
	go run github.com/steebchen/prisma-client-go db push --schema ./databases/schema.prisma

db-pull:
	go run github.com/steebchen/prisma-client-go db pull --schema ./databases/schema.prisma

migrate-dev:
	go run github.com/steebchen/prisma-client-go migrate dev --schema ./databases/schema.prisma

migrate-depoy:
	go run github.com/steebchen/prisma-client-go migrate deploy --schema ./databases/schema.prisma

generate-svc:
	goa gen github.com/adrisongomez/pti-ecommerce-site/backends/design -o ./internal

format-prisma:
	go run github.com/steebchen/prisma-client-go format --schema ./databases/schema.prisma
