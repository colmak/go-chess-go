package board

func (b *Board) IsCheck(isBlack bool) bool {
    var kingPos Position
    // Find the king's position
    for row := 0; row < 8; row++ {
        for col := 0; col < 8; col++ {
            piece := b.Squares[row][col]
            if piece != 0 && (piece&King) != 0 && (piece&Black != 0) == isBlack {
                kingPos = Position{Row: row, Col: col}
                break
            }
        }
    }

    // Check if any opposing piece can attack the king
    for row := 0; row < 8; row++ {
        for col := 0; col < 8; col++ {
            piece := b.Squares[row][col]
            if piece != 0 && (piece&Black != 0) != isBlack {
                if b.isLegalMove(piece, Position{Row: row, Col: col}, kingPos) {
                    return true // King is in check
                }
            }
        }
    }

    return false
}


// IsCheckmate checks if the current player is in checkmate
func (b *Board) IsCheckmate(isBlack bool) bool {
    if !b.IsCheck(isBlack) {
        return false // The king is not in check
    }

    // Check if there are any legal moves to get out of check
    for row := 0; row < 8; row++ {
        for col := 0; col < 8; col++ {
            piece := b.Squares[row][col]
            if piece != 0 && (piece&Black != 0) == isBlack {
                possibleMoves := b.GenerateMoves(Position{Row: row, Col: col})
                for _, move := range possibleMoves {
                    // Create a temporary copy of the board and simulate the move
                    tempBoard := *b
                    tempBoard.MovePiece(Position{Row: row, Col: col}, move)
                    if !tempBoard.IsCheck(isBlack) {
                        return false // There's a move that gets out of check
                    }
                }
            }
        }
    }

    return true // No legal moves, checkmate
}

// IsStalemate checks if the current player is in stalemate
func (b *Board) IsStalemate(isBlack bool) bool {
    if b.IsCheck(isBlack) {
        return false // Stalemate only occurs when the king is NOT in check
    }

    // Check if there are any legal moves
    for row := 0; row < 8; row++ {
        for col := 0; col < 8; col++ {
            piece := b.Squares[row][col]
            if piece != 0 && (piece&Black != 0) == isBlack {
                possibleMoves := b.GenerateMoves(Position{Row: row, Col: col})
                for _, move := range possibleMoves {
                    // Create a temporary copy of the board and simulate the move
                    tempBoard := *b
                    tempBoard.MovePiece(Position{Row: row, Col: col}, move)
                    if !tempBoard.IsCheck(isBlack) {
                        return false // There are legal moves, no stalemate
                    }
                }
            }
        }
    }

    return true // No legal moves, stalemate
}
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


func (b *Board) canCastleKingside(isBlack bool) bool {
    if isBlack {
        return !b.BlackKingMoved && !b.BlackRookMoved[1] && b.isPathClear(Position{7, 4}, Position{7, 7}) && !b.IsCheck(true)
    }
    return !b.WhiteKingMoved && !b.WhiteRookMoved[1] && b.isPathClear(Position{0, 4}, Position{0, 7}) && !b.IsCheck(false)
}

func (b *Board) canCastleQueenside(isBlack bool) bool {
    if isBlack {
        return !b.BlackKingMoved && !b.BlackRookMoved[0] && b.isPathClear(Position{7, 4}, Position{7, 0}) && !b.IsCheck(true)
    }
    return !b.WhiteKingMoved && !b.WhiteRookMoved[0] && b.isPathClear(Position{0, 4}, Position{0, 0}) && !b.IsCheck(false)
}

