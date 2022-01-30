package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func main() {
	srcFile := "./hostnames.txt"
	dstFile := "./exmaple_hosts.txt"

	file, err := os.Open(srcFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	dFile, err := os.Create(dstFile)
	if err != nil {
		log.Fatal(err)
	}
	defer dFile.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "xyz") {
			replacedLine := []byte(strings.Replace(line, "xyz", "exampledomain", -1) + "\n")
			_, err := dFile.Write(replacedLine)
			if err != nil {
				log.Fatal(err)
			}
		}
	}

}
