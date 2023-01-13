package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
)

type Creators struct {
	Creators []Creators `json:"creators"`
}
type creator struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

func main() {
	run()
}
func run() {
	jsonfile, err := os.Open("config.json")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("opend json")
	defer jsonfile.Close()

	byteValue, _ := io.ReadAll(jsonfile)
	var creators Creators

	json.Unmarshal(byteValue, &creators)
	for i := 0; i < len(creators.Creators); i++ {
		downloadContent(creators.Creators[i].Name, creators.Creators[i].Url)
	}
}

func newCreator(creatorName string) {
	err := os.Mkdir(creatorName, 0750)
	if err != nil && !os.IsExist(err) {
		log.Fatal(err)
	}
}

func downloadContent(creatorName string, url string) {
	_ = os.Chdir(creatorName)
	log.Printf("changed directory to: %v", creatorName)
	out, err := exec.Command("yt-dlp", "-i", url).Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(out))
	log.Printf("commad finnished with error: %v ", err)
}
