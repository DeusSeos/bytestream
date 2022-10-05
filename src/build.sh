#!/bin/bash

# This script is used to build the project.
rm -r bin/bytestream
go build -o ../bin/bytestream bytestream.go