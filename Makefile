build-cmd:
	go build -o ./  ./cmd/main.go ./cmd/root.go ./cmd/app.go

run:
	./main app

build-and-run:
	make build-cmd && make run