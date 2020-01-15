#!/bin/bash

SCRIPT_PATH="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"
cd "$SCRIPT_PATH/.."

if [ ! -f go.mod ]; then
	echo "Unknown location. Please run this from the ally repository. Abort."
	exit 1
fi

PACKAGE_MANAGER=
if [[ "$OSTYPE" == "linux-gnu" ]]; then
	if which yum &>/dev/null; then
		PACKAGE_MANAGER="yum"
	elif which apt-get &>/dev/null; then
		PACKAGE_MANAGER="apt-get"
	elif which pacman &>/dev/null; then
		PACKAGE_MANAGER="pacman"
	else
		echo "Unable to find supported package manager (yum, apt-get, or pacman). Abort"
		exit 1
	fi
elif [[ "$OSTYPE" == "darwin"* ]]; then
	if which brew &>/dev/null; then
		PACKAGE_MANAGER="brew"
	else
		echo "Missing package manager Homebrew (https://brew.sh/). Abort"
		exit 1
	fi
else
	echo "Unknown OS. Abort."
	exit 1
fi

if [[ $"$PACKAGE_MANAGER" == "apt-get" ]]; then
	echo "Updating apt-get......"
	sudo apt-get update
fi

echo "Installing Go......"
if which go &>/dev/null; then
	echo "Go is already installed"
else
	if [[ "$PACKAGE_MANAGER" == "yum" ]]; then
		sudo yum install golang -y
	elif [[ "$PACKAGE_MANAGER" == "apt-get" ]]; then
		sudo apt-get install golang -y
	elif [[ "$PACKAGE_MANAGER" == "pacman" ]]; then
		sudo pacman -Syu golang --noconfirm
	elif [[ "$PACKAGE_MANAGER" == "brew" ]]; then
		brew install go
	fi
fi

echo "Installing Beanstalk......"
if which beanstalkd &>/dev/null; then
	echo "Beanstalk is already installed"
else
	if [[ "$PACKAGE_MANAGER" == "yum" ]]; then
		sudo yum install beanstalkd -y
	elif [[ "$PACKAGE_MANAGER" == "apt-get" ]]; then
		sudo apt-get install beanstalkd -y
	elif [[ "$PACKAGE_MANAGER" == "pacman" ]]; then
		sudo pacman -Syu beanstalkd --noconfirm
	elif [[ "$PACKAGE_MANAGER" == "brew" ]]; then
		brew install beanstalkd
	fi
fi
