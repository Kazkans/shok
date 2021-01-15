package main

import (
	"bufio"
	"fmt"
	"log"
	"net/url"
	"os"
	"strings"
)

func readLinks() {
	f, err := os.OpenFile(*file, os.O_CREATE|os.O_RDONLY, 0644)
	if err != nil {
		log.Println(err)
	}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}

		parts := strings.Split(line, " ")
		links[parts[0]] = parts[1]
		ids[parts[1]] = parts[0]
	}

	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func getProtocol(link string) string {
	if len(link) < len("https://") {
		fmt.Println(link, len(link), len("https://"))
		return "https://" + link
	}

	if link[:len("http://")] == "http://" ||
		link[:len("https://")] == "https://" {
		return link
	} else {
		return "https://" + link
	}
}

func saveLink(link string, id string) {
	f, err := os.OpenFile(*file, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	if _, err := f.WriteString(fmt.Sprintf("%v %v\n", id, link)); err != nil {
		log.Fatal(err)
	}
}

func addLink(link string) string {
	link = getProtocol(link)
	_, err := url.ParseRequestURI(link)
	if err != nil {
		return ""
	}
	if _, ok := ids[link]; ok {
		return ids[link]
	}

	for {
		id := getRandomID()
		if _, ok := links[id]; !ok {
			fmt.Println("Added: ", link)
			links[id] = link
			ids[link] = id
			go saveLink(link, id)
			return id
		}
	}

}
