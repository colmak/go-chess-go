package board

import (
    "testing"
)

func TestNewBoard(t *testing.T) {
    b := NewBoard()
    if b.Squares[0][0] != (Rook | White) {
        t.Errorf("Expected Rook at (0,0), got %d", b.Squares[0][0])
    }
    if b.Squares[7][4] != (King | Black) {
        t.Errorf("Expected Black King at (7,4), got %d", b.Squares[7][4])
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

func TestKingsideCastling(t *testing.T) {
    b := NewBoard()

    // Clear the path for White castling
    b.Squares[0][5] = 0
    b.Squares[0][6] = 0
    b.WhiteKingMoved = false
    b.WhiteRookMoved[1] = false

    if !b.canCastleKingside(false) {
        t.Error("Expected White to be able to castle kingside")
    }
}

func TestQueensideCastling(t *testing.T) {
    b := NewBoard()

    // Clear the path for Black castling
    b.Squares[7][1] = 0
    b.Squares[7][2] = 0
    b.Squares[7][3] = 0
    b.BlackKingMoved = false
    b.BlackRookMoved[0] = false

    if !b.canCastleQueenside(true) {
        t.Error("Expected Black to be able to castle queenside")
    }
}

func TestIsCheck(t *testing.T) {
    b := NewBoard()

    // Place a Black rook in a position to check the White king
    b.Squares[1][4] = (Rook | Black)

    // Print White king's position for debugging
    whiteKingPos := b.findKing(false)
    t.Logf("White King Position: %v", whiteKingPos)

    if !b.IsCheck(false) {
        t.Error("Expected White king to be in check")
    }
}


func TestIsCheckmate(t *testing.T) {
    b := NewBoard()

    // Simulate a checkmate by placing Black pieces
    b.Squares[7][4] = (King | White)
    b.Squares[6][5] = (Queen | Black)
    b.Squares[6][3] = (Rook | Black)

    if !b.IsCheckmate(false) {
        t.Error("Expected White to be in checkmate")
    }
}

func TestIsStalemate(t *testing.T) {
    b := NewBoard()

    // Set up a stalemate position (Black king on h8, White king on f6, White pawn on g7)
    b.Squares[7][7] = (King | Black)
    b.Squares[5][5] = (King | White)
    b.Squares[6][6] = (Pawn | White)

    if !b.IsStalemate(true) {
        t.Error("Expected Black to be in stalemate")
    }
}

// --- Piece-specific movement rules ---
func TestPawnMovement(t *testing.T) {
    b := NewBoard()

    // Test pawn single step forward
    if !b.MovePiece(Position{1, 0}, Position{2, 0}) {
        t.Error("Expected pawn to move forward by 1 step")
    }

    // Test pawn double step forward on initial move
    b = NewBoard()
    if !b.MovePiece(Position{1, 0}, Position{3, 0}) {
        t.Error("Expected pawn to move forward by 2 steps")
    }

    // Test pawn diagonal capture
    b.Squares[2][1] = (Pawn | Black)
    if !b.MovePiece(Position{1, 0}, Position{2, 1}) {
        t.Error("Expected pawn to capture diagonally")
    }

    // Test en passant
    b = NewBoard()
    
    // Black pawn moves two steps forward
    if !b.MovePiece(Position{6, 1}, Position{4, 1}) {
        t.Error("Expected black pawn to move forward by 2 steps")
    }
    
    // White pawn moves two steps forward
    if !b.MovePiece(Position{1, 0}, Position{3, 0}) {
        t.Error("Expected white pawn to move forward by 2 steps")
    }
    
    // Check en passant capture (white pawn captures black pawn)
    if !b.canCaptureEnPassant(Position{3, 0}, Position{4, 1}, false) {
        t.Error("Expected en passant to be possible for white pawn")
    }
    
    // Perform en passant move
    if !b.MovePiece(Position{3, 0}, Position{4, 1}) {
        t.Error("Expected en passant capture to be successful")
    }

    // Check that the black pawn was captured
    if b.GetPieceAt(Position{4, 1}) != (Pawn | White) {
        t.Error("Expected white pawn to be at (4,1) after en passant")
    }
    if b.GetPieceAt(Position{4, 0}) != 0 {
        t.Error("Expected black pawn to be captured")
    }
}

func TestRookMovement(t *testing.T) {
    b := NewBoard()

    // Test rook horizontal move
    b.Squares[0][0] = (Rook | White)
    if !b.MovePiece(Position{0, 0}, Position{0, 5}) {
        t.Error("Expected rook to move horizontally")
    }

    // Test rook vertical move
    b = NewBoard()
    if !b.MovePiece(Position{0, 0}, Position{5, 0}) {
        t.Error("Expected rook to move vertically")
    }
}

func TestKnightMovement(t *testing.T) {
    b := NewBoard()

    // Test knight L-shaped move
    if !b.MovePiece(Position{0, 1}, Position{2, 2}) {
        t.Error("Expected knight to move in L-shape")
    }
}

func TestBishopMovement(t *testing.T) {
    b := NewBoard()

    // Test bishop diagonal move
    b.Squares[0][2] = (Bishop | White)
    if !b.MovePiece(Position{0, 2}, Position{3, 5}) {
        t.Error("Expected bishop to move diagonally")
    }
}

func TestQueenMovement(t *testing.T) {
    b := NewBoard()

    // Test queen horizontal move
    b.Squares[0][3] = (Queen | White)
    if !b.MovePiece(Position{0, 3}, Position{0, 5}) {
        t.Error("Expected queen to move horizontally")
    }

    // Test queen diagonal move
    if !b.MovePiece(Position{0, 5}, Position{3, 8}) {
        t.Error("Expected queen to move diagonally")
    }
}

func TestKingMovement(t *testing.T) {
    b := NewBoard()

    // Test king one step move
    if !b.MovePiece(Position{0, 4}, Position{1, 4}) {
        t.Error("Expected king to move one square")
    }

    // Test castling kingside
    b = NewBoard()
    b.Squares[0][5], b.Squares[0][6] = 0, 0 // Clear squares between king and rook
    if !b.MovePiece(Position{0, 4}, Position{0, 6}) {
        t.Error("Expected kingside castling to work")
    }

    // Test castling queenside
    b = NewBoard()
    b.Squares[0][1], b.Squares[0][2], b.Squares[0][3] = 0, 0, 0 // Clear squares
    if !b.MovePiece(Position{0, 4}, Position{0, 2}) {
        t.Error("Expected queenside castling to work")
    }
}

// --- Check and Checkmate ---
func TestCheck(t *testing.T) {
    b := NewBoard()
    b.Squares[1][4] = (Rook | Black) // Place Black rook to check White king
    if !b.IsCheck(false) { // White is in check
        t.Error("Expected White to be in check")
    }
}

func TestCheckmate(t *testing.T) {
    b := NewBoard()
    b.Squares[0][4] = (King | White) // White king
    b.Squares[7][4] = (King | Black) // Black king
    b.Squares[7][7] = (Rook | Black) // Black rook to deliver checkmate
    if !b.IsCheckmate(false) {
        t.Error("Expected White to be in checkmate")
    }
}

func TestStalemate(t *testing.T) {
    b := NewBoard()
    b.Squares[0][4] = (King | White) // White king in stalemate position
    b.Squares[1][5] = (Rook | Black) // Black rook blocks
    if !b.IsStalemate(false) {
        t.Error("Expected White to be in stalemate")
    }
}

// --- Captures ---
func TestCaptures(t *testing.T) {
    b := NewBoard()

    // Test capturing opponent piece
    b.Squares[1][0] = (Rook | White)
    b.Squares[2][0] = (Pawn | Black)
    if !b.MovePiece(Position{1, 0}, Position{2, 0}) {
        t.Error("Expected White rook to capture Black pawn")
    }

    // Test attempting to capture same color piece
    b.Squares[3][0] = (Pawn | White)
    if b.MovePiece(Position{1, 0}, Position{3, 0}) {
        t.Error("Expected move to fail, cannot capture same color piece")
    }
}

// --- Special Rules ---
func TestPawnPromotion(t *testing.T) {
    b := NewBoard()
    b.Squares[6][0] = (Pawn | White)
    b.MovePiece(Position{6, 0}, Position{7, 0}) // Move pawn to promotion rank
    if b.Squares[7][0] != (Queen | White) {     // Check if promoted to Queen
        t.Error("Expected pawn promotion to Queen")
    }
}

func Test50MoveRule(t *testing.T) {
    b := NewBoard()

    // Simulate 50 moves without a pawn move or capture
    for i := 0; i < 50; i++ {
        b.MovePiece(Position{0, 1}, Position{0, 2}) // Move some random pieces
    }

    if !b.IsDrawByFiftyMoveRule() {
        t.Error("Expected game to be a draw by the 50-move rule")
    }
}

func TestThreefoldRepetition(t *testing.T) {
    b := NewBoard()

    // Simulate threefold repetition of a position
    b.MovePiece(Position{1, 0}, Position{2, 0})
    b.MovePiece(Position{2, 0}, Position{1, 0}) // Repeat back and forth

    b.MovePiece(Position{1, 0}, Position{2, 0})
    b.MovePiece(Position{2, 0}, Position{1, 0}) // Repeat again

    b.MovePiece(Position{1, 0}, Position{2, 0})
    b.MovePiece(Position{2, 0}, Position{1, 0}) // Third time

    if !b.IsDrawByThreefoldRepetition() {
        t.Error("Expected game to be a draw by threefold repetition")
    }
}

// --- Turn Management ---
func TestTurnManagement(t *testing.T) {
    b := NewBoard()

    // Test that White goes first
    if b.GetCurrentTurn() != White {
        t.Error("Expected White to go first")
    }

    // Test that turns alternate
    b.MovePiece(Position{1, 0}, Position{2, 0})
    if b.GetCurrentTurn() != Black {
        t.Error("Expected Black's turn after White's move")
    }
}