package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

type creator struct {
	name string
	url  string
}

func main() {
	creator := creator{name: "peter", url: "https://youtu.be/dQw4w9WgXcQ"}
	downloadContent(creator.name, creator.url)
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
