package ox_test

import "testing"
import "ox"

func TestNotWinner(t *testing.T) {
	game := ox.NewGame()
	game.Play("X", 1)
	game.Play("O", 2)
	game.Play("X", 3)
	if game.IsWinner() {
		t.Fatalf("Game should not have the winner")
	}
}

func TestWinnerXfOInFirstRow(t *testing.T) {
	game := ox.NewGame()
	game.Play("X", 1)
	game.Play("X", 2)
	game.Play("X", 3)
	if !game.IsWinner() {
		t.Fatalf("Game should have the winner")
	}
}

func TestWinnerXfXInDiagonal(t *testing.T) {
	game := ox.NewGame()
	game.Play("X", 1)
	game.Play("X", 5)
	game.Play("X", 9)
	if !game.IsWinner() {
		t.Fatalf("Game should have the winner")
	}
}

func TestWinnerXfXInFirstColumn(t *testing.T) {
	game := ox.NewGame()
	game.Play("X", 1)
	game.Play("X", 4)
	game.Play("X", 7)
	if !game.IsWinner() {
		t.Fatalf("Game should have the winner")
	}
}

func TestWinnerOfOInFirstRow(t *testing.T) {
	game := ox.NewGame()
	game.Play("O", 1)
	game.Play("O", 2)
	game.Play("O", 3)
	if !game.IsWinner() {
		t.Fatalf("Game should have the winner")
	}
}

func TestWinnerOfOInDiagonalLefttoRightWithTwoPlayers(t *testing.T) {
	game := ox.NewGame()
	game.Play("O", 1)
	game.Play("X", 2)
	game.Play("O", 5)
	game.Play("X", 6)
	game.Play("O", 9)
	if game.CurrentBoard().Diagonal[0] != 3 {
		t.Fatalf("Colume 1 should be 3 but got %v", game.CurrentBoard().Diagonal[0])
	}
}

func TestWinnerOfOInDiagonalLefttoRight(t *testing.T) {
	game := ox.NewGame()
	game.Play("O", 1)
	game.Play("O", 5)
	game.Play("O", 9)
	if game.CurrentBoard().Diagonal[0] != 3 {
		t.Fatalf("Colume 1 should be 3 but got %v", game.CurrentBoard().Diagonal[0])
	}
}

func TestWinnerOfOInFirstColumn(t *testing.T) {
	game := ox.NewGame()
	game.Play("O", 1)
	game.Play("O", 4)
	game.Play("O", 7)
	if game.CurrentBoard().Column[0] != 3 {
		t.Fatalf("Colume 1 should be 3 but got %v", game.CurrentBoard().Column[0])
	}
}

func TestWinnerOfOInSecondRow(t *testing.T) {
	game := ox.NewGame()
	game.Play("O", 4)
	game.Play("O", 5)
	game.Play("O", 6)
	if game.CurrentBoard().Row[1] != 3 {
		t.Fatalf("Row 2 should be 3 but got %v", game.CurrentBoard().Row[1])
	}
}

func TestEndGameWithDraw_GameShouldNotHaveStatusFinish(t *testing.T) {
	game := ox.NewGame()
	game.Play("O", 1)
	game.Play("X", 2)
	status := game.CurrentStatus()
	if status != "playing" {
		t.Fatalf("Status of game should be playing, but got %v", status)
	}
}

func TestEndGameWithDraw_GameShouldHaveStatusFinish(t *testing.T) {
	game := ox.NewGame()
	game.Play("O", 1)
	game.Play("X", 2)
	game.Play("O", 3)
	game.Play("X", 4)
	game.Play("O", 5)
	game.Play("X", 6)
	game.Play("O", 7)
	game.Play("X", 8)
	game.Play("O", 9)
	status := game.CurrentStatus()
	if status != "finish" {
		t.Fatalf("Status of game should be finish, but got %v", status)
	}
}

func TestEndGameWithDraw(t *testing.T) {
	expextedStatus := [9]string{"O", "X", "O", "X", "O", "X", "O", "X", "O"}
	game := ox.NewGame()
	game.Play("O", 1)
	game.Play("X", 2)
	game.Play("O", 3)
	game.Play("X", 4)
	game.Play("O", 5)
	game.Play("X", 6)
	game.Play("O", 7)
	game.Play("X", 8)
	game.Play("O", 9)
	board := game.CurrentBoard()
	if board.Status != expextedStatus {
		t.Fatalf("All position in board should be %v, but got %v", expextedStatus, board.Status)
	}
}

func TestPlayer1andPlayer2_ShouldPlayCorrectly(t *testing.T) {
	game := ox.NewGame()
	game.Play("O", 1)
	err := game.Play("X", 2)
	if err != nil {
		t.Fatalf("Should play collectly")
	}
}

func TestPlayer1andPlayer2_ShouldNotPlayInSamePosition(t *testing.T) {
	game := ox.NewGame()
	game.Play("O", 1)
	err := game.Play("X", 1)
	if err == nil {
		t.Fatalf("Can not play in the same or not empty position")
	}
}

func TestPlayWithX_inBoard_at_first_position_Shouldupdated_status_of_board(t *testing.T) {
	expextedStatus := [9]string{"X", "", "", "", "", "", "", "", ""}
	game := ox.NewGame()
	game.Play("X", 1)
	board := game.CurrentBoard()
	if board.Status != expextedStatus {
		t.Fatalf("All position in board should be %v, but got %v", expextedStatus, board.Status)
	}
}

func TestPlayWithO_inBoard_at_first_position_Shouldupdated_status_of_board(t *testing.T) {
	expextedStatus := [9]string{"O", "", "", "", "", "", "", "", ""}
	game := ox.NewGame()
	game.Play("O", 1)
	board := game.CurrentBoard()
	if board.Status != expextedStatus {
		t.Fatalf("All position in board should be %v, but got %v", expextedStatus, board.Status)
	}
}
func TestCreateNewGame_ShouldCreateEmptyBoard_and_have_emptyStatusInAllPosition(t *testing.T) {
	expextedStatus := [9]string{"", "", "", "", "", "", "", "", ""}
	game := ox.NewGame()
	board := game.CurrentBoard()
	if board.Status != expextedStatus {
		t.Fatalf("All position in board should be empty, but got %v", board.Status)
	}

}
func TestCreateNewGame_ShouldCreateEmpytyBoard_withSize3X3(t *testing.T) {
	game := ox.NewGame()

	if game.CurrentBoard().Size != 9 {
		t.Fatalf("After created new game, board should be size 3X3 but got %v", game.CurrentBoard().Size)
	}
}
