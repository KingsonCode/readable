package main

import (
    "fmt"
    "os"
)

func main() {
    // Get the GITHUB_TOKEN from environment variables
    token := os.Getenv("GITHUB_TOKEN")
    if token == "" {
        fmt.Println("No GITHUB_TOKEN found!")
        return
    }

    // Use the token for GitHub API calls, etc.
    fmt.Println("Using token:", token)

    // Your application logic...
}
