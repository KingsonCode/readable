name: Go Build and Test - Token List

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
          go-version: '1.20'

      - name: Install dependencies
        run: go mod download

      - name: Build the project
        run: go build ./...

      - name: Run tests
        run: go test ./...
