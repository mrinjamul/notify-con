#!/bin/bash

# Variables
APPNAME="$(basename $PWD)"
VERSION="v1.0.0"
ASSETS="assets README.md LICENSE"

# Functions

msg() {
    printf '%b\n' "$1" >&2
}

success() {
    msg "\33[32m[✔]\33[0m ${1}${2}"
}

error() {
    msg "\33[31m[✘]\33[0m ${1}${2}"
    exit 1
}

remove_build() {
    go clean
    rm -rf releases
    mkdir -p releases
    msg ""
}

linux_amd64_build() {
    msg "Building for Linux AMD64"
    sleep 1
    env GOOS=linux GOARCH=amd64 go build . && \
    tar czf releases/$APPNAME-linux-amd64-$VERSION.tar.gz $APPNAME $ASSETS && \
    success "Built for Linux AMD64"
}

linux_arm_build() {
    msg "Building for Linux ARM"
    sleep 1
    env GOOS=linux GOARCH=arm go build . && \
    tar czf releases/$APPNAME-linux-arm-$VERSION.tar.gz $APPNAME $ASSETS && \
    success "Built for Linux ARM"
}

macos_amd64_build() {
    msg "Building for Mac AMD64"
    sleep 1
    env GOOS=darwin GOARCH=amd64 go build . && \
    tar czf releases/$APPNAME-darwin-amd64-$VERSION.tar.gz $APPNAME $ASSETS && \
    success "Built for Mac AMD64"
}

windows_build() {
    msg "Building for Windows i386"
    sleep 1
    env GOOS=windows GOARCH=386 go build . && \
    zip -r releases/$APPNAME-windows-i386-$VERSION.zip $APPNAME.exe $ASSETS && \
    go clean && \
    success "Built for Windows i386"
}

initialize_app() {
    go build .
}

# Main

msg "Buiding $APPNAME with $VERSION"
msg ""

remove_build

linux_amd64_build
echo ""
linux_arm_build
echo ""
macos_amd64_build
echo ""

rm $APPNAME

windows_build
echo ""
initialize_app
success "All package releases are built"
echo ""
