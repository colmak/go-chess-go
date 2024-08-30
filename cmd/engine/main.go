package main

import (
    "fmt"
    "github.com/colmak/go-chess-go/pkg/board"
)

func initialize() {
    fmt.Println("Initializing the Go Chess Go Engine")

}

func main() {
    initialize()
    b := board.NewBoard()
    fmt.Println(b.Squares)
}