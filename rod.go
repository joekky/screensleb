// This program reads a CSV file of event data, extracts IP addresses from lines
// that contain a specific string, and writes the IP addresses to a text file.

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

func main() {
	// Open the CSV file
	file, err := os.Open("events.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Set up a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	// Set up a regular expression to match lines that contain IP addresses
	re := regexp.MustCompile("Source Network Address:\t([0-9.]+)")

	// Create a new text file to write the IP addresses to
	f, err := os.Create("firewall.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Loop through the lines in the CSV file
	for scanner.Scan() {
		line := scanner.Text()

		// Check if the line contains an IP address
		if re.MatchString(line) {
			ip := re.FindStringSubmatch(line)[1]

			// Write the IP address to the text file
			_, err := f.WriteString(ip + "\n")
			if err != nil {
				log.Fatal(err)
			}
		}
	}

	// Check for any errors during scanning or writing
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Print a message to indicate the program has finished running
	fmt.Println("Done!")
}
