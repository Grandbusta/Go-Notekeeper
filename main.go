package main

import (
	"fmt"
	"notekeeper/models"
)

func main() {
	fmt.Println("Notekeeper")
	newnote := models.Note{Id: 1, Content: "This is my first note to write"}
	fmt.Println(newnote)
}
