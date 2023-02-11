package main

import (
	"fmt"
	"github.com/eggysetiawan/go-email-blast/config"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	config.LoadEnv(".env")

	Download()

	//app.Console()

	//args := os.Args[1]
	//
	//switch args {
	//
	//case "cli":
	//	app.Console()
	//case "serve":
	//	app.Router()
	//}

}

func Split(url string) string {
	s := strings.Split(url, "/")

	fmt.Println()

	return s[5]

}

func Download() {
	//url from struct
	surl := Split("https://drive.google.com/file/d/1Z2nInO9jSqwA2XyjDq-6AXAN9M1gBAiT/view?usp=share_link,FALSE")

	url := fmt.Sprintf("https://drive.google.com/uc?id=%s&export=download", surl)
	fileName := "test_sertifikat2.pdf"
	fmt.Println("Downloading file...")

	output, err := os.Create("./tmp/" + fileName)
	defer output.Close()

	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error while downloading", url, "-", err.Error())
		return
	}
	defer response.Body.Close()

	n, err := io.Copy(output, response.Body)

	fmt.Println(n, "bytes downloaded")
}
