package main

import (
    "fmt"
    "net/http"
    "time"
    
    "github.com/gin-gonic/gin"
    "github.com/gin-contrib/cors"


    "github.com/colmak/go-chess-go/pkg/board"
)

// Global board instance to keep track of the game state
var gameBoard *board.Board

func initialize() {
    fmt.Println("Initializing the Go Chess Go Engine")
    gameBoard = board.NewBoard() // Initialize the board
    gameBoard.PrintBoard()       // Optionally print the initial board state
}

func main() {
    initialize() // Initialize the chess engine

    // Set up the Gin router
    r := gin.Default()

    // Add CORS middleware
    r.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"http://localhost:3000"},
        AllowMethods:     []string{"GET", "POST", "OPTIONS"},
        AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
        MaxAge:           12 * time.Hour,
    }))

    // API routes
    r.GET("/status", getStatus)
    r.POST("/move", makeMove)
    r.POST("/reset", resetGame)

    // Start the API server on port 8080
    r.Run(":8080")
}


// getStatus returns the current state of the board
func getStatus(c *gin.Context) {
    gameBoard.PrintBoard()
    c.JSON(http.StatusOK, gin.H{
        "board": gameBoard.Squares, // Return the board's squares array
        "turn":  gameBoard.CurrentTurn,
    })
}

// Move struct for receiving move data
type Move struct {
    StartRow int `json:"start_row"`
    StartCol int `json:"start_col"`
    EndRow   int `json:"end_row"`
    EndCol   int `json:"end_col"`
}

func makeMove(c *gin.Context) {
    var move Move
    if err := c.ShouldBindJSON(&move); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    start := board.Position{Row: move.StartRow, Col: move.StartCol}
    end := board.Position{Row: move.EndRow, Col: move.EndCol}

    fmt.Printf("Current turn before move: %d\n", gameBoard.CurrentTurn)
    fmt.Printf("Switching turn after move...\n")

    if gameBoard.MovePiece(start, end) {
        c.JSON(http.StatusOK, gin.H{
            "message": "Move successful",
            "board":   gameBoard.Squares,
            "turn":    gameBoard.CurrentTurn,
        })
    } else {
        c.JSON(http.StatusBadRequest, gin.H{
            "message": "Invalid move",
        })
    }
}

// resetGame resets the chess game
func resetGame(c *gin.Context) {
    gameBoard = board.NewBoard() // Reinitialize the board
    gameBoard.PrintBoard()       // Optionally print the reset board
    c.JSON(http.StatusOK, gin.H{
        "message": "Game reset",
        "board":   gameBoard.Squares,
    })
}
