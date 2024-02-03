#!/bin/sh
go build -o ./deb/usr/local/bin/wbwatcher ./main.go

dpkg-deb --build deb wbwatcher.deb
