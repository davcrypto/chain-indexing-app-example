all: build install-migrate
build: chain-indexing-app
chain-indexing-app:
	go build ./app/chain-indexing-app/
install-migrate:
	./pgmigrate.sh --install-dependency -- version
migrate:
	./pgmigrate.sh -- -verbose up
clean:
	rm chain-indexing-app

.PHONY: docker
docker:
	docker build -t chain-indexing -f Dockerfile .
