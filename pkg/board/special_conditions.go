package board

func (b *Board) IsCheck(isBlack bool) bool {
    kingPos := b.findKing(isBlack)
    
    // Check for rook/queen in the same row/column
    for r := 0; r < 8; r++ {
        if r != kingPos.Row {
            piece := b.GetPieceAt(Position{r, kingPos.Col})
            if piece != 0 && b.isEnemyPiece(piece, isBlack) && (piece&^Black == Rook || piece&^Black == Queen) {
                return true // Rook or Queen can attack the king
            }
        }
    }
    for c := 0; c < 8; c++ {
        if c != kingPos.Col {
            piece := b.GetPieceAt(Position{kingPos.Row, c})
            if piece != 0 && b.isEnemyPiece(piece, isBlack) && (piece&^Black == Rook || piece&^Black == Queen) {
                return true // Rook or Queen can attack the king
            }
        }
    }

    // Additional check for other pieces (knights, bishops, pawns) can be added here
    return false
}



// IsCheckmate checks if the current player is in checkmate
func (b *Board) IsCheckmate(isBlack bool) bool {
    // First, check if the player is in check
    if !b.IsCheck(isBlack) {
        return false // The king is not in check
    }

    // Go through each piece on the board to see if there is any legal move
    for row := 0; row < 8; row++ {
        for col := 0; col < 8; col++ {
            piece := b.Squares[row][col]
            
            // Ensure the piece belongs to the current player
            if piece != 0 && (piece&Black != 0) == isBlack {
                possibleMoves := b.GenerateMoves(Position{Row: row, Col: col})

                // For each possible move, check if it gets the player out of check
                for _, move := range possibleMoves {
                    // Create a copy of the board to simulate the move
                    tempBoard := *b

                    // Simulate the move on the temporary board
                    tempBoard.MovePiece(Position{Row: row, Col: col}, move)

                    // Check if the player is still in check after the move
                    if !tempBoard.IsCheck(isBlack) {
                        return false // Found a move that gets out of check, no checkmate
                    }
                }
            }
        }
    }

    // No legal moves were found, so it's checkmate
    return true
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
    lastMove := b.GetLastMove()

    // Check en passant eligibility
    if (isBlack && start.Row == 3) || (!isBlack && start.Row == 4) {
        sidePawnPos := Position{Row: start.Row, Col: end.Col}
        sidePawn := b.GetPieceAt(sidePawnPos)

        if sidePawn != 0 && sidePawn&Pawn != 0 && (sidePawn&Black != 0) != isBlack {
            if lastMove.Start.Row == sidePawnPos.Row+2*direction(isBlack) && lastMove.End == sidePawnPos {
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

// isEnemyPiece checks if the piece belongs to the enemy based on the current player's color
func (b *Board) isEnemyPiece(piece int, isBlack bool) bool {
    if isBlack {
        return (piece & White) != 0 // True if the piece is White
    }
    return (piece & Black) != 0 // True if the piece is Black
}

func (b *Board) IsDrawByFiftyMoveRule() bool {
    return b.FiftyMoveCount >= 50
}

func (b *Board) IsDrawByThreefoldRepetition() bool {
    for _, count := range b.PositionHistory {
        if count >= 3 {
            return true
        }
    }
    return false
}
