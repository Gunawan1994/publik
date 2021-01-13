GO_LDFLAGS = -ldflags "-s -w"

.PHONY: build build-linux build-fe test run

build:
	@echo "${NOW} == BUILDING BACKEND..."
	@go build -o bin/http-server $(GO_LDFLAGS) cmd/http-server/*.go

build-linux:
	@echo "${NOW} == BUILDING BACKEND FOR LINUX..."
	@GOOS=linux GOARCH=amd64 go build -o bin/http-server-linux $(GO_LDFLAGS) cmd/http-server/*.go

run-http: build
	@echo "${NOW} == RUN DEVELOPMENT SERVER..."
	@./bin/http-server -c http-config.yaml

test:
	@echo "${NOW} == TESTING..."
	@go test -cover -race ./...

run-mock:
	@java -jar mocking/wiremock-standalone-2.26.3.jar --disable-banner true --root-dir mocking/

run-docker:
	@docker-compose up -d --build

stop-docker:
	@docker-compose down