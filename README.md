Tic Tac Toe in Go
=======

###To Run:

1.  Install [Go](http://golang.org/doc/install)

     a.  A Good tutorial on how to set up your environment can be found [here](http://skife.org/golang/2013/03/24/go_dev_env.html).

     b.  Also, you may try to use the `go_setup` script, although it may not work on all systems

2.  clone this repository in your go source folder.  Alternatively "get get github.com/AndrewZures/tictactoe" should work

3. cd to tictactoe/

#####options
        1. go run main.go
        2. ./tictactoe
        3. go build; ./tictactoe

###To Test

before: run `.test_setup` which will download ginkgo and gomega test suite

#####options
        1. cd `ttt_test`, run `ginkgo`
        2. ./test_runner

###Game Rules:
    Simply choose board spot if using Human Player:
        1. 1-9 for 3x3 board
        2. 1-16 for 4x4 board
        3  1-25 for 4x4 board


