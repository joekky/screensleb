package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

func main() {
	file, err := os.Open("events.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	re := regexp.MustCompile("Source Network Address:\t([0-9.]+)")

	f, err := os.Create("firewall.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	for scanner.Scan() {
		line := scanner.Text()
		if re.MatchString(line) {
			ip := re.FindStringSubmatch(line)[1]
			_, err := f.WriteString(ip + "\n")
			if err != nil {
				log.Fatal(err)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Done!")
}
