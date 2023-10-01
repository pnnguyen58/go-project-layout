
protoc:
	protoc -I . \
		--go_out=./pkg \
		--go_opt=module=github.com/pnnguyen58/go-project-layout \
		--go-grpc_out=./pkg \
		--go-grpc_opt=module=github.com/pnnguyen58/go-project-layout \
		--grpc-gateway_out=./pkg \
		--grpc-gateway_opt module=github.com/pnnguyen58/go-project-layout \
		--grpc-gateway_opt logtostderr=true \
		--grpc-gateway_opt generate_unbound_methods=true \
		./pkg/proto/*.proto

buildup:
	./scripts/docker-compose/setup.sh
	docker compose -f ./infra/docker-compose.yml up --build -d

start:
	./scripts/docker-compose/setup.sh
	docker compose -f ./infra/docker-compose.yml up -d

stop:
	docker compose -f ./infra/docker-compose.yml down

reset:
	docker compose -f ./infra/docker-compose.yml down
	./scripts/docker-compose/reset.sh
