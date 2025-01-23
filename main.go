package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func main() {
	args := os.Args

	if len(args) != 2 {
		fmt.Println("Please provide a path to an HTML file")
		return
	}

	pt, err := filepath.Abs(strings.ReplaceAll(args[1], "'", ""))

	if err != nil {
		fmt.Println("Error getting absolute path")
		return
	}

	buf, err := os.ReadFile(pt)

	if err != nil {
		fmt.Println("Error opening file")
		return
	}

	regex := regexp.MustCompile(`id=(\d+)`)

	ids := regex.FindAllString(string(buf), -1)

	newIds := make([]string, len(ids))

	for i, id := range ids {
		newIds[i] = strings.Split(id, "=")[1]
	}

	fmt.Println(strings.Join(newIds, ","))
}
