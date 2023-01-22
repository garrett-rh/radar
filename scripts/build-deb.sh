#!/usr/bin/env bash

ROOTDIR=$(pwd)

function main () {
    mkdir -p dist
    cd dist

	mkdir -p sonar-${1} sonar-${1}/DEBIAN sonar-${1}/usr/local/bin

    cd sonar-${1}

	cat << EOF > DEBIAN/control
Source: sonar
Maintainer: garretth@tuta.io
Section: misc
Priority: optional
Package: sonar
Version: ${1}
Architecture: all
Description: Tool to help navigate docker registries
EOF

    cd ..
    dpkg-deb --build sonar-${1}

    return 0
}

main $1
