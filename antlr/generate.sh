#!/bin/sh

SAVE=$(pwd)

cd $(dirname "$0")

rm ../parser/*

java -jar antlr-4.11.1-complete.jar -no-listener -visitor -Dlanguage=Go -o ../parser Csql.g4

cd $SAVE