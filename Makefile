export DOCKER_BUILDKIT=1

VERSION := 0.3.0
PROJ_NAME := radar

.PHONY: build
build:
	docker build -t ${PROJ_NAME}:${VERSION} .

.PHONY: test
test:
	docker build --target build -t ${PROJ_NAME}:${VERSION} .
	docker run ${PROJ_NAME}:${VERSION} go test ./...

.PHONY: run
run: build
	docker run ${PROJ_NAME}:${VERSION}

.PHONY: local
local:
	docker build -t ${PROJ_NAME}:${VERSION} --output type=local,dest=dist .

.PHONY: deb
deb:
	docker build -t ${PROJ_NAME}:${VERSION} --output type=local,dest=./dist/${PROJ_NAME}-${VERSION}/usr/local/bin/ .
	./scripts/build-deb.sh ${VERSION}

.PHONY: local-install
local-install: deb
	sudo apt install ./dist/${PROJ_NAME}-${VERSION}.deb

.PHONY: uninstall
uninstall:
	sudo apt remove ${PROJ_NAME}

.PHONY: clean
clean:
	docker image rm -f ${PROJ_NAME}:${VERSION}
	docker image rm -f ${PROJ_NAME}:${VERSION}-test
	rm -r dist/
