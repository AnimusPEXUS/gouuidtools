export GONOPROXY=github.com/AnimusPEXUS/*

all: get

get: 
		$(MAKE) -C tests/test_01 get
		go get -u -v "./..."
		go mod tidy
