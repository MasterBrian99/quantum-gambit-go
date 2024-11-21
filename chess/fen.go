package chess

import (
	"fmt"
	"strings"
)

type Fen struct {
	ranks                                           [8]string
	activeColor                                     Color // who's turn it is
	whiteCanCastleQueenSide, whiteCanCastleKingSide bool
	blackCanCastleQueenSide, blackCanCastleKingSide bool
	enPassantSquare                                 int //TODO: use square
	halfMoveClock                                   int //use to  50 moves rule
	fullMoveNumber                                  int
	fen                                             string
}

// rnbqkbnr/pp1ppppp/8/2p5/4P3/8/PPPP1PPP/RNBQKBNR w KQkq c6 0 2
const INITIAL_FEN = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"

func (f *Fen) Init() error {
	_, err := f.ParseFen(INITIAL_FEN)
	if err != nil {
		return err
	}
	return nil
}

func ValidateFen(fen string) (bool, error) {
	parts := strings.Split(fen, " ")
	if len(parts) < 1 {
		return false, fmt.Errorf("fen is empty")
	}
	ranks := strings.Split(parts[0], "/")
	if len(ranks) != 8 {
		return false, fmt.Errorf("fen has wrong number of ranks")
	}
	castlingRights := strings.Split(parts[2], "")
	if len(castlingRights) != 4 {
		return false, fmt.Errorf("fen has wrong number of castling rights")
	}
	return true, nil
}

func (f *Fen) ParseFen(fen string) (*Fen, error) {
	valid, err := ValidateFen(fen)
	if !valid {
		return f, err
	}
	f.fen = fen
	parts := strings.Split(fen, " ")
	ranks := strings.Split(parts[0], "/")
	for i, rank := range ranks {
		f.ranks[i] = rank
	}
	if parts[1] == "w" {
		f.activeColor = White
	} else if parts[1] == "b" {
		f.activeColor = Black
	}
	castlingRights := parts[2]
	f.whiteCanCastleKingSide = strings.Contains(castlingRights, "K")
	f.whiteCanCastleQueenSide = strings.Contains(castlingRights, "Q")
	f.blackCanCastleKingSide = strings.Contains(castlingRights, "k")
	f.blackCanCastleQueenSide = strings.Contains(castlingRights, "q")
	return f, nil
}
