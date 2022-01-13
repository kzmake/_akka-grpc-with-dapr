SHELL = /bin/bash

.PHONY: all
all: proto fmt lint

.PHONY: compile
compile:
	sbt compile

.PHONY: test
test: compile
	sbt test

.PHONY: proto
proto:
	cd ./api && buf mod update && cd ..
	buf generate

.PHONY: fmt
fmt:
	sbt scalafixAll scalafmtAll

.PHONY: lint
lint:
	sbt "scalafixAll --check" scalafmtCheckAll

.PHONY: dev
dev:
	skaffold dev

.PHONY: run
run:
	skaffold run --tail --port-forward=user

.PHONY: http
http:
	curl -s -X GET localhost:8080/v1/wallets | jq .
	curl -s -X GET localhost:8080/v1/wallets/dummyid | jq .
	curl -s -X POST localhost:8080/v1/wallets -H "Content-Type: application/json" -d '{"owner": "alice"}' | jq .
	curl -s -X DELETE localhost:8080/v1/wallets/dummyid | jq .


.PHONY: grpc
grpc:
	grpcurl -plaintext -d '{"owner": "alice"}' localhost:50051 agwd.user.v1.UserService/Create
