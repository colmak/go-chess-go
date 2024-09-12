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