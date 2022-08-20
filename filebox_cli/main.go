package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/Shulammite-Aso/filebox-cli/auth"
	"github.com/Shulammite-Aso/filebox-cli/box"
	"github.com/Shulammite-Aso/filebox-cli/util"
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
	scanner := bufio.NewScanner(os.Stdin)

	actionCaller := func() {
		switch action {
		case Register:
			auth.Register()
		case Login:
			auth.Login()
		case SendFile:
			box.SendFile()
		case UpdateFile:
			box.UpdateFile()
		case GetFile:
			box.GetFile()
		case GetListOfAllFiles:
			box.GetListOfAllFiles()
		case DeleteFile:
			box.DeleteFile()
		case SendFileToPerson:
			box.SendFileToPerson()
		default:
			fmt.Println("Please only enter a number from the options")
			os.Exit(1)
		}
	}

	// Authenticate
	fmt.Println("What do you want to do? Choose the number and press enter\n(1) Register\n(2) Login")
	scanner.Scan()
	action, err := strconv.Atoi(scanner.Text())

	if err != nil {
		util.HandleError(err)
	}

	actionCaller()

	// After authentication, proceed with other actions
	fmt.Println("What do you want to do? Choose the number and press enter")
	fmt.Println("(3) Send file to my box")
	fmt.Println("(4) Replace a file")
	fmt.Println("(5) Get a file")
	fmt.Println("(6) Get list of all files in my box")
	fmt.Println("(7) Delete a file")
	fmt.Println("(8) Send file to someone else")
	scanner.Scan()
	action, err = strconv.Atoi(scanner.Text())

	if err != nil {
		util.HandleError(err)
	}

	actionCaller()
}
