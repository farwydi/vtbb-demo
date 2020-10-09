all: gen

generate:
	gqlgen \
	wire

gqlgen:
	go run -v github.com/99designs/gqlgen

wire:
	wire