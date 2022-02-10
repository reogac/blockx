package main

import (
	"fmt"
	//"blockx/models"
	"blockx/core"
)

func main() {
	fmt.Println("vim-go")
	app := core.NewApp()
	app.Run()
	app.Stop()
}
