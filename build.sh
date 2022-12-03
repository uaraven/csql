#!/bin/sh

rm -rf bin
mkdir bin

function build() {
  echo "Building for $1 ($2)"
  GOOS=$1 GOARCH=$2  go build -ldflags="-X 'main.Version=$(cat version.txt)'"
  if [[ -f "csql.exe" ]] ; then
    mv csql.exe bin/csql_$1_$2.exe
  else
    mv csql bin/csql_$1_$2
  fi
}

function increment_version() {
  version=$( cat version.txt | awk -F. -v OFS=. '{$NF += 1 ; print}')
  echo "$version" > version.txt
}

increment_version

echo "Building version $(cat version.txt)"

build darwin amd64
build darwin arm64
build linux amd64
build windows amd64

