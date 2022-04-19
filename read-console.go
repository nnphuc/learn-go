package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	c := make(chan string)

	go func() {
		for {
			text := <-c
			fmt.Println("console:", text)
		}
	}()

	for {
		fmt.Println("->")
		text, _ := reader.ReadString('\n')
		c <- text
	}
}
