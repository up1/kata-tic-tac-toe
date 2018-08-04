# kata-tic-tac-toe

```
$go test -v -cover -coverprofile=coverage.out  ./...
$go tool cover -html=coverage.out
```

### Results
```
=== RUN   TestNotWinner
--- PASS: TestNotWinner (0.00s)
=== RUN   TestWinnerXfOInFirstRow
--- PASS: TestWinnerXfOInFirstRow (0.00s)
=== RUN   TestWinnerXfXInDiagonal
--- PASS: TestWinnerXfXInDiagonal (0.00s)
=== RUN   TestWinnerXfXInFirstColumn
--- PASS: TestWinnerXfXInFirstColumn (0.00s)
=== RUN   TestWinnerOfOInFirstRow
--- PASS: TestWinnerOfOInFirstRow (0.00s)
=== RUN   TestWinnerOfOInDiagonalLefttoRightWithTwoPlayers
--- PASS: TestWinnerOfOInDiagonalLefttoRightWithTwoPlayers (0.00s)
=== RUN   TestWinnerOfOInDiagonalLefttoRight
--- PASS: TestWinnerOfOInDiagonalLefttoRight (0.00s)
=== RUN   TestWinnerOfOInFirstColumn
--- PASS: TestWinnerOfOInFirstColumn (0.00s)
=== RUN   TestWinnerOfOInSecondRow
--- PASS: TestWinnerOfOInSecondRow (0.00s)
=== RUN   TestEndGameWithDraw_GameShouldNotHaveStatusFinish
--- PASS: TestEndGameWithDraw_GameShouldNotHaveStatusFinish (0.00s)
=== RUN   TestEndGameWithDraw_GameShouldHaveStatusFinish
--- PASS: TestEndGameWithDraw_GameShouldHaveStatusFinish (0.00s)
=== RUN   TestEndGameWithDraw
--- PASS: TestEndGameWithDraw (0.00s)
=== RUN   TestPlayer1andPlayer2_ShouldPlayCorrectly
--- PASS: TestPlayer1andPlayer2_ShouldPlayCorrectly (0.00s)
=== RUN   TestPlayer1andPlayer2_ShouldNotPlayInSamePosition
--- PASS: TestPlayer1andPlayer2_ShouldNotPlayInSamePosition (0.00s)
=== RUN   TestPlayWithX_inBoard_at_first_position_Shouldupdated_status_of_board
--- PASS: TestPlayWithX_inBoard_at_first_position_Shouldupdated_status_of_board (0.00s)
=== RUN   TestPlayWithO_inBoard_at_first_position_Shouldupdated_status_of_board
--- PASS: TestPlayWithO_inBoard_at_first_position_Shouldupdated_status_of_board (0.00s)
=== RUN   TestCreateNewGame_ShouldCreateEmptyBoard_and_have_emptyStatusInAllPosition
--- PASS: TestCreateNewGame_ShouldCreateEmptyBoard_and_have_emptyStatusInAllPosition (0.00s)
=== RUN   TestCreateNewGame_ShouldCreateEmpytyBoard_withSize3X3
--- PASS: TestCreateNewGame_ShouldCreateEmpytyBoard_withSize3X3 (0.00s)
PASS
coverage: 100.0% of statements
ok  	ox	0.011s	coverage: 100.0% of statements
```
