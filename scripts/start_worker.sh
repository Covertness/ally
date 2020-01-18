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

if [[ -z $ETHERSCAN_APIKEY ]]; then
	echo "Please set the environment: ETHERSCAN_APIKEY."
	exit 1
fi

if [[ -z $CONTRACT_ADDRESS_FACTORY_ADDRESS ]]; then
	echo "Please set the environment: CONTRACT_ADDRESS_FACTORY_ADDRESS."
	exit 1
fi

go run cmd/worker/main.go
