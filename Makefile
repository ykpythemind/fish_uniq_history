build:
	go build .

test:
	go test

bench:
	go test -bench .

install:
	make build && cp fish_uniq_history /usr/local/bin/
