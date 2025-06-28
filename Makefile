build:
	go build -o gogen cmd/codegen/main.go
run-build:
	./gogen basic --name myapp --dir ./output