package main

import (
	"fmt"

	"github.com/shiSock/goshiSock/server"
)

func test(trans server.Transport, Sender func(string, string, string), args []string) {
	fmt.Println("trans: ", trans)
	fmt.Println("args: ", args)
	Sender(trans.Name, trans.Channel, "Hello")
}

func main() {
	fmt.Println("Starting Server...")
	var sock server.Server
	sock.MaxConnection = 20000
	main := sock.Start("./../Engine/shiSock/shiSock", "127.0.0.1", "8080")
	main.Listen("main", test, []string{"one", "two"})

	main.Run()

}
