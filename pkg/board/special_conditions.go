package board

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