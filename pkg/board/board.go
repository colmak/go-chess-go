package board

import "fmt"

const (
    Rook   = 1
    Knight = 2
    Bishop = 3
    Queen  = 4
    King   = 5
    Pawn   = 6

    White = 8
    Black = 16
)

type Board struct {
    Squares [8][8]int 
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

func (b *Board) GetPieceAt(pos Position) int {
    return b.Squares[pos.Row][pos.Col]
}

func (b *Board) IsEmpty(pos Position) bool {
    return b.GetPieceAt(pos) == 0
}

func (b *Board) MovePiece(start, end Position) bool {
    piece := b.GetPieceAt(start)
    if piece == 0 {
        return false // No piece to move
    }

    // Check if the move is within board boundaries
    if !isWithinBounds(start) || !isWithinBounds(end) {
        return false // Move out of bounds
    }

    // Ensure the destination is not occupied by a piece of the same color
    destPiece := b.GetPieceAt(end)
    if destPiece != 0 && (destPiece&White) == (piece&White) {
        return false // Can't capture your own piece
    }

    // Check if the move is legal for the piece
    if !b.isLegalMove(piece, start, end) {
        return false // The move is not legal for this piece
    }

    // Perform the move
    b.Squares[end.Row][end.Col] = piece
    b.Squares[start.Row][start.Col] = 0
    return true
}

// Helper function to check if a position is within the board bounds
func isWithinBounds(pos Position) bool {
    return pos.Row >= 0 && pos.Row < 8 && pos.Col >= 0 && pos.Col < 8
}

// Method to check if a move is legal for a given piece
func (b *Board) isLegalMove(piece int, start, end Position) bool {
    switch piece & 0b111 { // Mask to get the piece type
    case Rook:
        return b.isLegalRookMove(start, end)
    case Knight:
        return b.isLegalKnightMove(start, end)
    case Bishop:
        return b.isLegalBishopMove(start, end)
    case Queen:
        return b.isLegalQueenMove(start, end)
    case King:
        return b.isLegalKingMove(start, end)
    case Pawn:
        return b.isLegalPawnMove(piece, start, end)
    }
    return false // Unsupported piece type
}

// Implement legal move checks for each piece type (these can be expanded)
func (b *Board) isLegalRookMove(start, end Position) bool {
    // Rooks move in straight lines (same row or same column)
    return (start.Row == end.Row || start.Col == end.Col) && b.isPathClear(start, end)
}

func (b *Board) isLegalKnightMove(start, end Position) bool {
    // Knights move in an L-shape
    rowDiff := abs(start.Row - end.Row)
    colDiff := abs(start.Col - end.Col)
    return (rowDiff == 2 && colDiff == 1) || (rowDiff == 1 && colDiff == 2)
}

func (b *Board) isLegalBishopMove(start, end Position) bool {
    // Bishops move diagonally
    return abs(start.Row-end.Row) == abs(start.Col-end.Col) && b.isPathClear(start, end)
}

func (b *Board) isLegalQueenMove(start, end Position) bool {
    // Queens move like both a rook and a bishop
    return b.isLegalRookMove(start, end) || b.isLegalBishopMove(start, end)
}

func (b *Board) isLegalKingMove(start, end Position) bool {
    // Kings move one square in any direction
    return abs(start.Row-end.Row) <= 1 && abs(start.Col-end.Col) <= 1
}

func (b *Board) isLegalPawnMove(piece int, start, end Position) bool {
    // Determine the direction based on piece color
    direction := 1
    if piece&Black != 0 {
        direction = -1
    }
    
    // Pawns move forward by one or two squares on their first move
    if start.Col == end.Col {
        if start.Row+direction == end.Row && b.IsEmpty(end) {
            return true
        }
        if (start.Row == 1 || start.Row == 6) && start.Row+2*direction == end.Row && b.IsEmpty(end) && b.IsEmpty(Position{Row: start.Row + direction, Col: start.Col}) {
            return true
        }
    }
    
    // Pawns capture diagonally
    if abs(start.Col-end.Col) == 1 && start.Row+direction == end.Row {
        // Regular capture
        if !b.IsEmpty(end) && (b.GetPieceAt(end)&White != piece&White) {
            return true
        }
        
        // En passant capture
        if b.canCaptureEnPassant(start, end, piece&Black != 0) {
            return true
        }
    }
    
    return false
}

// Implement the canCaptureEnPassant method
func (b *Board) canCaptureEnPassant(start, end Position, isBlack bool) bool {
    // Check the position of the target square
    if (isBlack && start.Row == 3) || (!isBlack && start.Row == 4) {
        // Determine the pawn that could be captured en passant
        sidePawnPos := Position{Row: start.Row, Col: end.Col}
        sidePawn := b.GetPieceAt(sidePawnPos)

        // Check if the pawn is of the opposite color and just moved two squares
        if sidePawn != 0 && sidePawn&Pawn != 0 && ((sidePawn&Black != 0) != isBlack) {
            lastRow := 6
            if isBlack {
                lastRow = 1
            }
            if b.GetPieceAt(Position{Row: lastRow, Col: end.Col}) == sidePawn {
                return true
            }
        }
    }

    return false
}



// Check if the path is clear between the start and end positions (for rooks, bishops, queens)
func (b *Board) isPathClear(start, end Position) bool {
    rowStep := sign(end.Row - start.Row)
    colStep := sign(end.Col - start.Col)

    currentRow := start.Row + rowStep
    currentCol := start.Col + colStep

    for currentRow != end.Row || currentCol != end.Col {
        if !b.IsEmpty(Position{Row: currentRow, Col: currentCol}) {
            return false
        }
        currentRow += rowStep
        currentCol += colStep
    }
    return true
}


func (b *Board) UndoMove(move Move) {
    b.Squares[move.Start.Row][move.Start.Col] = move.Piece
    b.Squares[move.End.Row][move.End.Col] = 0
}

func (b *Board) IsValidMove(move Move) bool {
    // Check if the move is valid
    return true
}

func (b *Board) GenerateMoves(pos Position) []Position {
    // Get all possible moves for the piece at the given position
    return nil
}

func (b *Board) IsCheck() bool {
    // Check if the current player is in check
    return false
}

func (b *Board) IsCheckmate() bool {
    // Check if the current player is in checkmate
    return false
}

func (b *Board) IsStalemate() bool {
    // Check if the current player is in stalemate
    return false
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
    b := &Board{}
    b.initPosition()
    return b
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

// Utility functions
func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}

func sign(x int) int {
    if x < 0 {
        return -1
    }
    if x > 0 {
        return 1
    }
    return 0
}