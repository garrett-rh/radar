VERSION := 0.2.0

.PHONY: build
build:
	docker build -q -t radar:${VERSION} .

.PHONY: run
run: build
	docker run radar:${VERSION}

.PHONY: local
local:
	docker build -t radar:${VERSION} --output type=local,dest=dist .

.PHONY: deb
deb:
	docker build -t radar:${VERSION} --output type=local,dest=./dist/radar-${VERSION}/usr/local/bin/ .
	bash ./scripts/build-deb.sh ${VERSION}

.PHONY: local-install
local-install: deb
	sudo apt install ./dist/radar-${VERSION}.deb

.PHONY: uninstall
uninstall:
	sudo apt remove radar

.PHONY: clean
clean:
	docker image rm -f radar:${VERSION}
	rm -r dist/
