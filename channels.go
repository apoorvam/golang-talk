package main

import (
	"time"
    "fmt"
)
// START OMIT
func main() {
	var Ball int
	table := make(chan int)

	go player("Player 1", table)
	go player("Player 2", table)

	table <- Ball
	time.Sleep(1 * time.Second)
	<-table
}
func player(playerName string, table chan int) {
	for {
		ball := <-table
        fmt.Println(playerName)
		ball++
		time.Sleep(100 * time.Millisecond)
		table <- ball
	}
}
// END OMIT
