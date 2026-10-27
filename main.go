package main

import (
	"fmt"
)

func main() {
	//reader := bufio.NewReader(os.Stdin)
	//fmt.Print("What author would you like to search for? ")
	//author, err := reader.ReadString('\n')
	//author = strings.TrimSpace(author)
	//if err != nil {
	//	fmt.Println("Error reading input:", err)
	//	return
	//}
	bookData, err := getBookData("riordan")
	if err != nil {
		fmt.Println("Error fetching book data:", err)
		return
	}
	gui := *buildBookGUI(bookData)
	gui.ShowAndRun()
}
