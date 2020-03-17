#!/bin/bash

set -e

SCRIPT_DIR=$(cd "$(dirname "$0")" ; pwd -P)
APP_DIR="$SCRIPT_DIR/app"

JAR_NAME='chess-clojure-0.1.0-SNAPSHOT-standalone.jar'
JAR_PATH="$APP_DIR/target/uberjar/$JAR_NAME"

RED='\033[0;31m'
GREEN='\033[0;32m'
NC='\033[0m'

function goal_build() {
  cd app && lein uberjar
  exit 0
}

function goal_smoketest() {
  local ACTUAL=$(java -jar $JAR_PATH)
  local EXPECTED='Hello, World!'
  if [ "$ACTUAL" != "$EXPECTED" ]
  then
    echo -e "${RED}SMOKETEST FAILED${NC} - expected: $EXPECTED, received: $ACTUAL"
    exit 1
  fi
  echo -e "${GREEN}SMOKETEST PASSED${NC} - received: $ACTUAL"
  exit 0
}

if type -t "goal_$1" &>/dev/null; then
  goal_$1 ${@:2}
else
  echo "Usage: $0 <goal>
goal:
    build        -- create jar
    smoketest    -- check that jar works
"
  exit 1
fi