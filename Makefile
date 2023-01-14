VERSION := 0.1.0

.PHONY: build
build: 
	docker build -t rummage:${VERSION} .

.PHONY: run
run: 
	docker run rummage:${VERSION}

.PHONY: local
local: 
	docker build -t rummage:${VERSION} --output type=local,dest=dist .

.PHONY: deb
deb: 
	bash ./scripts/build-deb.sh ${VERSION}
	docker build -t rummage:${VERSION} --output type=local,dest=./dist/rummage-${VERSION}/usr/local/bin/ .

.PHONY: local-install
local-install: deb
	sudo apt install ./dist/rummage-${VERSION}.deb

.PHONY: uninstall
uninstall:
	sudo apt remove rummage

.PHONY: clean
clean: 
	docker image rm -f rummage:${VERSION}
	rm -r dist/