VERSION := 0.2.0

.PHONY: build
build: 
	docker build -q -t sonar:${VERSION} .

.PHONY: run
run: build
	docker run sonar:${VERSION}

.PHONY: local
local: 
	docker build -t sonar:${VERSION} --output type=local,dest=dist .

.PHONY: deb
deb: 
	docker build -t sonar:${VERSION} --output type=local,dest=./dist/sonar-${VERSION}/usr/local/bin/ .
	bash ./scripts/build-deb.sh ${VERSION}

.PHONY: local-install
local-install: deb
	sudo apt install ./dist/sonar-${VERSION}.deb

.PHONY: uninstall
uninstall:
	sudo apt remove sonar

.PHONY: clean
clean: 
	docker image rm -f sonar:${VERSION}
	rm -r dist/