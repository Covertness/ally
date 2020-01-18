#!/bin/bash

SCRIPT_PATH="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"
cd "$SCRIPT_PATH/.."

if [ ! -f go.mod ]; then
	echo "Unknown location. Please run this from the ally repository. Abort."
	exit 1
fi

if [[ -z $INFURA_ID ]]; then
	echo "Please set the environment: INFURA_ID."
	exit 1
fi

go run contracts/deploy/*
