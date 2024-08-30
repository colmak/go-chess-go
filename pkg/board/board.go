// pkg/board/board.go
package board

const (
    Empty = 0
    Pawn  = 1
    Knight = 2
    Bishop = 3
    Rook = 4
    Queen = 5
    King = 6
    White = 8
    Black = 16
)

func isWhite(piece int) bool {
    return piece & White == White
}

func isBlack(piece int) bool {
    return piece & Black == Black
}

func getPiece(piece int) int {
    return piece & 7 // mask the color bits
}

// Board represents the chessboard and its state.
type Board struct {
    squares [8][8]int
}

// NewBoard creates a new chess board.
func NewBoard() *Board {
    return &Board{
        // Initialize board state
    }
}

func (b *Board) initPosistion() {
    b.squares[0][0] = Rook | White
    b.squares[0][1] = Knight | White
    b.squares[0][2] = Bishop | White
    b.squares[0][3] = Queen | White
    b.squares[0][4] = King | White
    b.squares[0][5] = Bishop | White
    b.squares[0][6] = Knight | White
    b.squares[0][7] = Rook | White
    for i := 0; i < 8; i++ {
        b.squares[1][i] = Pawn | White
    }

    // Set up black pieces.
    b.squares[7][0] = Rook | Black
    b.squares[7][1] = Knight | Black
    b.squares[7][2] = Bishop | Black
    b.squares[7][3] = Queen | Black
    b.squares[7][4] = King | Black
    b.squares[7][5] = Bishop | Black
    b.squares[7][6] = Knight | Black
    b.squares[7][7] = Rook | Black
    for i := 0; i < 8; i++ {
        b.squares[6][i] = Pawn | Black
    }

}