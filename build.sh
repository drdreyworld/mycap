#!/bin/bash

path="`pwd`"
binpath="$path/bin/"

build_target="$1"

echo "MyCAP builder"

if [ ! -z "$build_target" ]
then
  echo "Build target: $build_target"
fi;


if [ "$build_target" == "agent" ] || [ -z "$build_target" ]
then
  echo "Build agent"
  cd $path/agent/bin && go build -o $binpath/agent
fi;

if [ "$build_target" == "server" ] || [ -z "$build_target" ]
then
  echo "Build server"
  cd $path/server/bin && go build -o $binpath/server
fi;

if [ "$build_target" == "web" ] || [ -z "$build_target" ]
then
  echo "Build web"
  cd $path/web/bin && go build -o $binpath/web
fi;

if [ "$build_target" == "tool" ] || [ -z "$build_target" ]
then
  echo "Build tool"
  cd $path/tool && go build -o $binpath/tool
fi;
