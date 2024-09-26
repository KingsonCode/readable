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
          go-version: '1.23'

      - name: Install dependencies
        run: |
          cd my-go-project/
          go mod download

      - name: Build the project
        run: |
          cd my-go-project/
          go build ./...

      - name: Run tests
        run: |
          cd my-go-project/
          go test ./...
