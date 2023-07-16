package main

import (
	"fmt"
  "math/rand"
  "time"
)

func main() {
  min := 1
  max := 10
  chances := 10
  randNum := rand.Intn(max - min) + min
  var num int
	fmt.Println("This is a number guessing game. You have to guess a number between 1 and 10. Remember, guess carefully, as you have only 10 guesses!")
  time.Sleep(100 * time.Millisecond)
  fmt.Println("Guess a number!")
  fmt.Scanln(&num)
  if num < randNum {
    fmt.Println("Number is too small")
    chances--
    fmt.Println("Number of guesses: ", chances)
  } else if num > randNum {
    fmt.Println("Number is too big")
    chances--
    fmt.Println("Number of guesses: ", chances)
  } else if num == randNum {
    fmt.Println("You've got the right number!")
    fmt.Println("You've guessed the number in ", 10-chances+1, "guesses")
  } else {
    fmt.Println("Try again")
  }

  if chances == 0 {
    fmt.Println("You didn't guess the number!")
    fmt.Println("The number you were trying to guess was ", randNum)
  }
  
  
}
