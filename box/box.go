package box

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/Shulammite-Aso/filebox-cli/auth"
	"github.com/Shulammite-Aso/filebox-cli/util"
)

var (
	host          = "http://34.72.244.219:3000/"
	authorization *string
)

func SendFile() {

	var path string

	fmt.Println("Enter path to file:")
	fmt.Scanln(&path)

	fmt.Println("please wait...")

	// Verify file size
	if err := util.VerifySize(path); err != nil {
		util.HandleError(err)
	}

	data, err := os.ReadFile(path)
	if err != nil {
		util.HandleError(err)
	}

	fileName := filepath.Base(path)

	values := map[string]interface{}{"fileName": fileName, "file": data}
	json_data, err := json.Marshal(values)

	if err != nil {
		util.HandleError(err)
	}

	authorization = &auth.Token

	request, err := http.NewRequest("POST", host+"/filebox/sendfile", bytes.NewBuffer(json_data))

	if err != nil {
		util.HandleError(err)
	}

	request.Header.Set("Content-Type", "application/json; charset=UTF-8")
	request.Header.Set("authorization", "Bearer "+*authorization)

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		util.HandleError(err)
	}
	defer response.Body.Close()

	var res map[string]string

	json.NewDecoder(response.Body).Decode(&res)

	if res["error"] != "" {
		util.HandleError(errors.New(res["error"]))
	}

	fmt.Println(res["message"])

}

func GetFile() {

	var fileName string

	fmt.Println("Enter file name:")
	fmt.Scanln(&fileName)

	fmt.Println("please wait...")

	authorization = &auth.Token

	request, err := http.NewRequest("GET", host+"/filebox/getfile/"+fileName, nil)

	if err != nil {
		util.HandleError(err)
	}

	request.Header.Set("authorization", "Bearer "+*authorization)

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		util.HandleError(err)
	}
	defer response.Body.Close()

	var res map[string]string

	json.NewDecoder(response.Body).Decode(&res)

	if res["error"] != "" {
		util.HandleError(errors.New(res["error"]))
	}

	fmt.Println("downloading file...")

	fileUrl := res["fileURL"]

	err = util.DownloadFile(fileUrl, fileName)

	if err != nil {
		util.HandleError(err)
	}

	fmt.Println("finished!")

}

func UpdateFile() {

	var path string

	fmt.Println("Enter path to file:")
	fmt.Scanln(&path)

	fmt.Println("please wait...")

	// Verify file size
	if err := util.VerifySize(path); err != nil {
		util.HandleError(err)
	}

	data, err := os.ReadFile(path)
	if err != nil {
		util.HandleError(err)
	}

	fileName := filepath.Base(path)

	values := map[string]interface{}{"fileName": fileName, "file": data}
	json_data, err := json.Marshal(values)

	if err != nil {
		util.HandleError(err)
	}

	authorization = &auth.Token

	request, err := http.NewRequest("PUT", host+"/filebox/updatefile", bytes.NewBuffer(json_data))

	if err != nil {
		util.HandleError(err)
	}

	request.Header.Set("Content-Type", "application/json; charset=UTF-8")
	request.Header.Set("authorization", "Bearer "+*authorization)

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		util.HandleError(err)
	}
	defer response.Body.Close()

	var res map[string]string

	json.NewDecoder(response.Body).Decode(&res)

	if res["error"] != "" {
		util.HandleError(errors.New(res["error"]))
	}

	fmt.Println(res["message"])

}

func GetListOfAllFiles() {

	fmt.Println("please wait...")

	authorization = &auth.Token

	request, err := http.NewRequest("GET", host+"/filebox/get-list-of-all-files", nil)

	if err != nil {
		util.HandleError(err)
	}

	request.Header.Set("authorization", "Bearer "+*authorization)

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		util.HandleError(err)
	}
	defer response.Body.Close()

	var res map[string][]string

	json.NewDecoder(response.Body).Decode(&res)

	allFiles := res["allFiles"]

	for _, file := range allFiles {
		fmt.Println(file)
	}

}

func SendFileToPerson() {

	var path string
	var receiverUsername string

	fmt.Println("Enter path to file:")
	fmt.Scanln(&path)

	fmt.Println("Enter receiver username:")
	fmt.Scanln(&receiverUsername)

	fmt.Println("please wait...")

	// Verify file size
	if err := util.VerifySize(path); err != nil {
		util.HandleError(err)
	}

	data, err := os.ReadFile(path)
	if err != nil {
		util.HandleError(err)
	}

	fileName := filepath.Base(path)

	values := map[string]interface{}{"fileName": fileName, "file": data, "receiverUsername": receiverUsername}
	json_data, err := json.Marshal(values)

	if err != nil {
		util.HandleError(err)
	}

	authorization = &auth.Token

	request, err := http.NewRequest("POST", host+"/filebox/sendfile-to-person", bytes.NewBuffer(json_data))

	if err != nil {
		util.HandleError(err)
	}

	request.Header.Set("Content-Type", "application/json; charset=UTF-8")
	request.Header.Set("authorization", "Bearer "+*authorization)

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		util.HandleError(err)
	}
	defer response.Body.Close()

	var res map[string]string

	json.NewDecoder(response.Body).Decode(&res)

	if res["error"] != "" {
		util.HandleError(errors.New(res["error"]))
	}

	fmt.Println(res["message"])

}

func DeleteFile() {

	var fileName string

	fmt.Println("Enter file name:")
	fmt.Scanln(&fileName)

	fmt.Println("please wait...")

	authorization = &auth.Token

	request, err := http.NewRequest("DELETE", host+"/filebox/deletefile/"+fileName, nil)

	if err != nil {
		util.HandleError(err)
	}

	request.Header.Set("authorization", "Bearer "+*authorization)

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		util.HandleError(err)
	}
	defer response.Body.Close()

	var res map[string]string

	json.NewDecoder(response.Body).Decode(&res)

	if res["error"] != "" {
		util.HandleError(errors.New(res["error"]))
	}

	fmt.Println(res["message"])

}
