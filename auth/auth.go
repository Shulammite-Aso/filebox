package auth

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

var Token string

func Register() {

	var email string
	var username string
	var password string
	var ComfirmPassword string

	fmt.Println("Enter email:")
	fmt.Scanln(&username)

	fmt.Println("Enter username:")
	fmt.Scanln(&username)

	fmt.Println("Enter password:")
	fmt.Scanln(&password)

	fmt.Println("Comfirm password:")
	fmt.Scanln(&ComfirmPassword)

	if password != ComfirmPassword {
		fmt.Println("Your passwords do not match")
		os.Exit(2)
	}

	values := map[string]string{"email": email, "username": username, "password": password}
	json_data, err := json.Marshal(values)

	if err != nil {
		log.Fatal(err)
	}

	resp, err := http.Post("/auth/register", "application/json",
		bytes.NewBuffer(json_data))

	if err != nil {
		log.Fatal(err)
	}

	var res map[string]string

	json.NewDecoder(resp.Body).Decode(&res)

	if res["error"] != "" {
		fmt.Println(res["error"])
		os.Exit(1)
	}

	Token = res["token"]

	////
	fmt.Println(Token)

	fmt.Println(res["message"])

}

func Login() {

	var username string
	var password string

	fmt.Println("Enter your username:")
	fmt.Scanln(&username)

	fmt.Println("Enter your password:")
	fmt.Scanln(&password)

	values := map[string]string{"username": username, "password": password}
	json_data, err := json.Marshal(values)

	if err != nil {
		log.Fatal(err)
	}

	resp, err := http.Post("https://3000-shulammitea-fileboxauth-o84h57cmzok.ws-eu60.gitpod.io/auth/login", "application/json",
		bytes.NewBuffer(json_data))

	if err != nil {
		log.Fatal(err)
	}

	var res map[string]string

	json.NewDecoder(resp.Body).Decode(&res)

	if res["error"] != "" {
		fmt.Println(res["error"])
		os.Exit(1)
	}

	Token = res["token"]

	////
	fmt.Println(Token)
	fmt.Println("login successful")

}
