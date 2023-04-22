Firewall Log Analyzer

This program reads a CSV file of firewall event data, extracts IP addresses from lines that contain a specific string, and writes the IP addresses to a text file.
Usage

    Clone or download the repository to your local machine.
    Open a command-line interface (Terminal, Command Prompt, etc.) and navigate to the repository directory.
    Run the command go run main.go to execute the program.
    The program will read the events.csv file in the repository directory, and create a new file called firewall.txt that contains a list of IP addresses extracted from the CSV file.

Note: The program assumes that the events.csv file contains lines that match the regular expression "Source Network Address:\t([0-9.]+)". If the CSV file format is different, the program may not work as expected.