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

type Board struct {
    Squares [8][8]int // Exported field
}

// NewBoard creates a new chess board.
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