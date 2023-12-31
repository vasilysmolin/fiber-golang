build-go:
	docker build -t cr.yandex/crp4b3g2mro7h6fi2tng/go:$(branch) -f Dockerfile .
	docker push cr.yandex/crp4b3g2mro7h6fi2tng/go:$(branch)

start:
	go run main.go

get:
	go get

mod:
	go mod tidy

test:
	go test ./tests...

lint:
	find . -name "*.go" -not -path "./src/*"  -exec gofmt -d {} \;

lint-fix:
	find . -name "*.go" -not -path "./src/*" -exec gofmt -w {} \;
