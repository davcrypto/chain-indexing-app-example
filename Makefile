all: build install-migrate install
build: chain-indexing-app
chain-indexing-app:
	go build ./app/chain-indexing-app/
install:
	go install ./cmd/chain-indexing/
install-migrate:
	./pgmigrate.sh --install-dependency -- version
migrate:
	./pgmigrate.sh -- -verbose up
clean:
	rm chain-indexing-app

.PHONY: docker
docker:
	docker build -t chain-indexing -f Dockerfile .
