ENVIRONMENT?=dev

run:
	go run main.go

test:
	go test ./... -coverprofile cover.out -race
	go tool cover -func cover.out