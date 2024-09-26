package main

name: Go Build and Test

on:
  push:
    branches:
      - main
  pull_request:
    branches:
      - main

jobs:
  build:
    runs-on: ubuntu-latest

    steps:
      - name: Checkout code
        uses: actions/checkout@v3

      - name: Set up Go
        uses: actions/setup-go@v4
        with:
          go-version: '1.23'  # Ensure this matches the version in your go.mod

      - name: Install dependencies
        run: go mod tidy  # Downloads dependencies and cleans up go.mod/go.sum files

      - name: Build the project
        run: go build ./...  # Builds the Go project

      - name: Run tests
        run: go test ./...  # Runs tests in all Go files
