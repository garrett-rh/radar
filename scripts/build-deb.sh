#!/usr/bin/env bash

ROOTDIR=$(pwd)

function main () {
    mkdir -p dist
    cd dist

	mkdir -p radar-${1} radar-${1}/DEBIAN radar-${1}/usr/local/bin

    cd radar-${1}

	cat << EOF > DEBIAN/control
Source: radar
Maintainer: garretth@tuta.io
Section: misc
Priority: optional
Package: radar
Version: ${1}
Architecture: all
Description: Tool to help navigate docker registries
EOF

    cd ..
    dpkg-deb --build radar-${1}

    return 0
}

main $1
