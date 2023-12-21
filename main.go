package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

//type Data struct {
//	Intro Intro
//}

type Options struct {
	Text string
	Arc  string
}

type Section struct {
	Title   string
	Story   []string
	Options []Options
}

func main() {

	content, err := os.ReadFile("./story.json")
	if err != nil {
		log.Fatal("Error opening file: ", err)
	}

	var payload map[string]Section
	err = json.Unmarshal(content, &payload)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}

	println("starting loop")
	var selection string = "intro"
	for selection != "exit" {

		selection = getselection(selection, payload)
		fmt.Println("Returned Arc is :", selection)

	}

	log.Printf(payload["Intro"].Title)
}

func getselection(currentSelection string, payload map[string]Section) string {

	println("Select next action")
	if section, ok := payload[currentSelection]; ok {
		optionSlice := section.Options
		for i := 0; i < len(optionSlice); i++ {
			option := optionSlice[i]
			fmt.Printf("%d. %s\n", i, option.Text)
		}
		in := bufio.NewReader(os.Stdin)
		var s string
		println("What Option would you like ? Please enter the number")
		s, _ = in.ReadString('\n')
		s = strings.TrimSpace(s)
		if index, err := strconv.Atoi(s); err == nil && index >= 0 && index < len(optionSlice) {
			return optionSlice[index-1].Arc
		} else {
			fmt.Println("Invalid option. Please enter a valid number.")
		}
	}
	return "exit"
}

func displaystory(currentSelection string, payload map[string]Section) {

}
