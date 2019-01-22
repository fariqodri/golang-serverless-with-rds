.PHONY: build clean deploy

build:
	SHELL=/bin/bash
	chmod +x ./scripts/build.sh
	./scripts/build.sh

clean:
	SHELL=/bin/bash
	chmod +x ./scripts/clean.sh
	./scripts/clean.sh

deploy: clean build
	SHELL=/bin/bash
	chmod +x ./scripts/deploy.sh
	./scripts/deploy.sh