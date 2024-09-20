package board

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

func (b *Board) GetPieceAt(pos Position) int {
    return b.Squares[pos.Row][pos.Col]
}

func (b *Board) IsEmpty(pos Position) bool {
    return b.GetPieceAt(pos) == 0
}

func (b *Board) findKing(isBlack bool) Position {
    king := King | White
    if isBlack {
        king = King | Black
    }
    for r := 0; r < 8; r++ {
        for c := 0; c < 8; c++ {
            if b.Squares[r][c] == king {
                return Position{r, c}
            }
        }
    }
    return Position{-1, -1} // Error case if king not found
}