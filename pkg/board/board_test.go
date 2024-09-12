package board

import (
    "testing"
)

func TestNewBoard(t *testing.T) {
    b := NewBoard()
    if b.Squares[0][0] != (Rook | White) {
        t.Errorf("Expected Rook at (0,0), got %d", b.Squares[0][0])
    }
}

func TestMovePiece(t *testing.T) {
    b := NewBoard()
    start := Position{Row: 1, Col: 0}
    end := Position{Row: 3, Col: 0}
    if !b.MovePiece(start, end) {
        t.Error("Expected move to be successful")
    }
    if b.Squares[3][0] != (Pawn | White) {
        t.Errorf("Expected Pawn at (3,0), got %d", b.Squares[3][0])
    }
    if b.Squares[1][0] != 0 {
        t.Errorf("Expected empty square at (1,0), got %d", b.Squares[1][0])
    }
}

func TestUndoMove(t *testing.T) {
    b := NewBoard()
    start := Position{Row: 1, Col: 0}
    end := Position{Row: 3, Col: 0}
    move := Move{Start: start, End: end, Piece: b.GetPieceAt(start)}
    b.MovePiece(start, end)
    b.UndoMove(move)
    if b.Squares[1][0] != (Pawn | White) {
        t.Errorf("Expected Pawn at (1,0), got %d", b.Squares[1][0])
    }
    if b.Squares[3][0] != 0 {
        t.Errorf("Expected empty square at (3,0), got %d", b.Squares[3][0])
    }
}

func TestIsEmpty(t *testing.T) {
    b := NewBoard()
    pos := Position{Row: 2, Col: 2}
    if !b.IsEmpty(pos) {
        t.Error("Expected position to be empty")
    }
    pos = Position{Row: 0, Col: 0}
    if b.IsEmpty(pos) {
        t.Error("Expected position to be occupied")
    }
}