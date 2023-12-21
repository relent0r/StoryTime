package main

import (
	"bufio"
	"encoding/json"
	"log"
	"os"
	"strings"
)

type Data struct {
	Intro Intro
}

type Options struct {
	Text string
	Arc  string
}

type Intro struct {
	Title   string
	Story   []string
	Options []Options
}

func main() {

	content, err := os.ReadFile("./story.json")
	if err != nil {
		log.Fatal("Error opening file: ", err)
	}

	var payload Data
	err = json.Unmarshal(content, &payload)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}

	in := bufio.NewReader(os.Stdin)
	var s string
	s, _ = in.ReadString('\n')
	s = strings.TrimSpace(s)
	if s == "exit" {
		println("you exited")
	}

	log.Printf(payload.Intro.Title)

}
