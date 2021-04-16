package main

import (
    "github.com/hideakikondo/youtube-manager-go/routes"
    "github.com/labstack/echo"
)

func main() {
    e := echo.New()
    // Routes
    routes.Init(e)

    // Start server
    e.Logger.Fatal(e.Start(":8080"))
}
