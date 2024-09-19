package board

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

    // Check if the move is legal for the piece, including castling logic for kings
    if !b.isLegalMove(piece, start, end) {
        return false // The move is not legal for this piece
    }

    // Handle castling logic
    if piece&King != 0 && abs(start.Col-end.Col) == 2 {
        if end.Col == 6 { // Kingside castling
            if piece&Black != 0 {
                b.Squares[7][5] = b.Squares[7][7] // Move rook
                b.Squares[7][7] = 0
            } else {
                b.Squares[0][5] = b.Squares[0][7] // Move rook
                b.Squares[0][7] = 0
            }
        } else if end.Col == 2 { // Queenside castling
            if piece&Black != 0 {
                b.Squares[7][3] = b.Squares[7][0] // Move rook
                b.Squares[7][0] = 0
            } else {
                b.Squares[0][3] = b.Squares[0][0] // Move rook
                b.Squares[0][0] = 0
            }
        }
    }

    // Perform the move
    b.Squares[end.Row][end.Col] = piece
    b.Squares[start.Row][start.Col] = 0

    // Track if the king or rooks moved (to disable castling)
    if piece&King != 0 {
        if piece&Black != 0 {
            b.BlackKingMoved = true
        } else {
            b.WhiteKingMoved = true
        }
    } else if piece&Rook != 0 {
        if piece&Black != 0 {
            if start.Row == 7 && start.Col == 0 {
                b.BlackRookMoved[0] = true // Queenside rook
            } else if start.Row == 7 && start.Col == 7 {
                b.BlackRookMoved[1] = true // Kingside rook
            }
        } else {
            if start.Row == 0 && start.Col == 0 {
                b.WhiteRookMoved[0] = true // Queenside rook
            } else if start.Row == 0 && start.Col == 7 {
                b.WhiteRookMoved[1] = true // Kingside rook
            }
        }
    }

    return true
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
    // Check if it's a normal one-square king move
    if abs(start.Row-end.Row) <= 1 && abs(start.Col-end.Col) <= 1 {
        return true
    }
    
    // Castling logic
    piece := b.GetPieceAt(start)
    isBlack := piece&Black != 0
    
    // Castling kingside
    if start.Row == end.Row && abs(start.Col-end.Col) == 2 {
        if end.Col == 6 { // Kingside castling
            if (isBlack && start.Row == 7 && b.canCastleKingside(true)) || (!isBlack && start.Row == 0 && b.canCastleKingside(false)) {
                // Move the rook during castling
                if isBlack {
                    b.Squares[7][5] = b.Squares[7][7] // Move rook
                    b.Squares[7][7] = 0
                } else {
                    b.Squares[0][5] = b.Squares[0][7] // Move rook
                    b.Squares[0][7] = 0
                }
                return true
            }
        }
        // Castling queenside
        if end.Col == 2 { 
            if (isBlack && start.Row == 7 && b.canCastleQueenside(true)) || (!isBlack && start.Row == 0 && b.canCastleQueenside(false)) {
                // Move the rook during castling
                if isBlack {
                    b.Squares[7][3] = b.Squares[7][0] // Move rook
                    b.Squares[7][0] = 0
                } else {
                    b.Squares[0][3] = b.Squares[0][0] // Move rook
                    b.Squares[0][0] = 0
                }
                return true
            }
        }
    }
    
    return false
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

func isWithinBounds(pos Position) bool {
    return pos.Row >= 0 && pos.Row < 8 && pos.Col >= 0 && pos.Col < 8
}
