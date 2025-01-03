package board

import "fmt"

type Board struct {
    Squares [8][8]int 
    WhiteKingMoved bool
    BlackKingMoved bool
    WhiteRookMoved [2]bool
    BlackRookMoved [2]bool
    CurrentTurn int     
    HalfMoveClock int
    MoveCount int  
    FiftyMoveCount int
    PositionHistory map[string]int 
    LastMove Move
}

type Position struct {
    Row int
    Col int
}

type Move struct {
    Start Position
    End   Position
    Piece int
}

func (b *Board) isPathClear(start, end Position) bool {
    rowStep := sign(end.Row - start.Row)
    colStep := sign(end.Col - start.Col)
    
    currentRow := start.Row + rowStep
    currentCol := start.Col + colStep

    for currentRow != end.Row || currentCol != end.Col {
        if !b.IsEmpty(Position{currentRow, currentCol}) {
            return false
        }
        currentRow += rowStep
        currentCol += colStep
    }

    return true
}

func (b *Board) PrintBoard() {
    for i := 0; i < 8; i++ {
        for j := 0; j < 8; j++ {
            fmt.Printf("%d ", b.Squares[i][j])
        }
        fmt.Println()
    }
}

func NewBoard() *Board {
    b := &Board{
        CurrentTurn:     White,
        HalfMoveClock:   0,
        PositionHistory: make(map[string]int),
    }
    b.initPosition() 
    return b
}

func (b *Board) GetCurrentTurn() int {
    return b.CurrentTurn
}

func (b *Board) initPosition() {
    b.Squares[0][0] = Rook | White
    b.Squares[0][1] = Knight | White
    b.Squares[0][2] = Bishop | White
    b.Squares[0][3] = Queen | White
    b.Squares[0][4] = King | White
    b.Squares[0][5] = Bishop | White
    b.Squares[0][6] = Knight | White
    b.Squares[0][7] = Rook | White
    for i := 0; i < 8; i++ {
        b.Squares[1][i] = Pawn | White
    }
    b.Squares[7][0] = Rook | Black
    b.Squares[7][1] = Knight | Black
    b.Squares[7][2] = Bishop | Black
    b.Squares[7][3] = Queen | Black
    b.Squares[7][4] = King | Black
    b.Squares[7][5] = Bishop | Black
    b.Squares[7][6] = Knight | Black
    b.Squares[7][7] = Rook | Black
    for i := 0; i < 8; i++ {
        b.Squares[6][i] = Pawn | Black
    }
}