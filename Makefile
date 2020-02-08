test:
	go test ./...

test_v:
	go test -v ./...

test_short:
	go test ./... -short

test_race:
	go test ./... -short -race

.PHONY: fmt
fmt:
	goimports -l -w pkg/
	goimports -l -w cmd/

.PHONY: vendor
vendor:
	go mod vendor

.PHONY: run
run:
	go run cmd/cleanarch-demo/*.go