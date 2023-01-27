export DOCKER_BUILDKIT=1

VERSION := 0.2.1
PROJ_NAME := radar

.PHONY: build
build:
	docker build -t ${PROJ_NAME}:${VERSION} .

.PHONY: test
test:
	docker build --target test -t ${PROJ_NAME}:${VERSION}-test .
	docker run ${PROJ_NAME}:${VERSION}-tests go test ./...

.PHONY: run
run: build
	docker run --network host ${PROJ_NAME}:${VERSION}

.PHONY: local
local:
	docker build -t ${PROJ_NAME}:${VERSION} --output type=local,dest=dist .

.PHONY: deb
deb:
	docker build -t ${PROJ_NAME}:${VERSION} --output type=local,dest=./dist/${PROJ_NAME}-${VERSION}/usr/local/bin/ .
	bash ./scripts/build-deb.sh ${VERSION}

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
