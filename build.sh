#!/bin/bash

#
#    Copyright 2022 Oleksiy Voronin
#
#    Licensed under the Apache License, Version 2.0 (the "License");
#    you may not use this file except in compliance with the License.
#    You may obtain a copy of the License at
#
#        http://www.apache.org/licenses/LICENSE-2.0
#
#    Unless required by applicable law or agreed to in writing, software
#    distributed under the License is distributed on an "AS IS" BASIS,
#    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
#    See the License for the specific language governing permissions and
#    limitations under the License.
#
#    SPDX-License-Identifier: Apache-2.0
#    SPDX-FileCopyrightText: (c) 2022 Oleksiy Voronin <ovoronin@gmail.com>
#

rm -rf bin
mkdir bin

function build() {
  echo "Building for $1 ($2)"
  GOOS=$1 GOARCH=$2  go build -ldflags="-s -w -X 'main.Version=$(cat version.txt)'"
  zip csql_$1_$2.zip csql*
  mv csql* bin/
}

function increment_version() {
  version=$( cat version.txt | awk -F. -v OFS=. '{$NF += 1 ; print}')
  echo "$version" > version.txt
}

if [[ "$1" != "--no-version" ]] ; then
  increment_version
else
  echo "Skipping version increment"
fi

echo "Building version $(cat version.txt)"

build darwin amd64
build darwin arm64
build linux amd64
build windows amd64

