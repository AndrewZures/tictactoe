package ttt

type TTTFactory struct {
  playerList []Player;
}


func (f TTTFactory) PlayerTypes(userInterface UserInterface) ([]Player) {
  playerList := make([]Player, 0)

  human := new(HumanPlayer)
  human.NewHumanPlayer("", "Human", userInterface)
  playerList = append(playerList, Player(human))

  computer := new(ComputerPlayer)
  rules := Rules(new(BasicRules))
  computer.NewComputerPlayer("", "Computer", rules)
  playerList = append(playerList, Player(computer))

  return playerList
}

func (f TTTFactory) Player(playerTemplate Player, userInterface UserInterface) (Player) {
    switch {

      case playerTemplate.Description() == "Human":
        human := new(HumanPlayer)
        human.NewHumanPlayer(playerTemplate.Symbol(), "Human", userInterface)
        return human

      case playerTemplate.Description() == "Computer":
        computer := new(ComputerPlayer)
        rules := Rules(new(BasicRules))
        computer.NewComputerPlayer(playerTemplate.Symbol(), "Computer", rules)
        return computer
    }

    return nil
}

func (f TTTFactory) BoardTypes() ([]Board) {
  boardList := make([]Board, 0)

  board3x3 := new(BasicBoard)
  board3x3.NewBoard(9,3,"", "3x3 Board")
  boardList = append(boardList, Board(board3x3))

  return boardList
}

func (f TTTFactory) Board(boardTemplate Board, startSymbol string) (Board) {
  var board Board

  switch {
  case boardTemplate.Description() == "3x3 Board":
    board = new(BasicBoard)
    board.NewBoard(9,3,startSymbol, "3x3 Board")

    return board
  }

  return nil
}
