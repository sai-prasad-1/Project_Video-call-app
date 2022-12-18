package main

import (
	"fmt"

	server "github.com/sai-prasad-1/Project_Video-call-app/server"
)

func main() {
	fmt.Println("Hello, playground")
	if err := server.Run(); err != nil {
		fmt.Println(err)
	}
}
