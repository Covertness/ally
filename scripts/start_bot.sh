#!/bin/bash

SCRIPT_PATH="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"
cd "$SCRIPT_PATH/.."

if [ ! -f go.mod ]; then
	echo "Unknown location. Please run this from the ally repository. Abort."
	exit 1
fi

if [[ -z $TELEGRAM_TOKEN ]]; then
	echo "Please set the environment: TELEGRAM_TOKEN."
	exit 1
fi

if [[ -z $FTX_KEY ]]; then
	echo "Please set the environment: FTX_KEY."
	exit 1
fi

if [[ -z $FTX_SECRET ]]; then
	echo "Please set the environment: FTX_SECRET."
	exit 1
fi

go run cmd/bot/main.go