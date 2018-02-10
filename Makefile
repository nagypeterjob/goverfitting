.PHONY: test

test:
	go test -v $(shell go list ./... | grep -v /vendor/)

.PHONY: rocker

rocker:
	rocker build .
	-./tools/clean_docker.sh

.PHONY: compose
compose:
	$(MAKE) rocker
	rocker-compose run --wait 5s

.PHONY: compose-clean
compose-clean:
	rocker-compose rm
	-./tools/clean_docker.sh