ifneq (,$(wildcard ./.env))
    include .env
    export
endif

build: gen 
	go build -o ./bin/app.exe ./cmd/app

run: build
	./bin/app

PB=pb
AUTHPB=$(PB)/authpb
AUTHPROTO=proto/auth.proto
SBPB=$(PB)/sspb
SBPROTO=proto/subscriptions-service.proto
ESPB=$(PB)/espb
ESPROTO=proto/event-service.proto
NSPB=$(PB)/nspb
NSPROTO=proto/notifications-service.proto
protobuf:
	cp -r ../proto .
	protoc  \
		--go_out=. \
		--go_opt=M$(NSPROTO)=$(NSPB) \
		--go-grpc_out=. \
		--go-grpc_opt=M$(NSPROTO)=$(NSPB) \
		$(NSPROTO)

	protoc  \
		--go_out=. \
		--go_opt=M$(AUTHPROTO)=$(AUTHPB) \
		--go-grpc_out=. \
		--go-grpc_opt=M$(AUTHPROTO)=$(AUTHPB) \
		$(AUTHPROTO)

	protoc  \
		--go_out=. \
		--go_opt=M$(SBPROTO)=$(SBPB) \
		--go-grpc_out=. \
		--go-grpc_opt=M$(SBPROTO)=$(SBPB) \
		$(SBPROTO)

	protoc  \
		--go_out=. \
		--go_opt=M$(ESPROTO)=$(ESPB) \
		--go-grpc_out=. \
		--go-grpc_opt=M$(ESPROTO)=$(ESPB) \
		$(ESPROTO)

wire-gen:
	wire ./internal/app

gen: protobuf wire-gen

coverage:
	go test -v -coverprofile cover.out ./...
	go tool cover -html cover.out -o cover.html
	rm cover.out

migrate.up:
	migrate -path ./migrations -database 'postgres://$(PG_USER):$(PG_PASS)@$(PG_HOST):$(PG_PORT)/$(PG_NAME)?sslmode=disable' up

migrate.down:
	migrate -path ./migrations -database 'postgres://$(PG_USER):$(PG_PASS)@$(PG_HOST):$(PG_PORT)/$(PG_NAME)?sslmode=disable' down

