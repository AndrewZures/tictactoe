package ttt_test

import (
	. "github.com/andrewzures/tictactoe/ttt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Console UI", func() {
	var messages Messages
	var boardFormatter BoardFormatter

	BeforeEach(func() {
		messages = BuildConsoleMessages()
		boardFormatter = BoardFormatter(new(ConsoleBoardFormatter))
	})

	It("returns properly formatted 3x3 board", func() {
		consoleBoard := "  |   |  \n---------\n  |   |  \n---------\n  |   |  \n"
		board := Generate3x3Board([]string{"", "", "", "", "", "", "", "", ""})
		Expect(boardFormatter.FormatBoard(board, messages)).To(Equal(consoleBoard))
	})

	It("returns board state with default x and o symbols", func() {
		consoleBoard := "X | O |  \n---------\n  |   |  \n---------\n  |   |  \n"
		board := Generate3x3Board([]string{"x", "o", "", "", "", "", "", "", ""})
		Expect(boardFormatter.FormatBoard(board, messages)).To(Equal(consoleBoard))
	})

	It("returns properly formatted 4x4 board", func() {
		consoleBoard := "  |   |   |  \n-------------\n  |   |   |  \n-------------\n  |   |   |  \n-------------\n  |   |   |  \n"
		board := GenerateEmpty4x4Board()
		Expect(boardFormatter.FormatBoard(board, messages)).To(Equal(consoleBoard))
	})

	It("returns properly formatted 5x5 board", func() {
		consoleBoard := "  |   |   |   |  \n-----------------\n  |   |   |   |  \n-----------------\n  |   |   |   |  \n-----------------\n  |   |   |   |  \n-----------------\n  |   |   |   |  \n"
		board := GenerateEmpty5x5Board()

		Expect(boardFormatter.FormatBoard(board, messages)).To(Equal(consoleBoard))
	})

})
