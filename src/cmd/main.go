package main

import (
    "project03/internal/di"
)

func main() {
    app := di.BuildContainer()
    app.Run()
}