package main

import (
	"fmt"
	"os"

	"github.com/Shulammite-Aso/filebox-cli/auth"
)

const (
	Register = iota + 1
	Login
	SendFile
	UpdateFile
	GetFile
	GetListOfAllFiles
	DeleteFile
	SendFileToPerson
)

func main() {

	var action int

	actionCaller := func() {
		switch action {
		case Register:
			auth.Register()
		case Login:
			auth.Login()
		case SendFile:
			fmt.Println("send file")
		case UpdateFile:
			fmt.Println("update file")
		case GetFile:
			fmt.Println("get file")
		case GetListOfAllFiles:
			fmt.Println("get file")
		case DeleteFile:
			fmt.Println("get file")
		case SendFileToPerson:
			fmt.Println("get file")
		default:
			fmt.Println("Please only enter a number from the options")
			os.Exit(2)
		}
	}

	fmt.Println("What do you want to do? Choose the number and press enter\n(1) Register\n(2) Login")
	fmt.Scanln(&action)

	actionCaller()

	fmt.Println("What do you want to do? Choose the number and press enter")
	fmt.Println("(3) Send file to my box")
	fmt.Println("(4) Replace a file")
	fmt.Println("(5) Get a file")
	fmt.Println("(6) Get list of all files in my box")
	fmt.Println("(7) Delete a file")
	fmt.Println("(8) Send file to someone else")
	fmt.Scanln(&action)

	actionCaller()
}
