export GONOPROXY=github.com/AnimusPEXUS/*

all: get

get: 
		make -C tests/test_01
		go get -u -v "./..."
		go mod tidy
