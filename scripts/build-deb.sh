#!/usr/bin/env bash 

ROOTDIR=$(pwd)

function main () {
    mkdir -p dist
    cd dist

	mkdir -p rummage-${1} rummage-${1}/DEBIAN rummage-${1}/usr/local/bin

    cd rummage-${1}

	cat << EOF > DEBIAN/control
Source: rummage
Maintainer: garretth@tuta.io
Section: misc
Priority: optional
Package: rummage
Version: ${1}
Architecture: all
Description: Tool to help navigate docker registries
EOF

    cd ..
    dpkg-deb --build rummage-${1}

    return 0
}

main $1