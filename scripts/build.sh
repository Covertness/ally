#!/bin/bash

SCRIPT_PATH="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"
cd "$SCRIPT_PATH/.."

if [ ! -f go.mod ]; then
	echo "Unknown location. Please run this from the ally repository. Abort."
	exit 1
fi

go build -o bin/api cmd/api/main.go
go build -o bin/dispatcher cmd/dispatcher/main.go
go build -o bin/worker cmd/worker/main.go