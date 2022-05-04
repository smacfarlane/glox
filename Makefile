.PHONY: generate
generate: 
	mkdir -p internal/ast/gen
	rm -f internal/ast/expressions.go
	go run ./tools/generator/ast/main.go | gofmt > internal/ast/expressions.go
