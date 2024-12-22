package board

import "fmt"

func (b *Board) MovePiece(start, end Position) bool {
    piece := b.GetPieceAt(start)
    if piece == 0 {
        return false // No piece to move
    }

    if !isWithinBounds(start) || !isWithinBounds(end) {
        return false // Move out of bounds
    }

    // Ensure the piece belongs to the current player
    if (piece&White == 0 || b.CurrentTurn == White) && (piece&Black == 0 || b.CurrentTurn == Black) {
        return true // It's valid to move the piece
    }

    destPiece := b.GetPieceAt(end)
    if destPiece != 0 && (destPiece&White) == (piece&White) {
        return false // Can't capture your own piece
    }

    if !b.isLegalMove(piece, start, end) {
        return false // The move is not legal for this piece
    }

    // Perform the move
    b.Squares[end.Row][end.Col] = piece
    b.Squares[start.Row][start.Col] = 0

    // Check for pawn promotion
    if piece&Pawn != 0 {
        if (piece&White != 0 && end.Row == 7) || (piece&Black != 0 && end.Row == 0) {
            b.Squares[end.Row][end.Col] = Queen | (piece & (White | Black)) // Promote to Queen
        }
    }

    // Update position history
    key := fmt.Sprintf("%v", b.Squares)
    b.PositionHistory[key]++

    b.MoveCount++

    // Update fifty-move rule
    if (piece & 0b111) != Pawn && destPiece == 0 {
        b.FiftyMoveCount++
    } else {
        b.FiftyMoveCount = 0
    }

    // Update last move
    b.LastMove = Move{Start: start, End: end, Piece: piece}

    if b.CurrentTurn == White {
        b.CurrentTurn = Black
    } else {
        b.CurrentTurn = White
    }

    return true
}



func (b *Board) isLegalMove(piece int, start, end Position) bool {
    pieceType := piece & 0b111

    switch pieceType {
    case Rook:
        return (start.Row == end.Row || start.Col == end.Col) && b.isPathClear(start, end)
    case Knight:
        rowDiff := abs(start.Row - end.Row)
        colDiff := abs(start.Col - end.Col)
        return (rowDiff == 2 && colDiff == 1) || (rowDiff == 1 && colDiff == 2)
    case Bishop:
        return abs(start.Row-end.Row) == abs(start.Col-end.Col) && b.isPathClear(start, end)
    case Queen:
        return ((start.Row == end.Row || start.Col == end.Col) || abs(start.Row-end.Row) == abs(start.Col-end.Col)) && b.isPathClear(start, end)
    case King:
        if abs(start.Row-end.Row) <= 1 && abs(start.Col-end.Col) <= 1 {
            return true
        }

        isBlack := piece&Black != 0
        if start.Row == end.Row && abs(start.Col-end.Col) == 2 {
            intermediateCol := (start.Col + end.Col) / 2
            intermediatePos := Position{Row: start.Row, Col: intermediateCol}
            if b.IsCheck(isBlack) || b.IsCheckAfterMove(intermediatePos, isBlack) {
                return false
            }
            if end.Col == 6 {
                return (isBlack && start.Row == 7 && b.canCastleKingside(true)) || (!isBlack && start.Row == 0 && b.canCastleKingside(false))
            }
            if end.Col == 2 {
                return (isBlack && start.Row == 7 && b.canCastleQueenside(true)) || (!isBlack && start.Row == 0 && b.canCastleQueenside(false))
            }
        }
        return false
    case Pawn:
        direction := 1
        if piece&Black != 0 {
            direction = -1
        }

        if start.Col == end.Col && b.IsEmpty(end) {
            if start.Row+direction == end.Row {
                return true
            }
            if (start.Row == 1 && piece&White != 0 || start.Row == 6 && piece&Black != 0) && start.Row+2*direction == end.Row && b.IsEmpty(Position{start.Row + direction, start.Col}) {
                return true
            }
        }

        if abs(start.Col-end.Col) == 1 && start.Row+direction == end.Row && !b.IsEmpty(end) && (b.GetPieceAt(end)&White != piece&White) {
            return true
        }

        return b.canCaptureEnPassant(start, end, piece&Black != 0)
    default:
        return false
    }
}


func (b *Board) IsCheckAfterMove(pos Position, isBlack bool) bool {
    // Create a temporary copy of the board
    tempBoard := *b

    // Simulate the king's move to the new position
    king := King
    if isBlack {
        king |= Black
    } else {
        king |= White
    }

    for row := 0; row < 8; row++ {
        for col := 0; col < 8; col++ {
            if tempBoard.Squares[row][col] == king {
                tempBoard.Squares[row][col] = 0
            }
        }
    }
    tempBoard.Squares[pos.Row][pos.Col] = king

    // Check if the king is in check in the new position
    return tempBoard.IsCheck(isBlack)
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

func (b *Board) GetLastMove() Move {
    return b.LastMove
}