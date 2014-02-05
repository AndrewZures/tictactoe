package ttt_test

import (
  . "github.com/andrewzures/tictactoe/ttt"
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
)

var _ = Describe("Basic Board", func() {
  var playerX, playerO Player
  var board Board

  BeforeEach(func(){
    playerX, playerO = getPlayers()
    board = Board(new(BasicBoard))
    board.NewBoard("o")
  })

  It("can implement the Board interface", func() {
    board := Board(new(BasicBoard))
    board.NewBoard("x")
    Expect(board.RecordMove(0,playerX.Symbol())).To(Equal(true))
  })

  Context("recording player moves", func() {
    var basicBoard *BasicBoard

    BeforeEach(func(){
      basicBoard = new(BasicBoard)
      basicBoard.NewBoard("x")
    })

    Context("validates moves", func() {

      It("determines if move choice is within board bounds", func(){
        Expect(basicBoard.RecordMove(-1,playerX.Symbol())).To(Equal(false))
        Expect(basicBoard.RecordMove(0,playerX.Symbol())).To(Equal(true))

        Expect(basicBoard.RecordMove(9,playerO.Symbol())).To(Equal(false))
        Expect(basicBoard.RecordMove(8,playerO.Symbol())).To(Equal(true))
      })

      It("determines if move choice has already been selected", func(){
        expectedResult := []string{"","","","x","","","","",""}

        Expect(basicBoard.RecordMove(3,playerX.Symbol())).To(Equal(true))
        Expect(basicBoard.Array()).To(Equal(expectedResult))

        Expect(basicBoard.RecordMove(3,playerO.Symbol())).To(Equal(false))
        Expect(basicBoard.Array()).To(Equal(expectedResult))
      })
    })

    Context("adds state to internal board array", func() {

      It("records a player's move", func() {
        expectedResult := []string{"x","","","","","","","",""}
        basicBoard.RecordMove(0,playerX.Symbol())
        Expect(basicBoard.Array()).To(Equal(expectedResult))
      })

      It("records multiple moves", func() {
        expectedResult := []string{"x","","","x","","","x","","x"}
        board := GenerateBoard([]string{"x","","","x","","","x","","x"})
        Expect(board.Array()).To(Equal(expectedResult))
      })
    })

    Context("keeps track of player turns", func() {

      It("records a player's move", func() {
        board := GenerateBoard([]string{"","","","","","","","",""})
        Expect(board.PlayerTurn()).To(Equal("o"))

        board.RecordMove(0, "x")

        Expect(board.PlayerTurn()).To(Equal("o"))

        board.RecordMove(1, "o")

        Expect(board.PlayerTurn()).To(Equal("x"))
      })

      It("records a player's move also", func() {
        board := GenerateBoard([]string{"","","","","","","","",""})
        moveResult := board.RecordMove(0, "o")
        Expect(moveResult).To(Equal(true))

        Expect(board.PlayerTurn()).To(Equal("x"))
        Expect(board.RecordMove(1,"o")).To(Equal(false))
        Expect(board.RecordMove(2,"o")).To(Equal(false))
        Expect(board.RecordMove(3,"o")).To(Equal(false))
        Expect(board.RecordMove(1,"x")).To(Equal(true))
      })
    })
  })

  Context("holding and accessing board status", func() {
    var emptyBoardState []string

    BeforeEach(func(){
      emptyBoardState = make([]string, 9, 9)
    })

    It("initalizes with empty board", func() {
      board := GenerateBoard([]string{"","","","","","","","",""})
      Expect(board.Array()).To(Equal(emptyBoardState))
      Expect(len(board.Array())).To(Equal(9))
    })

    It("keeps game state after attempted reset", func() {
      board := GenerateBoard([]string{"","x","","","x","","","",""})
      board.NewBoard("x")
      Expect(board.Array()).NotTo((Equal(emptyBoardState)))
    })

    It("returns list of available moves after single move", func() {
      gameState := []string{"","x","","","","","","",""}
      expectedResult := []int{0,2,3,4,5,6,7,8}
      Expect(board.OpenSpots(gameState)).To(Equal(expectedResult))
    })

    It("returns list of available moves after multiple moves", func() {
      gameState := []string{"o","x","","o","","","","","x"}
      expectedResult := []int{2,4,5,6,7}
      Expect(board.OpenSpots(gameState)).To(Equal(expectedResult))
    })
  })

  Context("when scoring a game", func() {

    Context("when getting game status", func() {

      It("returns winner if there is a row winner", func() {
        boardContents := []string{"x","x","x","","","","","",""}
        board := GenerateBoard(boardContents)
        Expect(board.Status()).To(Equal("x"))
      })

      It("returns winner if there is a column winner", func() {
        boardContents := []string{"o","","","o","","","o","",""}
        board := GenerateBoard(boardContents)
        Expect(board.Status()).To(Equal("o"))
      })

      It("returns winner if there is a diagonal winner", func() {
        boardContents := []string{"","","x","","x","","x","",""}
        board := GenerateBoard(boardContents)
        Expect(board.Status()).To(Equal("x"))
      })

      It("returns tie if no winner and no more avialable moves", func() {
        boardContents := []string{"o","x","o","o","x","o","x","o","x"}
        board := GenerateBoard(boardContents)
        Expect(board.Status()).To(Equal("tie"))
      })

      It("returns inproress if no winner but more moves available", func() {
        boardContents := []string{"o","","","","x","","","",""}
        board := GenerateBoard(boardContents)
        Expect(board.Status()).To(Equal("inprogress"))

        boardContents = []string{"o","x","x","","x","o","o","",""}
        board = GenerateBoard(boardContents)
        Expect(board.Status()).To(Equal("inprogress"))
      })

    })

    Context("when scoring a row", func() {

      It("finds a row winner", func() {
        boardContents := []string{"x","x","x","","","","","",""}
        Expect(board.RowWinner(boardContents)).To(Equal("x"))

        boardContents = []string{"","","","o","o","o","","",""}
        Expect(board.RowWinner(boardContents)).To(Equal("o"))

        boardContents = []string{"","","","","","","x","x","x"}
        Expect(board.RowWinner(boardContents)).To(Equal("x"))
      })

      It("does not generate false positives", func() {

        boardContents := []string{"","","","","","","","",""}
        board := GenerateBoard(boardContents)
        Expect(board.RowWinner(boardContents)).To(Equal(""))

        boardContents = []string{"x","","","x","","","x","",""}
        board = GenerateBoard(boardContents)
        Expect(board.RowWinner(boardContents)).To(Equal(""))

        boardContents = []string{"x","","","","x","","","","x"}
        board = GenerateBoard(boardContents)
        Expect(board.RowWinner(boardContents)).To(Equal(""))
      })

    })

    Context("when scoring a column", func() {

      It("finds a column winner", func() {
        boardContents := []string{"x","","","x","","","x","",""}
        Expect(board.ColumnWinner(boardContents)).To(Equal("x"))

        boardContents = []string{"","o","","","o","","","o",""}
        Expect(board.ColumnWinner(boardContents)).To(Equal("o"))

        boardContents = []string{"","","x","","","x","","","x"}
        Expect(board.ColumnWinner(boardContents)).To(Equal("x"))
      })

      It("does not generate false positives", func() {

        boardContents := []string{"","","","","","","","",""}
        Expect(board.ColumnWinner(boardContents)).To(Equal(""))

        boardContents = []string{"","","","x","x","x","","",""}
        Expect(board.ColumnWinner(boardContents)).To(Equal(""))

        boardContents = []string{"x","","","","x","","","","x"}
        Expect(board.ColumnWinner(boardContents)).To(Equal(""))
      })

    })

    Context("when scoring a diagonal", func() {

      It("finds a diagonal winner", func() {
        boardContents := []string{"x","","","","x","","","","x"}
        Expect(board.DiagonalWinner(boardContents)).To(Equal("x"))

        boardContents = []string{"","","o","","o","","o","",""}
        Expect(board.DiagonalWinner(boardContents)).To(Equal("o"))
      })

      It("does not generate false positives", func() {

        boardContents := []string{"","","","","","","","",""}
        Expect(board.DiagonalWinner(boardContents)).To(Equal(""))

        boardContents = []string{"","","","","","","x","x","x"}
        Expect(board.DiagonalWinner(boardContents)).To(Equal(""))

        boardContents = []string{"","","o","","","o","","","o"}
        Expect(board.DiagonalWinner(boardContents)).To(Equal(""))
      })

    })

  })

})

func getPlayers() (Player, Player) {
  human1 := new(HumanPlayer)
  human1.SetSymbol("x")

  human2 := new(HumanPlayer)
  human2.SetSymbol("o")

  return Player(human1), Player(human2)
}
