package main
import (
	"fmt"
	"math/rand"
	"time"
)


func main() {
   rand.Seed(time.Now().UnixNano());
   var target = rand.Intn(100) + 1;
   var guessNum int
   attempts := 0
   for i := 0; i < 10; i++ {
	   fmt.Println("Please enter your guess between 1 to 100 : ")
	   fmt.Scanln(&guessNum)
	   attempts++
	   if guessNum < target {
		 fmt.Println("Guess too low")
	   } else if guessNum > target {
		fmt.Println("Guess too high")
	   } else{
		fmt.Printf("Congratulations! You guessed it in %d attempts.\n", attempts)
		break
	   }
   }
}