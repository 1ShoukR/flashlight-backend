package main

import (
	"fmt"

	"github.com/1ShoukR/flashlight-backend/internal/server"
)

func main() {
    fmt.Println("Hello, World!")
    err := server.NewServer().ListenAndServe()

    if err != nil {
        panic(fmt.Sprintf("cannot start server: %s", err))
    }
}