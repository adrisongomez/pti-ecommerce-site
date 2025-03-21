.ONESHELL:

generate-clients:
	go run github.com/steebchen/prisma-client-go generate --schema ./databases/schema.prisma

db-push:
	go run github.com/steebchen/prisma-client-go push --schema ./databases/schema.prisma

db-pull:
	go run github.com/steebchen/prisma-client-go pull --schema ./databases/schema.prisma

migrate-dev:
	go run github.com/steebchen/prisma-client-go migrate dev --schema ./databases/schema.prisma

migrate-depoy:
	go run github.com/steebchen/prisma-client-go migrate deploy --schema ./databases/schema.prisma
