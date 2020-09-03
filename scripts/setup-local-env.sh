#!/usr/bin/env sh

[ ! -f .env ] && cp .env.example .env || true # Copy the .env.example to .env file

if ! command -v pre-commit &> /dev/null
then
	echo "Installing pre-commit"
    pip install pre-commit
fi

pre-commit install
