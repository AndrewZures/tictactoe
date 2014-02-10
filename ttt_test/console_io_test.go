package ttt_test

import (
  . "github.com/andrewzures/tictactoe/ttt"
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
  "bytes"
)

var _ = Describe("Console UI", func() {
  var writer bytes.Buffer
  var reader bytes.Buffer
  var inOut InOutInterface

  BeforeEach(func(){
    inOut = InOutInterface(ConsoleIO{&writer, &reader})
  })

  Context("when receiving input from user", func() {

    It("reads from console", func() {
      SetMockInput(&reader, "Test\nConsole\nInput\n")
      Expect(inOut.Read()).To(Equal("Test"))
      Expect(inOut.Read()).To(Equal("Console"))
      Expect(inOut.Read()).To(Equal("Input"))
    })
  })


  AfterEach(func(){
    writer.Reset()
    reader.Reset()
  })

})

